package scrapy

import (
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const (
	WinGoogleUserAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
	AndroidUserAgent   string = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Mobile Safari/537.36 Edg/115.0.1901.203"
)

func GetAndSavePic(path, name, url string) error {
	rsp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(path+"/"+name, data, os.ModeAppend)
}

func GetWeb(url string) (*goquery.Document, error) {
	return GetWebWithHeader(url, nil)
}

func GetWebWithHeader(url string, header map[string]string) (*goquery.Document, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", WinGoogleUserAgent)
	// req.Header.Add("Sec-Ch-Ua-Platform", AndroidUserAgent)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
