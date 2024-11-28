package main

import (
	// "amplicon-extractor/greet"
	"amplicon-extractor/check"
	"flag"
	"fmt"
	"os"
)

func main() {
	// * 命令行参数
	var (
		forward     = flag.String("f", "none", "上游引物序列.")
		reverse     = flag.String("r", "none", "下游引物序列.")
		maxMismatch = flag.Int("m", 3, "最大允许的错配数.")
		threads     = flag.Int("j", 1, "线程数.")
		fasta       = flag.String("F", "none", "输入基因组 fasta 文件.")
	)

	flag.Parse()
	fmt.Println("forward: ", *forward)
	fmt.Println("reverse: ", *reverse)
	fmt.Printf("maxMismatch: %d\n", *maxMismatch)
	fmt.Printf("threads: %d\n", *threads)
	fmt.Println("fasta: ", *fasta)

	if *forward == "none" || *reverse == "none" || check.IsWrongPrimer(*forward) || check.IsWrongPrimer(*reverse) {
		fmt.Print("Error: 请输入正确的上游引物和下游引物!\n\n")
		flag.Usage()
		os.Exit(1)
	} else if *fasta == "none" || check.IsNotExistFile(*fasta) {
		fmt.Print("Error: 请输入正确的基因组 fasta 文件!\n\n")
		flag.Usage()
		os.Exit(1)
	}
}
