package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{0}
	Countdown(buffer, spySleeper)

	got := buffer.String()

	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enuf calls made to sleep %d calls were made", spySleeper.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
