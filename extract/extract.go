package extract

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

func Extract(fastaFile string) {
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

	data := strings.NewReader(content)
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)
	sc := seqio.NewScanner(r)

	for sc.Next() {
		s := sc.Seq().(*linear.Seq)
		fmt.Fprintf(os.Stdout, "%q %q %s\n", s.ID, s.Desc, s.Seq)
	}
}
