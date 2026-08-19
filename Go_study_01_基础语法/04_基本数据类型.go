package main

import "fmt"

func main() {
	//int: int8 int16 int32 int64   后缀数字代表该类型的位数，例如 int8 为 8 位有符号整数   有符号
	//int: int8 int16 int32 int64	The suffix number indicates the bit width,e.g.,int8 is an 8-bit signed integer
	//uint:uint8 uint16 uint32 uint64   ...   无符号整数
	//uint:uint8 uint16 uint32 uint64   ...   Unsigned integer

	//注意观察下面的输出差异
	//Observe the output difference below
	//解释：当输入单个字符（byte）时，系统会依据 ASCII 码表将其映射为对应的数值并存储
	//Explanation:When a single character(byte) is input,the system map it to its corresponding numerical  value based on ASCII table for storage
	var a1 byte = 'a'
	fmt.Println(a1)
	fmt.Printf("%c %d\n", a1, a1)
	//验证上述的数值映射关系
	//Verify the mapping relationship mentioned above
	var a2 = 97
	fmt.Printf("%c,%d\n", a2, a2)

	//多字节字符（Go 中 rune 类型本质是 int32，用于表示 Unicode 码点）
	//Multi-byte characters(In Go,rune is an alias for int32,used for Unicode code points)
	var b1 rune = '中'
	fmt.Printf("%c,%d\n", b1, b1)

	//转义字符 \ 告诉编译器：“后面的字符作为普通文本输出，而非代码指令”
	//The escape character \ tells the compiler:"Treat the following character as a plaint literal,not as code instructions"
	fmt.Println("hello")
	fmt.Println("'hel'lo")
	fmt.Println("hel\tlo")
	fmt.Println("\"hel\"lo")
	fmt.Println("\\hel\\lo")

	// 多行/原始字符串。使用反引号 ` 包裹的字符串会被原样输出（转义符无效，支持直接换行）
	//Multi-line/Raw string literals.Characters enclosed by backticks ` are output verbatim(escape sequences are ignored,supports direct line breaks)
	fmt.Println(`你好呀，王鸿铭\n
今天是7.1`)

	//如果我们给一个基本数据的变量只声明不赋值，那么这个变量就默认为是这个数据类型的零值(bool -> false)
	//If we declare a variable of a basic type without assigning a value,it defects to "zero value" of that type(e.g.,the zero value for bool is false)
	var c1 string
	var c2 bool
	fmt.Println(c1, c2)
}
