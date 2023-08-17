package scrapy

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/PuerkitoBio/goquery"
)

func SavePic(name, url string) {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.WriteFile("./pic/"+name+".jpg", data, os.ModeAppend)
}

func GetWeb(url string) *goquery.Document {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return doc
}

func Read() {
	prefix := "https://pic.netbian.com"
	for page := 1; page <= 70; page++ {
		url := "https://pic.netbian.com/4kmeinv/"
		if page > 1 {
			url = url + fmt.Sprintf("index_%d.html", page)
		}
		rsp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		doc, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 1; i <= 20; i++ {
			child := doc.Find(fmt.Sprintf("#main > div.slist > ul > li:nth-child(%d) > a", i))
			link, exist := child.Attr("href")
			if exist {
				img := GetWeb(prefix + link).Find("#img > img")
				imgLink, _ := img.Attr("src")
				title, exist := img.Attr("title")
				if exist {
					name, err := io.ReadAll(transform.NewReader(bytes.NewReader([]byte(title)), simplifiedchinese.GBK.NewDecoder()))
					if err != nil {
						fmt.Println(err, imgLink, name)
						break
					}
					SavePic(string(name), prefix+imgLink)
				}
			}
		}
		time.Sleep(time.Second)
	}
}
