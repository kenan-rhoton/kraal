package dom

import (
	"fmt"
	"testing"
)

var testtags = []string{"div", "html", "p", "fake", "strong"}

func TestElemCreatedWithItsType(t *testing.T) {
	for i, data := range testtags {
		el := EmptyElem(data)
		if el.Type() != testtags[i] {
			fmt.Printf("%s element did not have type \"%s\"\n", data, testtags[i])
			t.FailNow()
		}
	}
}

func TestElemChildrenAreCorrect(t *testing.T) {
	parent := EmptyElem("parent")
	for i, data := range testtags {
		child := EmptyElem(data)
		parent.Append(child)
		if parent.Child(i).Type() != testtags[i] {
			fmt.Printf("Expected parent's child to be of type '%s', received '%s'", testtags[i], parent.Child(0).Type())
			t.FailNow()
		}
	}
	if len(parent.Children()) != len(testtags) {
		fmt.Printf("Expected final length to be %d, got %d", len(testtags), len(parent.Children()))
		t.FailNow()
	}
}
