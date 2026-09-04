package main

import (
	"fmt"
	"reflect"
)

// User1 携带一个方法 Call,演示用反射找到并调用它
// User1 carries a method Call, used to demonstrate finding and invoking it via reflection.
type User1 struct {
}

// Call 是 User1 的方法,接收一个字符串参数
// Call is a method of User1, taking a string argument.
func (User1) Call(name string) {
	fmt.Println("我被调用了", name)
}

// CallMethod 是包级函数:用反射遍历结构体的方法,找到名字为 "Call" 的那个并调用
// CallMethod is a package-level function: it uses reflection to enumerate a struct's methods, find the one named "Call", and invoke it.
func CallMethod(obj any) {
	v := reflect.ValueOf(obj).Elem()     // 拿到指针指向实体的 Value,Method(i) 需要从它上面取	Get the Value of the pointed-to entity; Method(i) is taken from it.
	t := reflect.TypeOf(obj).Elem()      // 对应的类型,Method(i) 提供方法的元信息	The matching type; Method(i) provides the method's metadata.
	for i := 0; i < v.NumMethod(); i++ { // NumMethod() 得到结构体的方法数量	NumMethod() returns the number of methods on the struct.
		m := t.Method(i)      // 每个方法的描述(含名字 Name)	Each method's description (includes its Name).
		if m.Name != "Call" { // 只处理名为 Call 的方法	Only handle the method named "Call".
			continue
		}
		method := v.Method(i)        // 拿到可调用的方法值	Obtain the callable method value.
		method.Call([]reflect.Value{ // Call 用 []reflect.Value 传入实参并真正执行该方法	Call passes arguments as []reflect.Value and really executes the method.
			reflect.ValueOf("WHM"),
		})
	}

}

func main() {
	s := User1{}
	CallMethod(&s) // 传指针才能通过 Elem() 访问到方法	Pass a pointer so Elem() can reach the methods.
	fmt.Println(s)
}
