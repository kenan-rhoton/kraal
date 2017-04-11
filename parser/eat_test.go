package parser

import (
	"fmt"
	"testing"
)

func TestEatAString(t *testing.T) {
	test := "La cançó que es cantà"
	runes := []rune(test)
	p := Load(test)
	for i := 0; i <= len(runes); i++ {
		r, err := p.Eat()
		if err != nil {
			if err.Error() != "end of file" {
				fmt.Println("Unexpected error: " + err.Error())
				t.FailNow()
			} else {
				if i != len(runes) {
					fmt.Printf("Wrong string length! Was %d but expected %d\n", i, len(runes))
					t.FailNow()
				}
			}
		} else {
			if r != runes[i] {
				fmt.Printf("Wrong rune received! Got %c(%q) but expected %c(%q)\n", r, r, runes[i], runes[i])
				t.Fail()
			}
		}
	}
}

func TestEatWhile(t *testing.T) {
	test := "Anglo-saxon merlin"
	p := Load(test)
	ru, err := p.EatWhile(func(r rune) bool { return r != '-' })
	res := string(ru)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	if res != "Anglo" {
		t.Fail()
	}
	_, _ = p.Eat()
	ru, err = p.EatWhile(func(r rune) bool { return r != ' ' })
	res = string(ru)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	if res != "saxon" {
		t.Fail()
	}
	_, _ = p.Eat()
	ru, err = p.EatWhile(func(r rune) bool { return true })
	res = string(ru)
	if err != nil {
		if err.Error() != "end of file" {
			fmt.Printf("Error: %s\n", err.Error())
		}
	} else {
		fmt.Println("Wat, no end of file?")
		t.Fail()
	}
	if res != "merlin" {
		t.Fail()
	}
}

func TestWhitespace(t *testing.T) {
	test := "	    \n	\t 	 B"
	p := Load(test)
	p.Whitespace()
	r, err := p.Eat()
	if err != nil {
		t.FailNow()
	}
	if r != 'B' {
		t.FailNow()
	}
}
