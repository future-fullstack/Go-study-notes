package main

type MyCode int        // 自定义类型
type MyAliasCode = int // 类型别名

// 自定义类型 与 类型别名 的差异：
// 1. 类型别名不能绑定方法，自定义类型可以
// 2. 类型别名打印类型还是原始类型，但是自定义类型就是 main.xx 格式
// 3. 在需要进行条件判断时，类型别名不需要进行类型转换，自定义类型需要

// 自定义类型的作用：
// 1. 扩展功能：可以为基础类型绑定专属方法（Method）
// 2. 类型安全：实现类型隔离（如区分 UserID 和 OrderID），防止误用与传错参数
// 3. 增强语义：提高代码可读性，明确变量的实际业务含义

// 类型别名的作用：
// 1. 无缝重构：大项目中迁移包或重构代码时，零成本兼容旧代码
// 2. 免费继承：搭顺风车，直接继承原类型的所有方法，且与原类型比较/计算无需强转
// 3. 简化命名：给冗长或复杂的类型起个简短好记的名字（如官方的 byte = uint8, rune = int32）

// 下面仅仅对第三点差异进行讨论
const mycode MyCode = 1
const myaliascode MyAliasCode = 1

func main() {
	age := 1 // 一个 int 类型的普通变量	An ordinary int variable.

	// 自定义类型
	// Custom type.
	// 不报错:自定义类型与整型字面量比较,无需转换(字面量会隐式匹配)
	// No error: comparing a custom type with an integer literal needs no conversion (literals match implicitly).
	if mycode == 1 {
		return
	}
	// 报错,需要类型转换才能比较:自定义类型与 int 变量是不同静态类型
	// Error, needs a type conversion: a custom type and an int variable are distinct static types.
	if mycode == age {
		return
	}

	// 类型别名
	// Type alias.
	// 都不报错:类型别名本质就是原类型(int),与 int 变量可直接比较
	// No error in either: a type alias is essentially the original type (int), comparable with int directly.
	if myaliascode == 1 {
		return
	}
	if myaliascode == age {
		return
	}

}
