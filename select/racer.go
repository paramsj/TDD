package racer

import (
	"fmt"
	"net/http"
	"time"
)

func measureDuration(url string) time.Duration {
	start := time.Now()
	resp, err := http.Get(url)

	if err == nil {
		resp.Body.Close()
	}
	return time.Since(start)
}

var TimeoutDuration = 10 * time.Second

func Racer(slowURL string, fastURL string) (string, error) {
	return ConfigurableRacer(slowURL, fastURL, TimeoutDuration)
}

func ConfigurableRacer(slowURL string, fastURL string, timeout time.Duration) (string, error) {
	select {
	case <-ping(slowURL):
		return slowURL, nil
	case <-ping(fastURL):
		return fastURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", slowURL, fastURL)
	}
}

func ping(slowURL string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(slowURL)
		if err == nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}
