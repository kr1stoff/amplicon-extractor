package main

import (
	"amplicon-extractor/cli"
	"amplicon-extractor/extract"
	myio "amplicon-extractor/io"
)

func main() {
	args := cli.GetFlag()
	finalAmplicons := extract.ExtractMultithread(args.Fasta, args.Forward, args.Reverse, args.MaxMismatch, args.Threads)
	myio.WriteFastaArrayToFile(finalAmplicons, args.OutFile)
}
