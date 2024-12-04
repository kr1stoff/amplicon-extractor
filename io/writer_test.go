package myio

import "testing"

func TestWriteFastaArrayToFile(t *testing.T) {
	// 准备测试数据
	fastaArray := [][2]string{
		{"amplicon1", "ATCG"},
		{"amplicon2", "ATCG"},
		{"amplicon3", "ATCG"},
	}
	filePath := "../test/test.fasta"
	// 调用函数
	WriteFastaArrayToFile(fastaArray, filePath)
}
