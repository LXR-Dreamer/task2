package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://blog.lenconda.top/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	f, err := os.Create("C:/Users/Lenovo/Desktop/任务2/获取.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find(".post-box").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())
		fmt.Println(selection.Text())
	})

}
