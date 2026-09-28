package racer

import (
	"net/http"
	"time"
)

func Racer(slowURL string, fastURL string) string {
	durationA := measureDuration(slowURL)
	durationB := measureDuration(fastURL)

	if durationA < durationB {
		return slowURL
	}
	return fastURL
}

func measureDuration(url string) time.Duration {
	start := time.Now()
	resp, err := http.Get(url)

	if err == nil {
		resp.Body.Close()
	}
	return time.Since(start)
}
