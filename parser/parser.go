package parser

import (
	"errors"
)

type Parser struct {
	pos   int64
	input []rune
}

func Load(in string) *Parser {
	return &Parser{pos: 0, input: []rune(in)}
}

func (p *Parser) Next() (rune, error) {
	if int64(len(p.input)) > p.pos {
		return p.input[p.pos], nil
	}
	return '.', errors.New("end of file")
}
