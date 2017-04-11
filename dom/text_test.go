package dom

import (
	"fmt"
	"testing"
)

func TestTextCreatedWithTextType(t *testing.T) {
	txt := Text("Potato")
	if txt.Type() != "text" {
		fmt.Println("Text element did not have type \"text\"")
		t.FailNow()
	}
}
