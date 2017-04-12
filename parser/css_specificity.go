package parser

import (
	"sort"
)

func (s *StyleSheet) SortRules() {
	s.Calculate()
	sort.Slice(s.Rules, func(i, j int) bool {
		return s.Rules[i].Specificity < s.Rules[j].Specificity
	})
}

func (s *StyleSheet) Calculate() {
	for i, r := range s.Rules {
		s.Rules[i].Specificity = Specify(&r.Select)
	}
}

func Specify(s *Selector) int {
	a, b, c := 0, 0, 0
	if s.Tag != "" && s.Tag != "*" {
		c = 1
	}
	if s.ID != "" {
		a = 1
	}
	b = len(s.Classes)
	children := 0
	if s.Child != nil {
		children = Specify(s.Child)
	} else if s.Descendant != nil {
		children = Specify(s.Descendant)
	}
	return a*100 + b*10 + c + children
}
