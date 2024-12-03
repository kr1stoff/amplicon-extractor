package extract

import (
	"amplicon-extractor/common"
	myio "amplicon-extractor/io"
	"strings"
	"sync"

	"github.com/agnivade/levenshtein"
	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

/*
@description: 提取最短的扩增子
@param:

	fastaFile: fasta 文件路径
	forward: 正向引物序列
	reverse: 反向引物序列
	maxMismatch: 最大允许的错配数
	threads: 线程数

@return:

	[][2]string: 最短的扩增子 ID 和 序列
*/
func Extract(fastaFile string, forward string, reverse string, maxMismatch int, threads int) [][2]string {
	// 读取 fasta 文件转入字符串
	content := myio.OpenFileToString(fastaFile)
	finalAmplicons := make([][2]string, 0)

	// 解析 fasta 文件
	data := strings.NewReader(content)
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)
	sc := seqio.NewScanner(r)

	// 迭代 fasta 序列
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)
		shortestAmplicon := getShortestAmplicon(forward, reverse, maxMismatch, *s)
		if shortestAmplicon[0] != "" {
			finalAmplicons = append(finalAmplicons, shortestAmplicon)
		}
	}
	return finalAmplicons
}

/*
@description: 提取最短的扩增子, 多线程版
@param:

	fastaFile: fasta 文件路径
	forward: 正向引物序列
	reverse: 反向引物序列
	maxMismatch: 最大允许的错配数
	threads: 线程数

@return:

	[][2]string: 最短的扩增子 ID 和 序列
*/
func ExtractMultithread(fastaFile string, forward string, reverse string, maxMismatch int, threads int) [][2]string {
	// 读取 fasta 文件转入字符串
	content := myio.OpenFileToString(fastaFile)
	finalAmplicons := make([][2]string, 0)

	// 解析 fasta 文件
	data := strings.NewReader(content)
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)

	// 创建一个WaitGroup用于等待所有goroutine完成
	var wg sync.WaitGroup
	// 创建一个channel用于goroutine间通信
	seqChan := make(chan *linear.Seq)

	// 迭代 fasta 序列
	go func() {
		sc := seqio.NewScanner(r)
		for sc.Next() {
			seqChan <- sc.Seq().(*linear.Seq)
		}
		close(seqChan)
	}()

	// 处理每一行
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sequence := range seqChan {
				shortestAmplicon := getShortestAmplicon(forward, reverse, maxMismatch, *sequence)
				// 直接追加结果
				if shortestAmplicon[0] != "" {
					finalAmplicons = append(finalAmplicons, shortestAmplicon)
				}
			}
		}()
	}
	// 等待所有goroutine完成
	wg.Wait()

	return finalAmplicons
}

/*
@description: 获取最短的扩增子
@param:

	forward: 正向引物序列
	reverse: 反向引物序列
	maxMismatch: 最大允许的错配数
	s: 基因组序列

@return:

	[2]string: 最短的扩增子 ID 和 序列
*/
func getShortestAmplicon(forward string, reverse string, maxMismatch int, s linear.Seq) [2]string {
	// ! 固定最小和最大的扩增子长度
	minAmpliconLength := 80   // ddPCR 最小的
	maxAmpliconLength := 2000 // Sanger 最长的
	matchAmplicons := make(map[int][2]string, 0)
	forwardPositions := getPrimerPositionOnGenome(forward, s.Seq.String(), maxMismatch, 'F')
	reversePositions := getPrimerPositionOnGenome(reverse, s.Seq.String(), maxMismatch, 'R')

	if len(forwardPositions) > 0 && len(reversePositions) > 0 {
		for _, forwardPosition := range forwardPositions {
			for _, reversePosition := range reversePositions {
				// 计算扩增子长度
				ampliconLength := reversePosition - forwardPosition
				// 检查扩增子长度是否在阈值范围内
				// * 通过
				if ampliconLength >= minAmpliconLength && ampliconLength <= maxAmpliconLength {
					// 提取扩增子序列
					matchAmplicons[ampliconLength] = [2]string{s.ID, s.Seq.String()[forwardPosition:reversePosition]}
				}
			}

			if len(matchAmplicons) > 0 {
				// 找到最短的扩增子
				var shortestAmpliconLength int
				for length := range matchAmplicons {
					if shortestAmpliconLength == 0 {
						shortestAmpliconLength = length
					} else if length < shortestAmpliconLength {
						shortestAmpliconLength = length
					}
				}
				// 提取最短的扩增子
				return matchAmplicons[shortestAmpliconLength]
			}
		}
	}
	return [2]string{}
}

/*
@description: 获取引物在基因组上的匹配位置
@param:

	primer: 引物序列
	genomeSequence: 基因组序列
	maxMismatch: 最大允许的错配数
	ForwardOrRerverse: 引物方向, 'F' 表示正向, 'R' 表示反向

@return:

	[]int: 引物在基因组上的匹配位置切片s
*/
func getPrimerPositionOnGenome(primer string, genomeSequence string, maxMismatch int,
	ForwardOrRerverse rune) []int {
	// 获取引物序列
	primers := ExpandDegenerateBases(primer)
	primerLength := len(primers[0])
	// 基因组上引物匹配位置切片
	var genomePositionSlices []int

	for i := 0; i < len(genomeSequence)-primerLength; i++ {
		genomeSlice := genomeSequence[i : i+primerLength]
		for _, primer := range primers {
			if ForwardOrRerverse == 'F' {
				if levenshtein.ComputeDistance(genomeSlice, primer) <= maxMismatch {
					genomePositionSlices = append(genomePositionSlices, i)
				}
			} else if ForwardOrRerverse == 'R' {
				// 反转引物序列
				revcomp := linear.NewSeq("example DNA", []alphabet.Letter(primer), alphabet.DNAredundant)
				revcomp.RevComp()
				if levenshtein.ComputeDistance(genomeSlice, revcomp.String()) <= maxMismatch {
					genomePositionSlices = append(genomePositionSlices, i+primerLength)
				}
			}
		}
	}
	uniqueSlice := common.RemoveDupplicates(genomePositionSlices)
	return uniqueSlice
}
