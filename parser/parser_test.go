package parser

import (
	"testing"
)

func TestLoadBackAndForth(t *testing.T) {
	teststring := "fadsfoe ^á ------\t \n\n lol"
	p := Load(teststring)
	if string(p.input) != teststring {
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	testdata := []string{
		"fasldfnasldnfads",
		"tjwetiiwoemfiem",
		"       aoieiar",
		"\t jub",
		"€pàle",
	}
	for _, v := range testdata {
		p := Load(v)
		if n, _ := p.Next(); n != []rune(v)[0] {
			t.Fail()
		}
	}
}

func TestNextEOF(t *testing.T) {
	p := Load("")
	_, err := p.Next()
	if err == nil {
		t.Fail()
	} else {
		if err.Error() != "end of file" {
			t.Fail()
		}
	}
}
