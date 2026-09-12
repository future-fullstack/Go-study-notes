# StartingGo ⚡

> 从 `Hello World` 出发，写一点真正能跑起来的 Go。

这是 **耄耄** 的 Go 学习实验场，也是我的第一个 GitHub 仓库。
这里没有“看完就会”的速成神话，只有一段持续提交、持续踩坑、持续变强的过程。

```text
基础语法 ──▶ 数据结构 ──▶ 并发编程 ──▶ 文件与测试
                         │
                         └──▶ 反射 ──▶ 网络编程 ──▶ Web Spider 🚧
```

## 当前进度

| 阶段 | 内容 | 状态 |
| :--- | :--- | :---: |
| 01 | 基础语法 | ✅ |
| 02 | 结构体与接口 | ✅ |
| 03 | 函数、闭包与 defer | ✅ |
| 04 | Goroutine、Channel 与线程安全 | ✅ |
| 05 | 异常处理 | ✅ |
| 06 | 泛型 | ✅ |
| 07 | 文件读写 | ✅ |
| 08 | 单元测试 | ✅ |
| 09 | 反射与 ORM 案例 | ✅ |
| 10 | TCP / HTTP 网络编程 | ✅ |
| 11 | Build 与可执行程序 | ✅ |
| 12 | Web Spider / Gin | 🚧 施工中 |

## 目录导航

```text
StartingGo/
├── Go_study_01_基础语法       # 变量、类型、数组、切片、map、流程控制
├── Go_study_02_结构体         # 结构体、组合、指针、自定义类型、接口
├── Go_study_03_函数           # 函数、闭包、init、defer
├── Go_study_04_协程           # Goroutine、Channel、Select、并发安全
├── Go_study_05_异常处理       # panic、recover 与程序中断
├── Go_study_06_泛型           # 泛型函数、结构体、切片与 map
├── Go_study_07_文件操作       # 文件读取与写入
├── Go_study_08_unitTest       # Go 单元测试入门
├── Go_study_09_反射           # reflect 与简单 ORM 思路
├── Go_study_10_网络编程       # TCP 服务端/客户端、HTTP 服务端/客户端
├── Go_study_11_Build          # 编译与构建
├── Go_study_12_webSpider      # Web Spider 学习区（未完成）
└── Whm's Wonder Lab           # 随手玩的小实验（已 gitignore，不纳入主线）
```

## 代码风格

学习代码统一使用 **中文 + English 双语注释**：

```go
fmt.Println("hello world") // 输出一行测试文本	Print one line of test text.
```

注释的目标不是把代码盖住，而是记录“这一行为什么存在”。有些实验代码并不完美，正因为如此，它们才保留了学习过程的痕迹。

## 快速开始

要求：

- Go `1.26+`
- 建议使用 GoLand 或 VS Code

克隆并进入项目：

```bash
git clone <your-repository-url>
cd StartingGo
go mod download
```

大多数示例都是独立文件，可以直接运行：

```bash
go run "Go_study_01_基础语法/01_变量定义.go"
go run "Go_study_10_网络编程/02_HTTP/01_HTTP服务端.go"
```

运行 HTTP 客户端前，先启动服务端；网络编程示例涉及端口占用，请确保 `127.0.0.1:801` 没有被其他程序使用。

## 我想在这里留下什么

```text
不是收藏一堆教程，
而是把每一个“我好像懂了”
都变成一段可以运行、可以解释、可以复盘的代码。
```

下一站：把 Web Spider 做完，再把这些零散的实验串成真正的小项目。

`Whm's Wonder Lab` 是我临时测试想法、练手和放飞脑洞的地方，已经加入 `.gitignore`，不会混入正式学习记录。

如果你也在学 Go，欢迎一起写、一起改、一起发现 bug。🚀
