package stats

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3}, 6},
		{[]float64{1.0, 1.1, 1.2, 2.2}, 5.5},
		{[]float64{1, -1, 2, -3}, -1},
	} {
		got, err := Sum(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Sum(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Sum([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkSumSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(makeFloatSlice(5))
	}
}

func BenchmarkSumLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(lf)
	}
}
