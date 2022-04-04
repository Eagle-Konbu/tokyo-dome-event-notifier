package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchTodayEvent() string {
	time.LoadLocation("Asia/Tokyo")

	res, err := http.Get("https://www.tokyo-dome.co.jp/dome/event/schedule.html")

	if err != nil {
		fmt.Println("Failed to scrape")
		panic(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	selector := "div.c-mod-tab__body:nth-child(2) > table > tbody"
	innerSelector := "tr.c-mod-calender__item"
	dateSelector := "th > span:nth-child(1)"
	categorySelector := "td:nth-child(2) > div > div:nth-child(1) > p > span"
	titleSelector := "td > div > div:nth-child(2) > p.c-mod-calender__links"
	timeSelector := "td > div > div:nth-child(2) > p:nth-child(2)"

	selection := doc.Find(selector)

	var event string
	selection.Find(innerSelector).Each(func(index int, s *goquery.Selection) {
		date, _ := strconv.Atoi(s.Find(dateSelector).Text())
		category := s.Find(categorySelector).Text()
		title := s.Find(titleSelector).Text()
		info := s.Find(timeSelector).Text()

		if date == time.Now().Day() {
			if title == "" {
				event = "イベントなし"
			} else {
				event = title + "（" + category + "）" + "\n" + info
			}
		}
	})
	return event
}
