package main

// main 展示写入文件的几种方式(先写后读、一次性写入、文件复制)与目录操作
// main demonstrates several ways to write a file (write-then-read, write at once, file copy) and directory ops.
func main() {
	/*
		// 先写后读:打开文件,写入内容,再把读取指针重置到开头读回
		// Write then read: open the file, write content, reset the read pointer and read it back.
		file, err := os.OpenFile("w1.txt", os.O_CREATE|os.O_RDWR, 0777) // 打开文件(不存在则创建,可读可写)	Open the file (create if absent, readable & writable)
		if err != nil {
			panic(err) // 出错直接终止	Panic on error
		}
		defer file.Close() // 延迟到函数结束再关闭文件	Close the file when the function returns

		file.Write([]byte("Hello\n")) // 写入字节内容	Write byte content

		_, err = file.Seek(0, 0) // 将读取指针重置到开头	Reset the read pointer to the beginning

		byteData, err := io.ReadAll(file) // 一次性读取全部剩余内容	Read all the remaining content at once
		if err != nil {
			fmt.Println(err) // 出错则打印错误并返回	Print the error and return
			return
		}
		fmt.Println(string(byteData)) // 转为字符串输出	Convert to string and print
	*/

	/*
		// 一次性写入:直接创建(或覆盖)文件并写入内容
		// Write at once: create (or overwrite) the file and write content directly.
		err := os.WriteFile("w1.txt", []byte("你好"), 0666) // 内容直接写入文件,权限 0666	Write the content to the file directly, permission 0666
		fmt.Print(err) // 打印错误(无错误时为 nil)	Print the error (nil when there is none)
	*/

	/*
		// 文件复制:打开源文件,创建目标文件,用 io.Copy 拷贝内容
		// File copy: open the source file, create the destination file, copy the content with io.Copy.
		rFile, err := os.Open("E:\\浏览器\\Folder-Ico-main\\Folder-Ico-main\\ico\\4k_downloader.ico") // 打开源文件	Open the source file
		if err != nil {
			fmt.Print(err)
			return
		}
		defer rFile.Close() // 延迟关闭源文件	Close the source file when the function returns

		wFile, err := os.OpenFile("4k.ico", os.O_CREATE|os.O_WRONLY, 0777) // 创建目标文件(只写)	Create the destination file (write-only)
		if err != nil {
			fmt.Print(err)
			return
		}
		defer wFile.Close() // 延迟关闭目标文件	Close the destination file when the function returns

		io.Copy(wFile, rFile) // 把源文件内容拷贝到目标文件	Copy the source file content into the destination file
	*/

	/*
		// 目录操作:读取目录下的所有条目并打印信息
		// Directory operations: read all entries in a directory and print their info.
		dir, err := os.ReadDir("Go_study_07_文件操作") // 读取目录下所有条目	Read all the entries in the directory
		if err != nil {
			fmt.Print(err)
			return
		}
		for _, entry := range dir { // 遍历每个条目	Iterate over each entry
			info, _ := entry.Info()                               // 获取条目的详细信息(如大小)	Get the entry's detailed info (e.g. size)
			fmt.Println(entry.IsDir(), entry.Name(), info.Size()) // 打印:是否目录、名称、大小	Print: is-dir, name, size
		}
	*/
}
