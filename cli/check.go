package cli

import (
	"os"
	"strings"
)

// 检查引物序列是否只包含ACGT
func isWrongPrimer(primer string) bool {
	// 支持简并
	allowed := "ATCGRYSWKMBDHVNatucgryswkmbdhvn"

	for _, c := range primer {
		if !strings.ContainsAny(string(c), allowed) {
			return true
		}
	}
	return false
}

// 检查输入 fasta 文件是否存在
func isNotExistFile(file string) bool {
	open, err := os.Open(file)

	// 不能打开
	if err != nil {
		return true
	}
	defer open.Close()

	// 不能读取
	stream := make([]byte, 0)
	_, err2 := open.Read(stream[:])
	return err2 != nil
}
