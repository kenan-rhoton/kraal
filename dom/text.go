package dom

type TextNode struct {
	text string
}

func (t *TextNode) Attrs() map[string]string {
	return nil
}

func (t *TextNode) Type() string {
	return "text"
}

func (t *TextNode) Text() string {
	return t.text
}

func Text(data string) *Node {
	return &Node{children: make([]*Node, 0, 1), node_type: &TextNode{text: data}}
}
