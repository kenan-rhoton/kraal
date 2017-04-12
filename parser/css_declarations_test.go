package parser

import (
	"testing"
)

func TestCSS_DeclarationNumber(t *testing.T) {
	testdata := []struct {
		a string
		b int
	}{
		{"h1{size:3;}", 1},
		{"h1{size:3;money:5;}", 2},
		{"*{size:3;heigth:10;breakfast: \"eggs and ham\";boopiness:42;}", 4},
	}
	for _, v := range testdata {
		css := ParseCSS(v.a)
		if len(css.Rules[0].Declare) != v.b {
			t.Fail()
		}
	}
}
