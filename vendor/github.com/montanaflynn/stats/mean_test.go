package stats

import (
	"testing"
)

func TestMean(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := Mean(c.in)
		if got != c.out {
			t.Errorf("Mean(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Mean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func BenchmarkMeanSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(makeFloatSlice(5))
	}
}

func BenchmarkMeanLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mean(lf)
	}
}

func TestGeometricMean(t *testing.T) {
	s1 := []float64{2, 18}
	s2 := []float64{10, 51.2, 8}
	s3 := []float64{1, 3, 9, 27, 81}

	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{s1, 6},
		{s2, 16},
		{s3, 9},
	} {
		gm, err := GeometricMean(c.in)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}

		gm, _ = Round(gm, 0)
		if gm != c.out {
			t.Errorf("Geometric Mean %v != %v", gm, c.out)
		}
	}

	_, err := GeometricMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestHarmonicMean(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	s2 := []float64{10, -51.2, 8}
	s3 := []float64{1, 0, 9, 27, 81}

	hm, err := HarmonicMean(s1)
	if err != nil {
		t.Errorf("Should not have returned an error")
	}

	hm, _ = Round(hm, 2)
	if hm != 2.19 {
		t.Errorf("Geometric Mean %v != %v", hm, 2.19)
	}

	hm, err = HarmonicMean(s2)
	if err == nil {
		t.Errorf("Should have returned a negative number error")
	}

	hm, err = HarmonicMean(s3)
	if err == nil {
		t.Errorf("Should have returned a zero number error")
	}

	_, err = HarmonicMean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
