package parser

import (
	"fmt"
	"testing"
)

func TestSpecify(t *testing.T) {
	testdata := []struct {
		data  string
		value int
	}{
		{"*{size:2;}", 0},
		{"h1{size:2;}", 1},
		{"h1 b{size:2;}", 2},
		{"h1 b.cat{size:2;}", 12},
		{"h1.boop.lol b.lol{size:2;}", 32},
		{"#dude{size:2;}", 100},
		{"#dude h1{size:2;}", 101},
		{"h1{size:2;}", 1},
	}
	for _, v := range testdata {
		css := ParseCSS(v.data)
		if res := Specify(&css.Rules[0].Select); res != v.value {
			t.Errorf("Wrong Specificity for \"%s\": expected %d, got %d", v.data, v.value, res)
		}
	}
}

func ExampleSort() {
	testdata := `
second.yeah #mustbesecond{s:0;}
#first #veryfirst{s:0;}
last{s:0;}
.third #justbarely{t:1;}
`
	css := ParseCSS(testdata)
	css.SortRules()
	for _, r := range css.Rules {
		s := r.Select
		if s.Tag != "" {
			fmt.Printf("%s - %d\n", s.Tag, r.Specificity)
		} else if len(s.Classes) > 0 {
			fmt.Printf("%s - %d\n", s.Classes[0], r.Specificity)
		} else {
			fmt.Printf("%s - %d\n", s.ID, r.Specificity)
		}
	}
	// Output:
	// last - 1
	// third - 110
	// second - 111
	// first - 200
}
