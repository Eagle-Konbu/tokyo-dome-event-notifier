package slack

import (
	"fmt"
	"os"
)

func SendEventInfo(text string) {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	// body := fmt.Sprintf(`{
	// 	"blocks": [
	// 		{
	// 			"type": "header",
	// 			"text": {
	// 				"type": "plain_text",
	// 				"text": "本日のイベント情報"
	// 			}
	// 		},
	// 		{
	// 			"type": "divider"
	// 		},
	// 		{
	// 			"type": "section",
	// 			"text": {
	// 				"type": "plain_text",
	// 				"text": "%s",
	// 				"emoji": true
	// 			}
	// 		}
	// 	]
	// }`, text)

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	// req.Header.Set("Content-Type", "application/json")

	// if err != nil {
	// 	panic(err)
	// }

	// client := new(http.Client)
	// res, err := client.Do(req)

	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(url)
	// defer res.Body.Close()
}
