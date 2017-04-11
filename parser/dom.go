package parser

import (
	"github.com/kenan-rhoton/kraal/dom"
)

func (p *Parser) ParseDOM() *dom.Node {
	doc := dom.EmptyElem("document")
	p.MakeChildren(doc)
	return doc
}

func (p *Parser) MakeChildren(n *dom.Node) {
	for {
		p.Whitespace()
		switch n := p.Next(); {
		case n == '<':
		default:
			t, err := p.EatText()
			if len(t) > 0 {
				n.Append(Text(t))
			}
			if err != nil {
				if err.Error() == "end of file" {
					break
				}
			}
		}
	}
}
