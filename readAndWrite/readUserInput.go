package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file, e := os.Open("product.txt")
	if e!=nil{
		fmt.Println("打开文件错误")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		inputString, error := reader.ReadString('\n')
		if error!=nil{
			fmt.Println("读取文件错误")
			return
		}
		fmt.Printf("读出来的是: %s \n",inputString)
		if error==io.EOF{
			return
		}
	}

}
