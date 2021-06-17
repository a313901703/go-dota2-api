package spider

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//Spider struct
type Spider struct {
	URL    string
	Header map[string]string
}

//GetDoc return a Doc of pageview
func (spider Spider) GetDoc() *goquery.Document {
	// Request the HTML page.
	client := &http.Client{}
	req, err := http.NewRequest("GET", spider.URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(spider.Header) > 0 {
		for inx, header := range spider.Header {
			req.Header.Set(inx, header)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client connect error %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//fmt.Println(res.Body)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
