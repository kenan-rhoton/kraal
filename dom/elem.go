package dom

type ElementNode struct {
	tag        string
	attributes map[string]string
}

func (e *ElementNode) Attrs() map[string]string {
	return e.attributes
}

func (e *ElementNode) Type() string {
	return e.tag
}

func (e *ElementNode) Text() string {
	return ""
}

func Elem(name string, attrs map[string]string, children []*Node) *Node {
	return &Node{children: children, node_type: &ElementNode{tag: name, attributes: attrs}}
}

func EmptyElem(name string) *Node {
	return Elem(name, make(map[string]string), make([]*Node, 0, 0))
}
