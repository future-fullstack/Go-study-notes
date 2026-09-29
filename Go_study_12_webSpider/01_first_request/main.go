package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://books.toscrape.com/" // 设置图书练习页地址	Set the URL of the book practice page.

	req, err := http.NewRequest("GET", url, nil) // 创建不带请求体的 GET 请求	Create a GET request without a request body.
	if err != nil {                              // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("创建请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return                      // 结束当前函数, 不再继续处理	Exit the current function without further processing.
	}

	client := http.DefaultClient // 使用标准库默认 HTTP 客户端	Use the default HTTP client from the standard library.

	resp, err := client.Do(req) // 发送请求并接收 HTTP 响应	Send the request and receive the HTTP response.
	if err != nil {             // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("发出请求失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return                      // 结束当前函数, 不再继续处理	Exit the current function without further processing.
	}
	defer resp.Body.Close() // 在当前函数返回时关闭响应体	Close the response body when the current function returns.

	fmt.Println("response Status:", resp.Status) // 打印 HTTP 响应状态	Print the HTTP response status.

	doc, err := goquery.NewDocumentFromReader(resp.Body) // 从响应体读取 HTML 并构建可查询的文档树	Read HTML from the response body and build a queryable document tree.
	if err != nil {                                      // 检查上一步操作是否失败	Check whether the previous operation failed.
		fmt.Println("解析HTML失败：", err) // 输出当前操作的错误信息	Print the error from the current operation.
		return                        // 结束当前函数, 不再继续处理	Exit the current function without further processing.
	}

	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) { // 遍历每个图书卡片, i 从 0 开始	Iterate over book cards with a zero-based index.
		title := s.Find("h3 a").AttrOr("title", "") // 读取链接的完整书名属性, 缺失时用空字符串	Read the full title attribute or use an empty string if absent.
		piece := s.Find("p.price_color").Text()     // 读取价格文本, 当前变量名为 piece	Read the price text into the variable currently named piece.
		stock := strings.TrimSpace(                 // 去掉库存文本首尾的空白字符	Trim whitespace from both ends of the availability text.
			s.Find("p.instock.availability").Text(), // 读取当前图书卡片中的库存文本	Read the availability text in the current book card.
		)
		ratingClass := s.Find("p.star-rating").AttrOr("class", "") // 读取星级元素的 class 属性	Read the class attribute of the rating element.
		rating := strings.TrimPrefix(ratingClass, "star-rating ")  // 去掉公共前缀, 留下英文星级	Remove the common prefix to obtain the rating word.
		fmt.Println(i, title, piece, stock, rating)                // 输出索引及图书信息	Print the index and book details.
	})
}
