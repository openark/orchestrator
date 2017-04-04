package stats

import (
	"reflect"
	"testing"
)

func TestPercentile(t *testing.T) {
	m, _ := Percentile([]float64{43, 54, 56, 61, 62, 66}, 90)
	if m != 62.0 {
		t.Errorf("%.1f != %.1f", m, 62.0)
	}
	m, _ = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50)
	if m != 5.5 {
		t.Errorf("%.1f != %.1f", m, 5.5)
	}
	m, _ = Percentile([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 99.9)
	if m != 10.0 {
		t.Errorf("%.1f != %.1f", m, 10.0)
	}
	_, err := Percentile([]float64{}, 99.9)
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
	_, err = Percentile([]float64{1, 2, 3, 4, 5}, 0)
	if err == nil {
		t.Errorf("Zero percent didn't return an error")
	}
	_, err = Percentile([]float64{1, 2, 3, 4, 5}, 0.13)
	if err == nil {
		t.Errorf("Too low percent didn't return an error")
	}
}

func TestPercentileSortSideEffects(t *testing.T) {
	s := []float64{43, 54, 56, 44, 62, 66}
	a := []float64{43, 54, 56, 44, 62, 66}
	Percentile(s, 90)
	if !reflect.DeepEqual(s, a) {
		t.Errorf("%.1f != %.1f", s, a)
	}
}

func BenchmarkPercentileSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Percentile(makeFloatSlice(5), 50)
	}
}

func BenchmarkPercentileLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Percentile(lf, 50)
	}
}

func TestPercentileNearestRank(t *testing.T) {
	f1 := []float64{35, 20, 15, 40, 50}
	f2 := []float64{20, 6, 7, 8, 8, 10, 13, 15, 16, 3}
	f3 := makeFloatSlice(101)

	for _, c := range []struct {
		sample  []float64
		percent float64
		result  float64
	}{
		{f1, 30, 20},
		{f1, 40, 20},
		{f1, 50, 35},
		{f1, 75, 40},
		{f1, 95, 50},
		{f1, 99, 50},
		{f1, 99.9, 50},
		{f1, 100, 50},
		{f2, 25, 7},
		{f2, 50, 8},
		{f2, 75, 15},
		{f2, 100, 20},
		{f3, 1, 100},
		{f3, 99, 9900},
		{f3, 100, 10000},
	} {
		got, err := PercentileNearestRank(c.sample, c.percent)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}
		if got != c.result {
			t.Errorf("%v != %v", got, c.result)
		}
	}

	_, err := PercentileNearestRank([]float64{}, 50)
	if err == nil {
		t.Errorf("Should have returned an empty slice error")
	}

	_, err = PercentileNearestRank([]float64{1, 2, 3, 4, 5}, -0.01)
	if err == nil {
		t.Errorf("Should have returned an percentage must be above 0 error")
	}

	_, err = PercentileNearestRank([]float64{1, 2, 3, 4, 5}, 110)
	if err == nil {
		t.Errorf("Should have returned an percentage must not be above 100 error")
	}

}

func BenchmarkPercentileNearestRankSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PercentileNearestRank(makeFloatSlice(5), 50)
	}
}

func BenchmarkPercentileNearestRankLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PercentileNearestRank(lf, 50)
	}
}
