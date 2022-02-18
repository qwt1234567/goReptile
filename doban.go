package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Spider()
}

func Spider() {
	// 1、发送请求
	// 创建客服端
	client := http.Client{}
	// 创建请求
	// req, err := http.NewRequest(method:"GET", url: "", body: nil)
	req, err := http.NewRequest("GET", "https://movie.douban.com/", nil)
	if err != nil {
		fmt.Println("req err", err)
	}

	// 添加请求头信息，防止服务器检测爬虫访问，伪造浏览器访问
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败", err)
	}

	defer resp.Body.Close()
	// 2、解析网页
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析失败", docDetail)
	}
	// 3、获取节点信息
	title := docDetail.Find("#content > div > div.article > div.gaia.gaia-lite.gaia-movie.slide-mode > div.list-wp > div > div.slide-container > div > div:nth-child(2) > a:nth-child(3) > p").Text()
	// 4、保存信息
	fmt.Println("title", title)
}

// #screening > div.screening-bd > ul > li:nth-child(22) > ul > li.title > a
