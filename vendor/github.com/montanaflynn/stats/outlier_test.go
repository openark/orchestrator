package stats

import (
	"testing"
)

func TestQuartileOutliers(t *testing.T) {
	s1 := []float64{-1000, 1, 3, 4, 4, 6, 6, 6, 6, 7, 8, 15, 18, 100}
	o, _ := QuartileOutliers(s1)

	if o.Mild[0] != 15 {
		t.Errorf("First Mild Outlier %v != 15", o.Mild[0])
	}

	if o.Mild[1] != 18 {
		t.Errorf("Second Mild Outlier %v != 18", o.Mild[1])
	}

	if o.Extreme[0] != -1000 {
		t.Errorf("First Extreme Outlier %v != -1000", o.Extreme[0])
	}

	if o.Extreme[1] != 100 {
		t.Errorf("Second Extreme Outlier %v != 100", o.Extreme[1])
	}

	_, err := QuartileOutliers([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
