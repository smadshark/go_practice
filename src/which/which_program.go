package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		return
	}

	file := arguments[1]
	pathEnv := os.Getenv("PATH")
	pathSplit := filepath.SplitList(pathEnv)

	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, file)

		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			fmt.Println(mode)

			if mode.IsRegular() {
				// 실행 파일인가?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
