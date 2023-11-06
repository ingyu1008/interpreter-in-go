package reader

import (
	"testing"
)

// Always write test first, then implement
func TestRead(t *testing.T) {
	input := `1234abcd`
	tests := []struct {
		idx int
		exp int
	}{
		{0, 49},
		{1, 50},
		{2, 51},
		{3, 52},
		{4, 97},
		{5, 98},
		{6, 99},
		{7, 100},
		{-1, -1},
		{50, -1},
	}

	r := New(input)

	for i, tt := range tests {
		x := r.ReadPos(tt.idx)
		if x != tt.exp {
			t.Fatalf("tests[%d] - wrong answer. expected=%d, got=%d", i, tt.exp, x)
		}
	}
}
