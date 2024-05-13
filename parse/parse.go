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
