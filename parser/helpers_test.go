package parser

import (
	"testing"
)

func TestIs(t *testing.T) {
	testdata := []struct {
		char   rune
		list   string
		result bool
	}{
		{'a', "abcdefg", true},
		{'€', "!\"·$€%&/()", true},
		{'G', "gFH", false},
		{' ', "\t\n", false},
		{'â', "_â-¨d", true},
		{'ë', "_â-¨e", false},
	}

	for _, v := range testdata {
		if is(v.char, []rune(v.list)) != v.result {
			t.Fail()
		}
	}
}
