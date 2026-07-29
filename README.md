# Go Study Notes 🚀

> 我是 **耄耄**，目标是成为全栈工程师。
> 这个仓库是我在 GitHub 上的第一个仓库，存放着我的 Go 语言学习笔记，会长期更新 ✨

按模块分类整理。

## 目录结构

```
Go_study_foundation/    基础语法
├── 01_变量定义          变量与常量
├── 02_输出              输出
├── 03_输入              输入
├── 04_基本数据类型       基本数据类型
├── 05_数组              数组
├── 06_切片              切片
├── 07_map               Map
├── 08_if语句            if 条件判断
├── 09_switch            switch 语句
└── 10_for循环           for 循环

Go_study_function/      函数篇
├── 01_简单函数          基本函数
├── 02_闭包              闭包
├── 03_init函数          init 函数
└── 04_defer函数         defer 语句

Go_study_struct/        结构体与进阶
├── 01_对象              结构体定义
├── 02_tag               结构体标签 (Tag)
├── 03_组合              结构体组合
├── 04_结构体指针         指针与方法
├── 05_自定义类型         自定义类型
├── 06_自定义类型和类型别名 自定义类型 vs 类型别名
├── 07_接口              接口
├── 08_协程              Goroutine
├── 09_channel           Channel
└── 10_select            Select
```

## 注释风格

所有代码均采用 **中英双语注释**，方便理解和查阅：

```go
// 结构体标签的作用是：当转换成 JSON 时规定属性的键名
// Struct tags specify the key names when marshalling to JSON.
```

## 学习路线

1. 基础语法 → `Go_study_foundation/`
2. 函数进阶 → `Go_study_function/`
3. 结构体、接口、并发 → `Go_study_struct/`
