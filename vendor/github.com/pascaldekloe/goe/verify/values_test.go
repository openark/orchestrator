package verify

import (
	"math/big"
	"reflect"
	"strings"
	"testing"
)

type O struct {
	B bool
	I int
	U uint
	F float64
	C complex64
	A [2]int
	X interface{}
	M map[int]interface{}
	P *O
	S []byte
	T string
	O P
}

type P struct {
	Name string
}

var goldenIdentities = []interface{}{
	0:  nil,
	1:  false,
	2:  true,
	3:  "",
	4:  "a",
	5:  'b',
	6:  byte(1),
	7:  2,
	8:  4.8,
	9:  float32(-0.0),
	10: 16i,
	11: O{},
	12: &O{},
	13: O{
		B: true,
		I: -9,
		U: 9,
		F: .8,
		C: 7i,
		A: [2]int{6, 5},
		X: "inner",
		M: map[int]interface{}{4: 'q', 3: nil},
		P: &O{
			T: "inner",
			P: &O{},
		},
		S: []byte{2, 1},
		T: "text",
		O: P{Name: "test"},
	},
	14: big.NewInt(32), // unexported fields & encoding.TextMarshaller
}

func TestGoldenIdentities(t *testing.T) {
	for i, x := range goldenIdentities {
		if !Values(t, "same", x, x) {
			t.Errorf("%d: Rejected: %#v", i, x)
		}
	}
}

var goldenDiffers = []struct {
	a, b interface{}
	msg  string
}{
	0: {true, false, "Got true, want false"},
	1: {"a", "b", `Got "a", want "b"`},
	2: {O{U: 10}, O{U: 11}, "/U: Got 10 (0xa), want 11 (0xb)"},

	3: {int(2), uint(2), "Got type int, want uint"},
	4: {O{X: O{}}, O{X: &O{}},
		"/X: Got type verify.O, want *verify.O"},

	5: {&O{}, nil, "Unwanted *verify.O"},
	6: {nil, &O{}, "Missing *verify.O"},
	7: {O{}, O{X: O{}}, "/X: Missing verify.O"},
	8: {O{X: O{}}, O{}, "/X: Unwanted verify.O"},

	9: {[]int{1, -2}, []int{1, -3},
		"[1]: Got -2, want -3"},
	10: {map[int]bool{0: false, 1: false, 2: true}, map[int]bool{0: false, 1: true, 2: true},
		"[1]: Got false, want true"},
	11: {map[rune]bool{'f': false}, map[rune]bool{'f': false, 't': true},
		"[116]: Missing bool"},
	12: {map[string]int{"false": 0, "true": 1, "?": 2}, map[string]int{"false": 0, "true": 1},
		`["?"]: Unwanted int`},

	13: {O{X: func() int { return 9 }}, O{X: func() int { return 0 }},
		"/X: Can't compare functions"},
	14: {
		O{
			P: &O{S: []byte{3, 10}},
		},
		O{
			P: &O{S: []byte{5, 11}},
		},
		"/P/S[0]: Got 3, want 5\n/P/S[1]: Got 10 (0xa), want 11 (0xb)"},

	15: {O{T: "abcdefghijklmnoprstuvwxyz"}, O{T: "abcdefghijklmnopqrstuvwxyz"},
		"/T: Got \"abcdefghijklmnoprstuvwxyz\", want \"abcdefghijklmnopqrstuvwxyz\"\n                         ^"},
	16: {O{O: P{Name: "Jo"}}, O{O: P{Name: "Joe"}},
		`/O/Name: Got "Jo", want "Joe"`},

	17: {O{X: big.NewInt(1000)}, O{X: big.NewInt(1001)},
		"/X: Type big.Int with unexported fields [\"neg\" \"abs\"] not equal"},
}

func TestGoldenDiffers(t *testing.T) {
	for i, gold := range goldenDiffers {
		tr := &travel{}
		tr.values(reflect.ValueOf(gold.a), reflect.ValueOf(gold.b), nil)

		msg := tr.report("case")
		if len(msg) == 0 {
			t.Errorf("%d: No report", i)
			continue
		}

		if i := strings.IndexRune(msg, '\n'); i < 0 {
			t.Fatalf("%d: Report header absent in %q", i, msg)
		} else {
			msg = msg[i+1:]
		}

		if msg != gold.msg {
			t.Errorf("%d:\nGot:  %q\nWant: %q", i, msg, gold.msg)
		}
	}
}

func TestNilEquivalents(t *testing.T) {
	var sn, se []byte
	var mn, me map[int]string
	se = make([]byte, 0, 1)
	me = make(map[int]string, 1)

	type Containers struct {
		S []byte
		M map[int]string
	}

	Values(t, "nil vs empty slice", sn, se)
	Values(t, "nil vs empty map", mn, me)
	Values(t, "empty vs nil slice", se, sn)
	Values(t, "empty vs nil map", me, mn)
	Values(t, "nil vs empty embedded", Containers{}, Containers{S: se, M: me})
	Values(t, "empty vs nil embedded", Containers{S: se, M: me}, Containers{})
}
