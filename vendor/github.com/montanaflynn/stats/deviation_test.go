package stats

import (
	"math"
	"testing"
)

func TestMedianAbsoluteDeviation(t *testing.T) {
	_, err := MedianAbsoluteDeviation([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestMedianAbsoluteDeviationPopulation(t *testing.T) {
	s, _ := MedianAbsoluteDeviation([]float64{1, 2, 3})
	m, err := Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.00 {
		t.Errorf("%.10f != %.10f", m, 1.00)
	}

	s, _ = MedianAbsoluteDeviation([]float64{-2, 0, 4, 5, 7})
	m, err = Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 3.00 {
		t.Errorf("%.10f != %.10f", m, 3.00)
	}

	m, _ = MedianAbsoluteDeviation([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestStandardDeviation(t *testing.T) {
	_, err := StandardDeviation([]float64{1, 2, 3})
	if err != nil {
		t.Errorf("Returned an error")
	}
}

func TestStandardDeviationPopulation(t *testing.T) {
	s, _ := StandardDeviationPopulation([]float64{1, 2, 3})
	m, err := Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.82 {
		t.Errorf("%.10f != %.10f", m, 0.82)
	}
	s, _ = StandardDeviationPopulation([]float64{-1, -2, -3.3})
	m, err = Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 0.94 {
		t.Errorf("%.10f != %.10f", m, 0.94)
	}

	m, _ = StandardDeviationPopulation([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}

func TestStandardDeviationSample(t *testing.T) {
	s, _ := StandardDeviationSample([]float64{1, 2, 3})
	m, err := Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.0 {
		t.Errorf("%.10f != %.10f", m, 1.0)
	}
	s, _ = StandardDeviationSample([]float64{-1, -2, -3.3})
	m, err = Round(s, 2)
	if err != nil {
		t.Errorf("Returned an error")
	}
	if m != 1.15 {
		t.Errorf("%.10f != %.10f", m, 1.15)
	}

	m, _ = StandardDeviationSample([]float64{})
	if !math.IsNaN(m) {
		t.Errorf("%.1f != %.1f", m, math.NaN())
	}
}
