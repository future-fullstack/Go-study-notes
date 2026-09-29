package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	for i := 1; true; i++ { // 从第 1 页开始递增, 直到执行 break	Increment page numbers from one until a break is reached.
		url := "https://quotes.toscrape.com/page/" + strconv.Itoa(i) + "/" // 把页码转成字符串并拼接页面地址	Convert the page number to a string and construct the page URL.
		req, err := http.NewRequest("GET", url, nil)                       // 创建不带请求体的 GET 请求	Create a GET request without a request body.
		if err != nil {                                                    // 检查上一步操作是否失败	Check whether the previous operation failed.
			fmt.Println(err) // 打印错误以便定位失败原因	Print the error to help identify the failure.
			continue         // 跳过本轮, 继续尝试下一页	Skip this iteration and try the next page.
		}

		client := http.DefaultClient // 使用标准库默认 HTTP 客户端	Use the default HTTP client from the standard library.

		resp, err := client.Do(req) // 发送请求并接收 HTTP 响应	Send the request and receive the HTTP response.
		if err != nil {             // 检查上一步操作是否失败	Check whether the previous operation failed.
			fmt.Println(err) // 打印错误以便定位失败原因	Print the error to help identify the failure.
			continue         // 跳过本轮, 继续尝试下一页	Skip this iteration and try the next page.
		}
		resp.Body.Close() // 关闭响应体 (待修正: 此处早于下方读取)	Close the response body (TODO: this currently precedes the read below).

		doc, err := goquery.NewDocumentFromReader(resp.Body) // 从响应体读取 HTML 并构建可查询的文档树	Read HTML from the response body and build a queryable document tree.
		if err != nil {                                      // 检查上一步操作是否失败	Check whether the previous operation failed.
			fmt.Println(err) // 打印错误以便定位失败原因	Print the error to help identify the failure.
			continue         // 跳过本轮, 继续尝试下一页	Skip this iteration and try the next page.
		}

		fmt.Println(resp.Status)                    // 打印当前页的 HTTP 状态	Print the HTTP status of the current page.
		if resp.StatusCode == http.StatusNotFound { // 若执行到此处且状态为 404, 则结束翻页	Stop pagination if execution reaches here with status 404.
			break // 退出当前翻页循环	Exit the current pagination loop.
		}

		doc.Find("div.quote").Each(func(i int, s *goquery.Selection) { // 遍历页面中的名言容器	Iterate over the quote containers on the page.
			quote := s.Find("span.text").Text()                            // 提取当前名言的正文	Extract the current quote text.
			author := s.Find("small.author").Text()                        // 提取当前名言的作者	Extract the author of the current quote.
			fmt.Println(quote, author)                                     // 输出名言和作者	Print the quote and its author.
			s.Find("div.tags").Each(func(i int, tags *goquery.Selection) { // 读取整个标签容器, 文本可能包含栏目标题	Select the entire tag container, whose text may include its heading.
				tag := tags.Text() // 读取标签区域的文本	Read the text of the selected tag region.
				fmt.Println(tag)   // 输出标签区域的文本	Print the tag region text.
			})
		})

	}
}
