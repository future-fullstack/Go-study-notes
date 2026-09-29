package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// spider 输出名言并通过指针更新下一页地址, 返回下一页链接
// spider prints quotes, updates the next-page URL through a pointer, and returns its link.
func spider(doc *goquery.Document, url *string) string {
	doc.Find("div.quote").Each(func(i int, div *goquery.Selection) { // 逐条处理文档中的名言	Process each quote in the document.
		var ls []string                                              // 准备当前名言的输出字段	Prepare the output fields for the current quote.
		quote := div.Find("span.text").Text()                        // 读取名言正文	Read the quote text.
		author := div.Find("small.author").Text()                    // 读取作者名称	Read the author's name.
		ls = append(ls, quote, author)                               // 按顺序保存名言和作者	Append the quote and author in order.
		var tagList []string                                         // 为当前名言准备标签列表	Prepare a tag list for the current quote.
		div.Find("a.tag").Each(func(i int, tag *goquery.Selection) { // 遍历当前名言的各个标签	Iterate over individual tags for the current quote.
			tagList = append(tagList, tag.Text()) // 把当前标签文本加入列表	Append the current tag text to the list.
		})
		ls = append(ls, strings.Join(tagList, " ")) // 把标签用空格连接后加入输出字段	Join tags with spaces and append them to the output fields.
		fmt.Println(i+1, ls)                        // 使用从 1 开始的序号打印当前名言	Print the current quote with a one-based index.
	})
	href := doc.Find("li.next a").AttrOr("href", "") // 获取下一页相对链接, 不存在则返回空字符串	Read the relative next-page link or use an empty string if absent.
	if href == "" {                                  // 检查是否已没有下一页	Check whether there is no next page.
		return "" // 返回空字符串作为翻页结束信号	Return an empty string to signal the end of pagination.
	}
	fmt.Println("-------本页打印完毕-------")         // 输出本页与下一页之间的分隔提示	Print a separator between pages.
	*url = "https://quotes.toscrape.com" + href // 通过指针更新调用方 URL, 此处针对该站的根相对链接	Update the caller's URL through the pointer for this site's root-relative links.
	return href                                 // 返回下一页链接供调用方判断	Return the next-page link for the caller to inspect.
}

// send 请求页面并解析为 HTML 文档, 失败时返回 nil
// send fetches a page and parses an HTML document, returning nil on failure.
func send(url string) *goquery.Document {
	req, err := http.NewRequest("GET", url, nil) // 创建不带请求体的 GET 请求	Create a GET request without a request body.
	if err != nil {                              // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("创建请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return nil                  // 失败时返回空文档指针	Return a nil document pointer on failure.
	}

	client := http.DefaultClient // 使用标准库默认 HTTP 客户端	Use the default HTTP client from the standard library.

	resp, err := client.Do(req) // 发送请求并接收 HTTP 响应	Send the request and receive the HTTP response.
	if err != nil {             // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("发出请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return nil                  // 失败时返回空文档指针	Return a nil document pointer on failure.
	}
	defer resp.Body.Close() // 在当前函数返回时关闭响应体	Close the response body when the current function returns.

	fmt.Println("response Status:", resp.Status)         // 打印 HTTP 响应状态	Print the HTTP response status.
	doc, err := goquery.NewDocumentFromReader(resp.Body) // 从响应体读取 HTML 并构建可查询的文档树	Read HTML from the response body and build a queryable document tree.
	if err != nil {                                      // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println(err) // 打印错误以便定位失败原因	Print the error to help identify the failure.
		return nil       // 失败时返回空文档指针	Return a nil document pointer on failure.
	}
	return doc // 返回解析后的 HTML 文档	Return the parsed HTML document.
}

func main() {
	var url = "https://quotes.toscrape.com" // 设置翻页请求的起始地址	Set the initial URL for pagination.
	for {                                   // 持续翻页, 直到循环内主动退出	Keep paginating until an explicit loop exit.
		doc := send(url)          // 请求并解析当前页 (待补充 nil 检查)	Fetch and parse the current page (TODO: check for nil).
		href := spider(doc, &url) // 提取本页内容并传入 URL 地址以更新下一页	Extract page data and pass the URL address to update the next page.
		if href == "" {           // 检查是否已没有下一页	Check whether there is no next page.
			break // 退出当前翻页循环	Exit the current pagination loop.
		}
	}
}
