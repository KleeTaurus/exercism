package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func newNode(id int) *Node {
	return &Node{
		ID:       id,
		Children: make([]*Node, 0),
	}
}

func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	mapping := make(map[int]*Node)
	for i := 0; i < len(records); i++ {
		if i != records[i].ID {
			return nil, errors.New("None continuous")
		}

		if i == 0 {
			if records[i].Parent != 0 {
				return nil, errors.New("The root Parent must be 0")
			}
		} else {
			if records[i].ID == records[i].Parent {
				return nil, errors.New("Cycle directly")
			}
		}

		node := newNode(records[i].ID)
		mapping[records[i].ID] = node

		if i != 0 {
			parent, ok := mapping[records[i].Parent]
			if !ok {
				return nil, errors.New("Can not find parent")
			}
			parent.Children = append(parent.Children, node)
		}

	}

	return mapping[0], nil
}
