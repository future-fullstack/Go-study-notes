package main

// Number 是一个自定义的类型约束(interface):它限定了泛型类型参数必须是下列这些整数类型中的一种
// Number is a custom type constraint (interface): it restricts generic type parameters to one of these integer types.
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// plus 是一个泛型函数:T 被约束为 Number,因此可以同时适配多种整数类型进行加法运算
// plus is a generic function: T is constrained to Number, so it works with many integer types at once.
func plus[T Number](n1, n2 T) T {
	return n1 + n2
}

// MyPrint 展示了多个类型参数的写法:一个类型参数对应一种约束
// MyPrint demonstrates multiple type parameter: each type parameter has its own constraint.
func MyPrint[T int, K string | int](n1 T, n2 K) {

}

func main() {
	var u1, u2 = uint(1), uint(2) // 使用 uint 类型调用泛型函数,无需类型转换	Call the generic function with uint, no type conversion needed.
	plus(u1, u2)                  // 编译器会自动推导 T 为 uint	The compiler infers T as uint automatically.
}
