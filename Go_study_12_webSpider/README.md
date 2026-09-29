# Go 爬虫学习导航

这里按“请求网页 → 提取静态内容 → 翻页 → 判断动态数据来源 → 提取 script 数据”的顺序阅读。目录状态描述的是代码现状，不代表已经掌握或完成验收。

## 文件地图

| 文件 | 作用 | 当前注意点 |
| --- | --- | --- |
| [01_first_request/main.go](01_first_request/main.go) | 从图书页读取书名、价格、库存与星级 | 最初的 HTTP + goquery 示例 |
| [homework_1/main.go](01_first_request/homework_1_拉取网页第一页内容/main.go) | 从名言首页读取正文、作者和标签 | HTML 解析失败后缺少 return |
| [homework_2/main_1.go](01_first_request/homework_2_实现翻页/main_1.go) | 自增页码、拼接 URL 的翻页尝试 | 响应体提前关闭；停止条件需再检查 |
| [homework_2/main_2.go](01_first_request/homework_2_实现翻页/main_2.go) | 跟随下一页链接，拆分 send 与 spider | send 返回 nil 时，调用方尚未拦截 |
| [01_find_api/main.go](02_API_Scraping/01_find_api/main.go) | 数据来源侦察的学习占位 | 只有 package 声明，无 main 函数；未完成，不加注释 |
| [02_extract_script_data/main.go](02_API_Scraping/02_extract_script_data/main.go) | 开始请求动态页面并解析 HTML | 草稿：URL 含 view-source:，doc 尚未使用 |

两种翻页写法保留在原处，方便对照思路。其余目录和文件名也保持原样。

## 如何运行

在仓库根目录 D:\StartingGo 中，指定单个文件：

```powershell
go run "Go_study_12_webSpider/01_first_request/main.go"
go run "Go_study_12_webSpider/01_first_request/homework_1_拉取网页第一页内容/main.go"
go run "Go_study_12_webSpider/01_first_request/homework_2_实现翻页/main_2.go"
```

main_1.go 和 main_2.go 各自包含 main 函数，因此分别运行，不一起传给 go run。上面的命令会实际请求练习网站；静态检查通过不代表网络请求和异常分支已经验证。

## 接着上次的讨论

来源：[共享对话《Go爬虫》](https://chatgpt.com/share/6abb28d5-edf0-83ee-b7ef-79415494b82b)。

目前的讨论落在“从 script 提取嵌入数据”：先拿到 HTML，再定位包含 var data 的 script，分离其中的数据，最后尝试用 encoding/json 解析。HTML 解析器不会执行 JavaScript；JavaScript 赋值语句整体也不是 JSON，提取出的数据是否符合 JSON 语法仍需确认。

下一次从 02_extract_script_data/main.go 接着写，一次解决一个节点。已成形的静态页面练习补双语注释；02_API_Scraping 下的两个未完成文件保持无注释，不代写后续逻辑，也不标记学习完成。

## 已发现、留待一起处理的问题

- main_1.go 在读取响应体之前调用 Close；即使修正顺序，只靠 404 结束也不充分，应结合空列表或下一页链接判断。它在错误后继续增加页码，持续失败时可能不断请求。
- 首页作业解析失败后仍会继续使用 doc；main_2.go 也需要在 send 失败后处理 nil。
- script 草稿中的 view-source: 是浏览器查看源码的方式，不能作为 net/http 的请求协议；doc 未使用会导致编译失败。
- 请求示例尚未设置显式超时，HTTP 非成功状态也未统一处理。这些是后续完善项。
