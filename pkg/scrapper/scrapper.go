package scrapper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)
const(
	TitleNotFound = 1
	CompanyNotFound = 3
	PriceNotFound = 5
)

func ScrapeLink(link string)(Title string,Company string,Price float64,Status int){
		doc:=getDocument(link)
		s1:=0
		s2:=0
		s3:=0
		priceString,ok:=doc.Find(`div#cerberus-data-metrics`).Attr("data-asin-price")
		if !ok{
			log.Print("not ok")
		}
		doc.Find("div#centerCol").Each(func(i int, s *goquery.Selection) {
			Title=s.Find("span#productTitle").Text()
			Company=s.Find("a#bylineInfo").Text()
		})

	if Title==""{
			s1=TitleNotFound
		}else {
			Title = strings.TrimSpace(Title)
		}
		if Company==""{
			s2=CompanyNotFound
		}else {
			Company = strings.TrimSpace(Company)
		}
		if priceString==""{
			s3=PriceNotFound
		}else{
			temp:=strings.TrimSpace(priceString)
			Price,_=strconv.ParseFloat(temp,32)
		}
		Status=s1+s2+s3
	return
}
func getDocument(url string) *goquery.Document {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	return doc
}
