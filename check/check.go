package check

import (
	"fmt"
	"os"
	"strings"
)

// 检查引物序列是否只包含ACGT
func IsWrongPrimer(primer string) bool {
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
func IsNotExistFile(file string) bool {
	open, err := os.Open(file)

	// 不能打开
	if err != nil {
		fmt.Println(err)
		return true
	}

	defer open.Close()
	stream := make([]byte, 0)
	_, err2 := open.Read(stream[:])

	// 不能读取
	if err2 != nil {
		fmt.Println(err2)
		return true
	}
	return false
}
