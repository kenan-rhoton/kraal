package frame

import (
	"github.com/kenan-rhoton/kraal/dom"
	"github.com/kenan-rhoton/kraal/parser"
	"strings"
)

func BuildObject(n *dom.Node) *FrameObject {
	o := &FrameObject{}
	o.Tag = n.Type()
	o.Attributes = n.Attrs()
	o.Styles = make(map[string]string)
	if o.Attributes["style"] != "" {
		p := parser.Load(o.Attributes["style"])
		dec, err := parser.EatDeclarations(p)
		if err == nil {
			for _, d := range dec {
				o.ApplyStyleLocal(d)
			}
		}
		if o.Attributes["class"] != "" {
			o.Classes = strings.Split
		}
	}
	o.Children = make([]*FrameObject, 0)
	for _, c := range n.Children() {
		o.Children = append(o.Children, BuildObject(c))
	}
	return o
}

func (o *FrameObject) ApplyStyleLocal(d parser.Declaration) {
	o.Styles[d.Name()] = d.Value()
}

func (o *FrameObject) ApplyStyleRecursive(d parser.Declaration) {
	o.ApplyStyleLocal(d)
        for _, c := range o.Children {
            o.ApplyStyleRecursive(d)
        }
}

func (o *FrameObject) MatchSelector(sel *parser.Selector) bool {
	if sel.ID != "" && sel.ID != o.ID {
		return false
	}
	if sel.Tag != "" && sel.Tag != o.Tag {
		return false
	}
	if sel.Classes != nil {
		if o.Classes == nil {
			return false
		}
		for _, c := range sel.Classes {
			found = false
			for _, find := range o.Classes {
				if c == find {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func (o *FrameObject) ApplyStyleNow(r *parser.Rule) {
	if o.MatchSelector(r.Select) {
		if r.Select.Child != nil {
			for _, c := range o.Children {
				c.ApplyStyleNow(r.Select.Child)
			}
		} else if r.Select.Descendant != nil {
			for _, c := range o.Children {
				c.ApplyStyle(r.Select.Descendant)
			}
		} else {
			for _, dec := range r.Declare {
				if isInheritable(dec) {
					o.ApplyStyleRecursive(dec)
				} else {
					o.ApplyStyleLocal(dec)
				}
			}
		}
		return true
	} else {
		return false
	}
}

func (o *FrameObject) ApplyStyle(r *parser.Rule) {
	if o.ApplyStyleNow(r) == false {
		for _, c := range o.Children {
			c.ApplyStyle(r)
		}
	}
}

func (o *FrameObject) ApplyStyles(css *parser.StyleSheet) {
	for _, r := range css.Rules {
		o.ApplyStyle(r)
	}
}

type FrameObject struct {
	Tag        string
	Attributes map[string]string
	Classes    []string
	ID         string
	Styles     map[string]string
	Children   []*FrameObject
}
