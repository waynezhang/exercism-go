package pov

import "fmt"

// Node doc here
type Node struct {
	data     string
	parent   *Node
	children []*Node
}

// Graph doc here
type Graph struct {
	root  *Node
	nodes map[string]*Node
}

// New doc here
func New() *Graph {
	return &Graph{root: &Node{children: []*Node{}}, nodes: make(map[string]*Node)}
}

// AddNode doc here
func (g *Graph) AddNode(nodeLabel string) {
	node := Node{data: nodeLabel, children: []*Node{}}
	g.nodes[nodeLabel] = &node
}

// AddArc doc here
func (g *Graph) AddArc(from, to string) {
	if g.nodes[from] == nil {
		g.AddNode(from)
	}
	if g.nodes[to] == nil {
		g.AddNode(to)
	}
	g.nodes[from].children = append(g.nodes[from].children, g.nodes[to])
	g.nodes[to].parent = g.nodes[from]

	for node := g.nodes[from]; ; node = node.parent {
		if node.parent == nil {
			g.root = node
			break
		}
	}
}

// ArcList doc here
func (g *Graph) ArcList() []string {
	arcs := make([]string, 0)

	nodes := []*Node{g.root}
	for {
		if len(nodes) == 0 {
			break
		}
		newNodes := []*Node{}
		for _, n := range nodes {
			for _, child := range n.children {
				arcs = append(arcs, fmt.Sprintf("%s -> %s", n.data, child.data))
				newNodes = append(newNodes, child)
			}
		}
		nodes = newNodes
	}
	return arcs
}

// ChangeRoot doc here
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	g.root = g.nodes[newRoot]
	node := g.root
	for {
		parent := node.parent
		if parent == nil {
			break
		}
		for i, n := range parent.children {
			if n == node {
				parent.children = append(parent.children[:i], parent.children[i+1:]...)
				break
			}
		}
		node.children = append(node.children, parent)
		node = parent
	}
	return g
}
