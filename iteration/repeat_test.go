package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

// okay so you cannot put "aaaaa", you req it just as is
func ExampleRepeat() {
	concat := Repeat("a")
	fmt.Println(concat)
	// Output: aaaaa
}

// need to look into
