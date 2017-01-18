package stats

import (
	"testing"
)

func TestError(t *testing.T) {
	err := statsErr{"test error"}
	if err.Error() != "test error" {
		t.Errorf("Error method message didn't match")
	}
}
