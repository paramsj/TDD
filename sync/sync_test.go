package sync

import (
	"sync"
	"testing"
)

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("COnkcurrency", func(t *testing.T) {
		wanted := 3000

		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(wanted)

		for i := 0; i < wanted; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wanted)
	})
}
