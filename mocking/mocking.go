package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// write implements io.Writer and is why coundown can allow spycountdown as an io.Writer
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}

	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
