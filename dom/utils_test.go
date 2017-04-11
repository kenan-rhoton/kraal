package dom

import (
	"fmt"
	"testing"
)

func TestIndent(t *testing.T) {
	in := "This is a\n    Very annoying ex\n ample but it shoudn'\n            t matter"
	out := "  This is a\n      Very annoying ex\n   ample but it shoudn'\n              t matter"
	if indent(in) != out {
		fmt.Printf("Expected:\n\n%s\nGot:\n\n%s", out, indent(in))
		t.FailNow()
	}
}
