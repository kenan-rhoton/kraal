package frame

import (
	"github.com/kenan-rhoton/kraal/dom"
	"github.com/kenan-rhoton/kraal/parser"
)

func BuildObject(n *dom.Node) *FrameObject {
	o := &FrameObject{}
	o.Class = n.Type()
	o.Attributes = n.Attrs()
	o.Styles = make(map[string]string)
	if o.Attributes["style"] != "" {
		css := parser.ParseCSS(o.Attributes["style"])
		for _, r := range css.Rules {
			for _, d := range r.Declare {
				o.ApplyStyleLocal(d)
			}
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

func (o *FrameObject) ApplyStyles(css *parser.StyleSheet) {

}

type FrameObject struct {
	Class      string
	Attributes map[string]string
	Styles     map[string]string
	Children   []*FrameObject
}
