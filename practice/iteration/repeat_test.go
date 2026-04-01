package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat a character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		// The Repeat function is expected to repeat the character 'a' 5 times
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("repeat a character 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		// The Repeat function is expected to return an empty string when repeatCount is 0
		expected := ""
		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
