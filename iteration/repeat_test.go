package iteration

import (
	"fmt"
	"testing"
)

func TestReapeat(t *testing.T) {
	t.Run("5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("3 times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 4)
	fmt.Println(repeated)
	// Output: cccc
}
