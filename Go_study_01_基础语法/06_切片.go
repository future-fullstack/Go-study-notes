package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的创建(类似于python的列表)
	//Slice creation(similar to python lists)
	//这里定义了一个空数组，fmt.Println(namelist == nil) 可以验证
	//This defines an empty slice.Verify via fmt.Println(namelist == nil)
	var namelist1 []string //只能是同类型的
	namelist1 = append(namelist1, "张三")
	namelist1 = append(namelist1, "李四")
	namelist1 = append(namelist1, "王五")
	fmt.Println(namelist1)

	//上述切片可以像列表一样索引
	//This slice above can be indexed like a list

	//为了避免生成的是nil切片，我们可以使用如下方法进行初始化
	//To avoid generating a nil slice,we can initialize it using the following methods
	var namelist2 []string = []string{}
	fmt.Println(namelist2 == nil)

	namelist3 := []string{}
	fmt.Println(namelist3 == nil)

	var namelist4 []string
	namelist4 = make([]string, 0)
	fmt.Println(namelist4 == nil)

	namelist5 := make([]string, 0) //这个也挺省事的
	fmt.Println(namelist5 == nil)

	//创建全零切片(利用了基本数据类型不赋值就默认为该数据类型的零值的特性)
	//Create an all-zero slice(take the advantage of the fact that unsigning basic types default to their zero value)
	namelist6 := make([]int, 3)
	fmt.Println(namelist6)

	//命名为切片的原因
	//the reason it's called "slice"
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[:] //从头切到尾
	fmt.Println(slice)

	//切片排序
	//Sorting slice
	var ints = []int{4, 8, 3, 2}
	sort.Ints(ints) // 升序，这个是简化排序写法	ascending order,this is the simplified sort syntax
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序，这个是标准化写法	descending order,this is standardized syntax
	fmt.Println(ints)
}
