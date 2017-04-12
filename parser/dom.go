package parser

import (
	"github.com/kenan-rhoton/kraal/dom"
)

func ParseDOM(input string) *dom.Node {
	p := Load(input)
	doc := dom.EmptyElem("document")
	p.MakeChildren(doc)
	return doc
}

func (p *Parser) MakeChildren(d *dom.Node) {
	for {
		p.Whitespace()
		switch n, _ := p.Next(); {
		case n == '<':
			_, _ = p.Eat()
			if i, _ := p.Next(); i == '/' {
				_, _ = p.Eat()
				_, _ = p.EatTagName()
				p.EatTagClosed()
				return
			}
			tag, _ := p.EatTagName()
			if len(tag) > 0 {
				attr, _ := p.EatAttributes()
				e := dom.Elem(string(tag), attr, make([]*dom.Node, 0, 0))
				if p.EatTagClosed() == false {
					p.MakeChildren(e)
				}
				d.Append(e)
			}
		default:
			t, err := p.EatText()
			if len(t) > 0 {
				d.Append(dom.Text(string(t)))
			}
			if err != nil {
				if err.Error() == "end of file" {
					return
				}
			}
		}
	}
}

func (p *Parser) EatAttributes() (map[string]string, error) {
	p.Whitespace()
	res := make(map[string]string)
	for {
		if n, err := p.Next(); n == '>' || n == '/' {
			return res, err
		}
		name, err := p.EatAttrName()
		if len(name) == 0 {
			return res, err
		}
		value, _ := p.EatAttrValue()
		res[string(name)] = string(value)
	}
}

func (p *Parser) EatTagClosed() bool {
	res := false
	for r, err := p.Eat(); r != '>' && err == nil; {
		if r == '/' {
			res = true
		}
		r, err = p.Eat()
	}
	return res
}
