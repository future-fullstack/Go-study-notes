package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{
		"James": 32,
		"Mike":  23,
		"John":  24,
		"donk":  0,
	}
	fmt.Printf("%T\n", m)
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Whm"]) // 如果索引不存在的键，那么输出对应值的数据类型的'零值'	 indexing a non-existent key returns zero value of its data type

	//区分是否被索引的key在map中
	//Check whether the indexed key exists in the map
	value, ok := m["donk"] //value接收该键对应的值，ok接收对于是否存在该键的结果	"value" receives the value for the key,"ok" report whether the key exists
	fmt.Println(value, ok)

	//map中key的重新赋值
	//update the key's value
	m["James"] = 26
	fmt.Println(m["James"])

	//删除map中的键值
	//delect the key-value pair from the map
	delete(m, "donk")
	fmt.Println(m)

	//Map 变量必须经过初始化，否则无法向其添加键值对
	//A map variable must be initialized before adding key-value pairs
	Amap := map[string]int{} //或者是 Amap := make(map[string]int)	or Amap := make(map[string]int)
	fmt.Println(Amap)
}
