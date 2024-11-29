package extract

import (
	"amplicon-extractor/common"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

func Extract(fastaFile string, forward string, reverse string, maxMismatch int, threads int) {
	// 打开文件
	file, err := os.Open(fastaFile)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer file.Close()

	// 读取文件内容到字符串
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件内容出错:", err)
		return
	}

	// 解析 fasta 文件
	data := strings.NewReader(content)
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)
	sc := seqio.NewScanner(r)
	// 迭代 fasta 序列
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)
		fmt.Fprintf(os.Stdout, "%q %q %s\n", s.ID, s.Desc, s.Seq)
	}
}

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
				revcomp := linear.NewSeq("example DNA", []alphabet.Letter(primer), alphabet.DNA)
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
