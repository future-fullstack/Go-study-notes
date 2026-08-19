package main

// main 展示读取文件的几种方式(一次性读取、分片读、带缓冲读)
// main demonstrates several ways to read a file (at once, by chunks, buffered).
func main() {
	/*
		// 一次性读取:直接读入整个文件,简单但占用内存较大(适合小文件)
		// Read at once: loads the whole file into memory, simple but memory-heavy (fine for small files).
		byteData, err := os.ReadFile("Go_study_07_文件操作/Hello.txt") // 读入整个文件内容	Read the whole file content
		if err != nil {
			fmt.Println(err) // 出错则打印错误并返回	Print the error and return
			return
		}
		fmt.Println(string(byteData)) // 转为字符串输出	Convert to string and print
	*/

	/*
		// 分片读:指定缓冲区大小循环读取,适合大文件
		// Read by chunks: loop reading with a fixed buffer, suitable for large files.
		file, err := os.Open("Go_study_07_文件操作/Hello.txt") // 打开文件	Open the file
		if err != nil {
			fmt.Println(err) // 出错则打印错误并返回	Print the error and return
			return
		}
		defer file.Close() // 延迟到函数结束再关闭文件	Close the file when the function returns
		for true {
			// 每次读取 12 字节到缓冲区
			// Read 12 bytes into the buffer each time
			var byteData = make([]byte, 12)
			n, err := file.Read(byteData) // 把一段内容读进缓冲区	Read a chunk into the buffer
			if err == io.EOF { // 读到文件末尾,结束循环	Reached end of file, break the loop.
				break
			}
			fmt.Println(string(byteData), n) // 打印本次读到的内容与字节数	Print the content and the byte count
		}
	*/

	/*
		// 带缓冲读
		// Buffered read.
		file, err = os.Open("Go_study_07_文件操作/Hello.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// 创建缓冲器
		// Create a buffered reader
		buf := bufio.NewReader(file)

		// 逐行读
		// Read line by line
		line, _, err := buf.ReadLine()
		fmt.Println(string(line), err)

		line, _, err = buf.ReadLine()
		fmt.Println(string(line), err)
	*/

	/*
		// 指定分隔符
		// Specify a delimiter
		file, err := os.Open("Go_study_07_文件操作/Hello.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// 使用 Scanner 按分隔符逐段读取
		// Use a Scanner to read segment by segment
		buf := bufio.NewScanner(file)
		buf.Split(bufio.ScanLines) // 设置按行分割	Split by lines
		for buf.Scan() {           // 循环扫描	Scan in a loop
			fmt.Println(buf.Text()) // 打印当前段	Print the current segment
		}
	*/
}
