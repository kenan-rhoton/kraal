package dom

import (
	"fmt"
	"strings"
)

type Node struct {
	children  []*Node
	node_type NodeType
}

type NodeType interface {
	Type() string
	Attrs() map[string]string
	Text() string
}

func (n *Node) Type() string {
	return n.node_type.Type()
}

func (n *Node) Append(c *Node) {
	n.children = append(n.children, c)
}

func (n *Node) Child(i int) *Node {
	if len(n.children) <= i {
		return nil
	}
	return n.children[i]
}

func (n *Node) Children() []*Node {
	return n.children
}

func (n *Node) Attrs() map[string]string {
	return n.node_type.Attrs()
}

func (n *Node) AttrString() string {
	attrs := n.node_type.Attrs()
	if attrs == nil {
		return ""
	} else {
		results := make([]string, 0, len(attrs))
		for key, value := range attrs {
			results = append(results, key+"="+value)
		}
		return strings.Join(results, " ")
	}
}

func (n *Node) PrettyString() string {
	if n.node_type.Type() == "text" {
		return n.node_type.Text() + "\n"
	}
	if len(n.children) == 0 {
		return fmt.Sprintf("<%s%s/>\n", n.node_type.Type(), n.AttrString())
	} else {
		res := fmt.Sprintf("<%s%s>\n", n.node_type.Type(), n.AttrString())
		for _, child := range n.children {
			res = res + indent(child.PrettyString())
		}
		res = res + fmt.Sprintf("</%s>\n", n.node_type.Type())
		return res
	}
}

func (n *Node) PrettyPrint() {
	fmt.Println(n.PrettyString())
}
