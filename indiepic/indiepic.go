package main


import (
	"fmt"
	"github.com/anjunact/learn-go/indiepic/crawldata"
)

func main () {
	// 使用crawldata包里面的Crawl()抓取需要的数据存到数据库
	crawldata.Crawl()
	fmt.Println("主函数")
}