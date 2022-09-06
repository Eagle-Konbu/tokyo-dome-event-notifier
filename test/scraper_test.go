package scraper

import (
	"testing"

	"github.com/tokyo-dome-event-notifier/scraper"
)

func TestFetchTodayEvent(t *testing.T) {
	res, err := scraper.FetchTodayEvent()

	if err != nil {
		t.Errorf("Failed to fetch today's event")
	} else if res == "" {
		t.Errorf("Today's event not found")
	} else {
		t.Logf("Today's event is \"%s\"", res)
	}
}
