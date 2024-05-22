package parse

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

func internalHTTPGet(url string) (respBody io.ReadCloser, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		err = fmt.Errorf("HTTP status code: %d", res.StatusCode)
		return
	}
	if err != nil {
		return
	}

	respBody = res.Body
	return
}

func CheckUrl(str string) (err error) {
	if len(str) == 0 {
		err = fmt.Errorf("URL is empty")
		return
	}

	if len(str) > 2048 {
		err = fmt.Errorf("URL is too long")
		return
	}

	if str[0] < 'a' || str[0] > 'z' {
		err = fmt.Errorf("URL is not valid")
		return
	}
	return
}

func fixUrl(url string) (fixedUrl string) {
	// 不是 http 或 https 开头的 url，加上 https://
	if url[:4] != "http" {
		fixedUrl = "https://" + url
	} else {
		fixedUrl = url
	}
	return
}

func Scrape(url string) []string {
	var imageUrls []string

	respReadCloser, err := internalHTTPGet(url)
	defer respReadCloser.Close()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(respReadCloser)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "og:image" {
			if content, exists := s.Attr("content"); exists {
				imageUrls = append(imageUrls, content)
			}
		}
	})

	return imageUrls
}

func RoboParse(url string) (imageUrls []string, err error) {
	err = CheckUrl(url)
	if err != nil {
		return
	}

	fixedUrl := fixUrl(url)
	imageUrls = Scrape(fixedUrl)
	return
}
