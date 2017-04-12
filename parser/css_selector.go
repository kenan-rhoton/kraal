package parser

import (
	"errors"
)

type Selector struct {
	Tag        string
	ID         string
	Classes    []string
	Child      *Selector
	Descendant *Selector
}

func EatSelector(p *Parser) string {
	res, _ := p.EatWhile(func(r rune) bool { return is(r, []rune(alpha+num+"*-_")) })
	return string(res)
}

func EatSelectors(p *Parser) (Selector, error) {
	err := errors.New("invalid")
	sel := Selector{Classes: make([]string, 0, 0)}
	for {
		switch n, err2 := p.Next(); {
		case err2 != nil:
			return sel, err
		case n == '.':
			_, _ = p.Eat()
			ret := EatSelector(p)
			if ret == "" {
				return sel, errors.New("invalid")
			}
			sel.Classes = append(sel.Classes, ret)
			err = nil
		case n == '#':
			_, _ = p.Eat()
			ret := EatSelector(p)
			if ret == "" {
				return sel, errors.New("invalid")
			}
			sel.ID = ret
			err = nil
		case n == ' ' || n == '>':
			p.Whitespace()
			switch l, _ := p.Next(); {
			case l == '>':
				_, _ = p.Eat()
				p.Whitespace()
				child, selerr := EatSelectors(p)
				if selerr == nil {
					sel.Child = &child
				} else {
					err = selerr
				}
				return sel, err
			case l == '{':
				return sel, err
			default:
				child, selerr := EatSelectors(p)
				if selerr == nil {
					sel.Descendant = &child
				} else {
					err = selerr
				}
				return sel, err
			}
		case n == '{':
			return sel, err
		default:
			ret := EatSelector(p)
			if ret == "" {
				return sel, errors.New("invalid")
			}
			sel.Tag = ret
			err = nil
		}
	}
	return sel, err
}
