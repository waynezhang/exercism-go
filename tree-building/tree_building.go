package tree

import (
	"errors"
	//"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		var r1 = records[i]
		var r2 = records[j]
		return r1.Parent <= r2.Parent && r1.ID <= r2.ID
	})

	var nodes = make(map[int]*Node)
	var root *Node
	for _, rec := range records {
		var node = &Node{ID: rec.ID}

		// duplicate check
		if nodes[node.ID] != nil {
			return nil, errors.New("")
		}

		// root check
		if rec.ID == 0 {
			if root != nil || rec.Parent != 0 {
				return nil, errors.New("")
			}
			root = node
			nodes[0] = root
			continue
		}

		// validation check
		if rec.ID <= rec.Parent {
			return nil, errors.New("")
		}

		// should has parent
		parent := nodes[rec.Parent]
		if parent == nil {
			return nil, errors.New("")
		}

		// append children
		parent.Children = append(parent.Children, node)
		sort.Slice(parent.Children, func(i, j int) bool {
			return parent.Children[i].ID < parent.Children[j].ID
		})
		nodes[rec.ID] = node
	}

	// check continuous
	for i := range records {
		if nodes[i] == nil {
			return nil, errors.New("")
		}
	}

	return root, nil
}
