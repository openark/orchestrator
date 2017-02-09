package stats

import (
	"testing"
	"time"
)

var allTestData = []struct {
	actual   interface{}
	expected Float64Data
}{
	{
		[]interface{}{1.0, "2", 3.0, 4, "4.0", 5, time.Duration(6), time.Duration(-7)},
		Float64Data{1.0, 2.0, 3.0, 4.0, 4.0, 5.0, 6.0, -7.0},
	},
	{
		[]interface{}{"-345", "223", "-654.4", "194", "898.3"},
		Float64Data{-345.0, 223.0, -654.4, 194.0, 898.3},
	},
	{
		[]interface{}{7862, 4234, 9872.1, 8794},
		Float64Data{7862.0, 4234.0, 9872.1, 8794.0},
	},
	{
		[]interface{}{true, false, true, false, false},
		Float64Data{1.0, 0.0, 1.0, 0.0, 0.0},
	},
	{
		[]interface{}{14.3, 26, 17.7, "shoe"},
		Float64Data{14.3, 26.0, 17.7},
	},
	{
		[]uint{34, 12, 65, 230, 30},
		Float64Data{34.0, 12.0, 65.0, 230.0, 30.0},
	},
	{
		[]bool{true, false, true, true, false},
		Float64Data{1.0, 0.0, 1.0, 1.0, 0.0},
	},
	{
		[]float64{10230.9823, 93432.9384, 23443.945, 12374.945},
		Float64Data{10230.9823, 93432.9384, 23443.945, 12374.945},
	},
	{
		[]int{-843, 923, -398, 1000},
		Float64Data{-843.0, 923.0, -398.0, 1000.0},
	},
	{
		[]time.Duration{-843, 923, -398, 1000},
		Float64Data{-843.0, 923.0, -398.0, 1000.0},
	},
	{
		[]string{"-843.2", "923", "hello", "-398", "1000.5"},
		Float64Data{-843.2, 923.0, -398.0, 1000.5},
	},
	{
		map[int]int{0: 456, 1: 758, 2: -9874, 3: -1981},
		Float64Data{456.0, 758.0, -9874.0, -1981.0},
	},
	{
		map[int]float64{0: 68.6, 1: 72.1, 2: -33.3, 3: -99.2},
		Float64Data{68.6, 72.1, -33.3, -99.2},
	},
	{
		map[int]uint{0: 4567, 1: 7580, 2: 98742, 3: 19817},
		Float64Data{4567.0, 7580.0, 98742.0, 19817.0},
	},
	{
		map[int]string{0: "456", 1: "758", 2: "-9874", 3: "-1981", 4: "68.6", 5: "72.1", 6: "-33.3", 7: "-99.2"},
		Float64Data{456.0, 758.0, -9874.0, -1981.0, 68.6, 72.1, -33.3, -99.2},
	},
	{
		map[int]bool{0: true, 1: true, 2: false, 3: true, 4: false},
		Float64Data{1.0, 1.0, 0.0, 1.0, 0.0},
	},
}

func equal(actual, expected Float64Data) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, actualVal := range actual {
		if actualVal != expected[k] {
			return false
		}
	}

	return true
}

func TestLoadRawData(t *testing.T) {
	for _, data := range allTestData {
		actual := LoadRawData(data.actual)
		if !equal(actual, data.expected) {
			t.Fatalf("Transform(%v). Expected [%v], Actual [%v]", data.actual, data.expected, actual)
		}
	}
}
