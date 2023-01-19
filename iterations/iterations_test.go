package iterations

import "testing"

func TestRepeat(t *testing.T) {
  result := Repeat("a", 5)
  expected := "aaaaa"

  if result != expected {
    t.Errorf("Expected '%s', got '%s'", expected, result)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i:=0; i<b.N; i++ {
    Repeat("a", 5)
  }
}

