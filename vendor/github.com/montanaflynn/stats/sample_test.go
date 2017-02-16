package stats

import (
	"testing"
)

func TestSample(t *testing.T) {
	_, err := Sample([]float64{}, 10, false)
	if err == nil {
		t.Errorf("Returned an error")
	}

	_, err2 := Sample([]float64{0.1, 0.2}, 10, false)
	if err2 == nil {
		t.Errorf("Returned an error")
	}
}

func TestSampleWithoutReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result, _ := Sample(arr, 5, false)
	checks := map[float64]bool{}
	for _, res := range result {
		_, ok := checks[res]
		if ok {
			t.Errorf("%v already seen", res)
		}
		checks[res] = true
	}
}

func TestSampleWithReplacement(t *testing.T) {
	arr := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	numsamples := 100
	result, _ := Sample(arr, numsamples, true)
	if len(result) != numsamples {
		t.Errorf("%v != %v", len(result), numsamples)
	}
}
