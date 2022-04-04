package slack

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func SendEventInfo(text string) {
	time.LoadLocation("Asia/Tokyo")
	t := time.Now()
	weekdayja := strings.NewReplacer(
		"Sun", "日",
		"Mon", "月",
		"Tue", "火",
		"Wed", "水",
		"Thu", "木",
		"Fri", "金",
		"Sat", "土",
	)
	date := weekdayja.Replace(t.Format("2006年1月2日(Mon曜日)"))

	url := os.Getenv("SLACK_WEBHOOK_URL")
	body := fmt.Sprintf(`{
		"text": "%sのイベント情報",
		"blocks": [
			{
				"type": "header",
				"text": {
					"type": "plain_text",
					"text": "%sのイベント情報"
				}
			},
			{
				"type": "divider"
			},
			{
				"type": "section",
				"text": {
					"type": "plain_text",
					"text": "%s",
					"emoji": true
				}
			}
		]
	}`, date, date, text)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := new(http.Client)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
