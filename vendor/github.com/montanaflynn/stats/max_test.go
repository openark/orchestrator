package stats

import (
	"testing"
)

func TestMax(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 5.0},
		{[]float64{10.5, 3, 5, 7, 9}, 10.5},
		{[]float64{-20, -1, -5.5}, -1.0},
		{[]float64{-1.0}, -1.0},
	} {
		got, err := Max(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Max(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Max([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func BenchmarkMaxSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(makeFloatSlice(5))
	}
}

func BenchmarkMaxLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(lf)
	}
}
