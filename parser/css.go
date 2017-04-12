package parser

type StyleSheet struct {
	Rules []Rule
}

func (s *StyleSheet) AddRule(sl Selector, d []Declaration) {
	s.Rules = append(s.Rules, Rule{Select: sl, Declare: d})
}

type Rule struct {
	Select      Selector
	Declare     []Declaration
	Specificity int
}

func ParseCSS(input string) *StyleSheet {
	p := Load(input)
	css := &StyleSheet{Rules: make([]Rule, 0, 0)}
	for {
		p.Whitespace()
		_, err := p.Next()
		if err != nil {
			return css
		}
		sel, err := EatSelectors(p)
		if err != nil {
			return css
		}
		dec, err := EatDeclarations(p)
		if err != nil {
			return css
		}
		css.AddRule(sel, dec)
	}
}
