package inst

import (
	"testing"

	test "github.com/openark/golib/tests"
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
	}

	for _, p := range patterns {
		if match := RegexpMatchPatterns(p.s, p.patterns); match != p.expected {
			t.Errorf("RegexpMatchPatterns failed with: %q, %+v, got: %+v, expected: %+v", p.s, p.patterns, match, p.expected)
		}
	}
}

func TestRedactGtidSetUUID(t *testing.T) {
	gtidSet := `00020192-1111-1111-1111-111111111111:1-838,00020194-3333-3333-3333-333333333333:1-5`
	{
		redacted := redactGtidSetUUID(gtidSet, "00020192-1111-1111-1111-111111111111")
		test.S(t).ExpectEquals(redacted, ",00020194-3333-3333-3333-333333333333:1-5")
	}
	{
		redacted := redactGtidSetUUID(gtidSet, "00020194-3333-3333-3333-333333333333")
		test.S(t).ExpectEquals(redacted, "00020192-1111-1111-1111-111111111111:1-838,")
	}
	{
		redacted := redactGtidSetUUID(gtidSet, "00020193-2222-2222-2222-222222222222")
		test.S(t).ExpectEquals(redacted, gtidSet)
	}
	{
		gtidSet := `00020192-1111-1111-1111-111111111111:1-838,00020194-3333-3333-3333-333333333333:1-5,00020192-1111-1111-1111-111111111111:839-840`
		redacted := redactGtidSetUUID(gtidSet, "00020192-1111-1111-1111-111111111111")
		test.S(t).ExpectEquals(redacted, ",00020194-3333-3333-3333-333333333333:1-5,")
	}
	{
		gtidSet := ""
		redacted := redactGtidSetUUID(gtidSet, "00020193-2222-2222-2222-222222222222")
		test.S(t).ExpectEquals(redacted, "")
	}
}
