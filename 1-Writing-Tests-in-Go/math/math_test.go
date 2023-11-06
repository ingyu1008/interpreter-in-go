package math

import "testing"

func TestGCD(t *testing.T) {
	inputs := []struct {
		a int
		b int
	}{
		{12, 60},
		{3, 5},
		{10, 66},
		{500, 75},
	}

	tests := []int{12, 1, 2, 25}

	for i, tt := range tests {
		ans := gcd(inputs[i].a, inputs[i].b)
		if ans != tt {
			t.Fatalf("tests[%d] - wrong answer. expected-%d, got-%d", i, tt, ans)
		}
	}

}
