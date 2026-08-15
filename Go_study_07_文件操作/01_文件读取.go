package main

import (
	"bufio"
	"fmt"
	"os"
)

// main 展示读取文件的几种方式(一次性读取、分片读、带缓冲读)
// main demonstrates several ways to read a file (at once, by chunks, buffered).
func main() {
	/*
		// 一次性读取:直接读入整个文件,简单但占用内存较大(适合小文件)
		// Read at once: loads the whole file into memory, simple but memory-heavy (fine for small files).
		byteData, err := os.ReadFile("Go_study_07_文件操作/Hello.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(byteData))
	*/

	/*
		// 分片读:指定缓冲区大小循环读取,适合大文件
		// Read by chunks: loop reading with a fixed buffer, suitable for large files.
		file, err := os.Open("Go_study_07_文件操作/Hello.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		for true {
			var byteData = make([]byte, 12)
			n, err := file.Read(byteData)
			if err == io.EOF { // 读到文件末尾,结束循环	Reached end of file, break the loop.
				break
			}
			fmt.Println(string(byteData), n)
		}
	*/

	// 带缓冲读
	// Buffered read.
	file, err := os.Open("Go_study_07_文件操作/Hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建缓冲器
	buf := bufio.NewReader(file)
	line, _, err := buf.ReadLine()
	fmt.Println(string(line), err)
}
