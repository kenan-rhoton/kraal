package parser

import (
	"errors"
	"unicode"
)

func (p *Parser) Eat() (rune, error) {
	if int64(len(p.input)) > p.pos {
		res := p.input[p.pos]
		p.pos++
		return res, nil
	}
	return '.', errors.New("end of file")
}

func (p *Parser) EatWhile(f func(rune) bool) ([]rune, error) {
	res := make([]rune, 0, 1)
	for {
		r, err := p.Eat()
		if err != nil {
			return res, err
		}
		if f(r) == false {
			p.pos--
			return res, nil
		}
		res = append(res, r)
	}
}

func (p *Parser) Whitespace() error {
	_, err := p.EatWhile(unicode.IsSpace)
	return err
}

func (p *Parser) EatText() ([]rune, error) {
	return p.EatWhile(isText)
}

func (p *Parser) EatTagName() ([]rune, error) {
	return p.EatWhile(isTagName)
}

func (p *Parser) EatAttrName() ([]rune, error) {
	return p.EatWhile(isAttrName)
}

func (p *Parser) EatAttrValue() ([]rune, error) {
	return p.EatWhile(isAttrValue)
}
