package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {

	// 创建 chrome 实例
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置 timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var example string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://pkg.go.dev/time"),
		chromedp.WaitVisible("body > footer"),
		chromedp.Click("#example-After", chromedp.NodeVisible),
		chromedp.Value("#example-After textarea"+
			"", &example),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Go's time.After example:\n%s", example)

	//fetcher := collect.BaseFetcher{}
	//
	//b, err := fetcher.Get(context.TODO(), "https://book.douban.com/subject/1007305/")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf(string(b))

	//doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//doc.Find("div.small_cardcontent__BTALp h2").Each(func(i int, selection *goquery.Selection) {
	//	log.Println(selection.Text())
	//})
}
