package stats

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	for _, c := range []struct {
		number   float64
		decimals int
		result   float64
	}{
		{0.1111, 1, 0.1},
		{-0.1111, 2, -0.11},
		{5.3253, 3, 5.325},
		{5.3258, 3, 5.326},
		{5.3253, 0, 5.0},
		{5.55, 1, 5.6},
	} {
		m, err := Round(c.number, c.decimals)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if m != c.result {
			t.Errorf("%.1f != %.1f", m, c.result)
		}

	}
	_, err := Round(math.NaN(), 2)
	if err == nil {
		t.Errorf("Round should error on NaN")
	}
}

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(0.1111, 1)
	}
}
