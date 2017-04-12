package parser

import (
	"testing"
)

func TestCSS_Rules(t *testing.T) {
	testdata := []struct {
		a string
		b int
	}{
		{"h1{size:3;}", 1},
		{"h1{size:3;}h1{size:3;}", 2},
		{"h1{size:3;}h1{size:3;}h1{size:3;}h1{size:3;}", 4},
	}
	for _, v := range testdata {
		css := ParseCSS(v.a)
		if len(css.Rules) != v.b {
			t.Fail()
		}
	}
}
