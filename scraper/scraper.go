package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func FetchTodayEvent() (string, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	url := "https://www.tokyo-dome.co.jp/dome/event/schedule.html"

	c := colly.NewCollector()

	var err error
	c.OnError(func(_ *colly.Response, e error) {
		fmt.Println("Failed to scrape")
		err = e
	})

	selector := "div.c-mod-tab__body:nth-child(2) > table > tbody"
	innerSelector := "tr.c-mod-calender__item"
	dateSelector := "th > span:nth-child(1)"
	categorySelector := "td:nth-child(2) > div > div:nth-child(1) > p > span"
	titleSelector := "td > div > div:nth-child(2) > p.c-mod-calender__links"
	timeSelector := "td > div > div:nth-child(2) > p:nth-child(2)"

	var event string
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach(innerSelector, func(_ int, s *colly.HTMLElement) {
			date := s.ChildText(dateSelector)
			category := s.ChildText(categorySelector)
			title := s.ChildText(titleSelector)
			info := s.ChildText(timeSelector)
			today := time.Now().In(jst).Format("02")

			if date == today {
				if title == "" {
					event = "イベントなし"
				} else {
					event = title + "（" + category + "）" + "\n" + info
				}
			}
		})
	})

	c.Visit(url)

	if err != nil {
		return "", err
	}

	return event, nil
}
