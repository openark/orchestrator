package inst

import (
	"testing"
)

type testPatterns struct {
	s        string
	patterns []string
	expected bool
}

func TestRegexpMatchPatterns(t *testing.T) {
	patterns := []testPatterns{
		{"hostname", []string{}, false},
		{"hostname", []string{"blah"}, false},
		{"hostname", []string{"blah", "blah"}, false},
		{"hostname", []string{"host", "blah"}, true},
		{"hostname", []string{"blah", "host"}, true},
		{"hostname", []string{"ho.tname"}, true},
		{"hostname", []string{"ho.tname2"}, false},
		{"hostname", []string{"ho.*me"}, true},
		{"10.0.0.3", []string{"10.0.0.3"}, true},
		{"10.0.0.3", []string{"10.0.0.3:3306"}, true},
		{"10.0.0.3", []string{"10.0.0.38"}, false},
	}

	for _, p := range patterns {
		t.Run(p.s, func(t *testing.T) {
			k := &InstanceKey{Hostname: p.s, Port: 3306}
			if match := FiltersMatchInstanceKey(k, p.patterns); match != p.expected {
				t.Errorf("FiltersMatchInstanceKey failed with: %q, %+v, got: %+v, expected: %+v", p.s, p.patterns, match, p.expected)
			}
		})
	}
}
