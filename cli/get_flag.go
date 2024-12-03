package cli

import (
	"flag"
	"fmt"
	"os"
)

// 定义一个结构体保存命令行参数
type CommandLineArgs struct {
	Forward     string
	Reverse     string
	MaxMismatch int
	Threads     int
	Fasta       string
}

func GetFlag() CommandLineArgs {
	var args CommandLineArgs
	// * 命令行参数

	flag.StringVar(&args.Forward, "f", "none", "上游引物序列.")
	flag.StringVar(&args.Reverse, "r", "none", "下游引物序列.")
	flag.IntVar(&args.MaxMismatch, "m", 3, "最大允许的错配数.")
	flag.IntVar(&args.Threads, "j", 1, "线程数.")
	flag.StringVar(&args.Reverse, "F", "none", "输入基因组 fasta 文件.")

	flag.Parse()
	fmt.Println("Forward: ", args.Forward)
	fmt.Println("Reverse: ", args.Reverse)
	fmt.Printf("MaxMismatch: %d\n", args.MaxMismatch)
	fmt.Printf("Threads: %d\n", args.Threads)
	fmt.Println("Fasta: ", args.Fasta)

	if args.Forward == "none" || args.Reverse == "none" || isWrongPrimer(args.Forward) || isWrongPrimer(args.Reverse) {
		fmt.Print("Error: 请输入正确的上游引物和下游引物!\n\n")
		flag.Usage()
		os.Exit(1)
	} else if args.Fasta == "none" || isNotExistFile(args.Fasta) {
		fmt.Print("Error: 请输入正确的基因组 fasta 文件!\n\n")
		flag.Usage()
		os.Exit(1)
	}

	return args
}
