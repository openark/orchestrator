package discovery

import (
	"github.com/montanaflynn/stats"
)

// internal routine to return the average value or 0
func mean(values stats.Float64Data) float64 {
	s, err := stats.Mean(values)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the requested percentile value or 0
func percentile(values stats.Float64Data, percent float64) float64 {
	s, err := stats.Percentile(values, percent)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the maximum value or 0
func max(values stats.Float64Data) float64 {
	s, err := stats.Max(values)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the minimum value or 9e9
func min(values stats.Float64Data) float64 {
	s, err := stats.Min(values)
	if err != nil {
		return 9e9 // a large number (should use something better than this but it's ok for now)
	}
	return s
}

// internal routine to return the median or 0
func median(values stats.Float64Data) float64 {
	s, err := stats.Median(values)
	if err != nil {
		return 0
	}
	return s
}
