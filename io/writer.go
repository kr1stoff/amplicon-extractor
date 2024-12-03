package myio

import "os"

func WriteFastaArrayToFile(fastaArray [][2]string, filePath string) {
	// 打开文件
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 写入文件
	for _, amplicon := range fastaArray {
		_, err := file.WriteString(">" + amplicon[0] + "\n" + amplicon[1] + "\n")
		if err != nil {
			panic(err)
		}
	}
}
