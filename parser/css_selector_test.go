package parser

import (
	"testing"
)

func TestCSS_TagSelector(t *testing.T) {
	css := ParseCSS("h1{size:3;}")
	if css.Rules[0].Select.Tag != "h1" {
		t.Fail()
	}
}

func TestCSS_UnnamedSelector(t *testing.T) {
	css := ParseCSS("*{size:3;}")
	if css.Rules[0].Select.Tag != "*" {
		t.Fail()
	}
}

func TestCSS_ClassSelector(t *testing.T) {
	css := ParseCSS(".boop{size:3;}")
	if css.Rules[0].Select.Classes[0] != "boop" {
		t.Fail()
	}
}

func TestCSS_IDSelector(t *testing.T) {
	css := ParseCSS("#boop{size:3;}")
	if css.Rules[0].Select.ID != "boop" {
		t.Fail()
	}
}

func TestCSS_ChildSelector(t *testing.T) {
	css := ParseCSS("h1>b{size:3;}")
	if css.Rules[0].Select.Child == nil {
		t.FailNow()
	}
	if css.Rules[0].Select.Child.Tag != "b" {
		t.Fail()
	}
}

func TestCSS_DescendantSelector(t *testing.T) {
	css := ParseCSS("h1 b{size:3;}")
	if css.Rules[0].Select.Descendant == nil {
		t.FailNow()
	}
	if css.Rules[0].Select.Descendant.Tag != "b" {
		t.Fail()
	}
}

func TestCSS_CombinedSelectors(t *testing.T) {
	css := ParseCSS("h1.boop > #all{size:3;}")
	s := css.Rules[0].Select
	if s.Tag != "h1" || s.Classes[0] != "boop" || s.Child.ID != "all" {
		t.Fail()
	}
}
