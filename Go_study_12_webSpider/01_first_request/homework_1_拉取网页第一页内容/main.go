package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://quotes.toscrape.com/" // 设置名言首页地址	Set the URL of the first quote page.

	req, err := http.NewRequest("GET", url, nil) // 创建不带请求体的 GET 请求	Create a GET request without a request body.
	if err != nil {                              // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("创建请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return                      // 结束当前函数, 不再继续处理	Exit the current function without further processing.
	}

	client := http.DefaultClient // 使用标准库默认 HTTP 客户端	Use the default HTTP client from the standard library.

	resp, err := client.Do(req) // 发送请求并接收 HTTP 响应	Send the request and receive the HTTP response.
	if err != nil {             // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("发送请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return                      // 结束当前函数, 不再继续处理	Exit the current function without further processing.
	}
	defer resp.Body.Close() // 在当前函数返回时关闭响应体	Close the response body when the current function returns.

	fmt.Println("response status", resp.Status) // 打印 HTTP 响应状态	Print the HTTP response status.

	doc, err := goquery.NewDocumentFromReader(resp.Body) // 从响应体读取 HTML 并构建可查询的文档树	Read HTML from the response body and build a queryable document tree.
	if err != nil {                                      // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("解析HTML失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
	}

	doc.Find("div.quote").Each(func(i int, s *goquery.Selection) { // 遍历页面中的名言容器	Iterate over the quote containers on the page.
		quote := s.Find("span.text").Text()                         // 提取当前名言的正文	Extract the current quote text.
		author := s.Find("small.author").Text()                     // 提取当前名言的作者	Extract the author of the current quote.
		var Tags []string                                           // 为当前名言准备标签切片	Prepare a tag slice for the current quote.
		fmt.Println(i, quote, author)                               // 打印索引、名言及作者	Print the index, quote, and author.
		s.Find("a.tag").Each(func(j int, tags *goquery.Selection) { // 遍历当前名言的标签链接	Iterate over tag links for the current quote.
			tag := tags.Text()       // 读取标签区域的文本	Read the text of the selected tag region.
			Tags = append(Tags, tag) // 将标签追加到当前名言的切片	Append the tag to the current quote's slice.
		})
		fmt.Println(Tags) // 输出当前名言的全部标签	Print all tags for the current quote.

	})
}
