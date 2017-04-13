package parser

import (
	"errors"
)

type Declaration struct {
	name  string
	value string
}

func (d *Declaration) Name() string {
	return d.name
}

func (d *Declaration) Value() string {
	return d.value
}

func EatDeclName(p *Parser) string {
	res, _ := p.EatWhile(func(r rune) bool { return r != ':' })
	return string(res)
}

func EatDeclValue(p *Parser) string {
	res, _ := p.EatWhile(func(r rune) bool { return r != ';' })
	return string(res)
}

func EatDeclarations(p *Parser) ([]Declaration, error) {
	decs := make([]Declaration, 0, 0)
	p.Whitespace()
	if n, _ := p.Next(); n == '{' {
		_, _ = p.Eat()
		p.Whitespace()
	}
	for {
		if n, _ := p.Next(); n == '}' {
			_, _ = p.Eat()
			return decs, nil
		}
		name := EatDeclName(p)
		if name == "" {
			return decs, errors.New("invalid")
		}
		p.Whitespace()
		if n, _ := p.Next(); n == ':' {
			_, _ = p.Eat()
			p.Whitespace()
		} else {
			return decs, errors.New("invalid")
		}
		value := EatDeclValue(p)
		p.Whitespace()
		if n, _ := p.Next(); n == ';' {
			_, _ = p.Eat()
			p.Whitespace()
		} else {
			return decs, errors.New("invalid")
		}
		decs = append(decs, Declaration{name, value})
	}
}
