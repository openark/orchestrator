package stats

import (
	"testing"
)

func TestQuartile(t *testing.T) {
	s1 := []float64{6, 7, 15, 36, 39, 40, 41, 42, 43, 47, 49}
	s2 := []float64{7, 15, 36, 39, 40, 41}

	for _, c := range []struct {
		in []float64
		Q1 float64
		Q2 float64
		Q3 float64
	}{
		{s1, 15, 40, 43},
		{s2, 15, 37.5, 40},
	} {
		quartiles, err := Quartile(c.in)
		if err != nil {
			t.Errorf("Should not have returned an error")
		}

		if quartiles.Q1 != c.Q1 {
			t.Errorf("Q1 %v != %v", quartiles.Q1, c.Q1)
		}
		if quartiles.Q2 != c.Q2 {
			t.Errorf("Q2 %v != %v", quartiles.Q2, c.Q2)
		}
		if quartiles.Q3 != c.Q3 {
			t.Errorf("Q3 %v != %v", quartiles.Q3, c.Q3)
		}
	}

	_, err := Quartile([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestInterQuartileRange(t *testing.T) {
	s1 := []float64{102, 104, 105, 107, 108, 109, 110, 112, 115, 116, 118}
	iqr, _ := InterQuartileRange(s1)

	if iqr != 10 {
		t.Errorf("IQR %v != 10", iqr)
	}

	_, err := InterQuartileRange([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestMidhinge(t *testing.T) {
	s1 := []float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13}
	mh, _ := Midhinge(s1)

	if mh != 7.5 {
		t.Errorf("Midhinge %v != 7.5", mh)
	}

	_, err := Midhinge([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestTrimean(t *testing.T) {
	s1 := []float64{1, 3, 4, 4, 6, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 10, 11, 12, 13}
	tr, _ := Trimean(s1)

	if tr != 7.25 {
		t.Errorf("Trimean %v != 7.25", tr)
	}

	_, err := Trimean([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
