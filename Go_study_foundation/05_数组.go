package main

import "fmt"

func main() {
	//数组的创建
	//Creation of an array
	var namelist [3]string = [3]string{"Hello", "World", "Hi"}
	fmt.Println(namelist)

	//数组的索引	Go语言支持正向索引，不支持负向索引
	//Array indexing:GO support forward indexing,not reverse indexing
	fmt.Println(namelist[0])

	fmt.Println(len(namelist)) //len函数能取得数组长度  len() gets the array length

	//数组的修改
	//Modifying an array
	namelist[0] = "hello"
	fmt.Println(namelist)

	//Go语言不支持已创建数组元素的添加
	//Go doesn't support adding elements to an already  created array
}
