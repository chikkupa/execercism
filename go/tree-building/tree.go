package tree

// Record Struct for containing id and parent
type Record struct {
	ID     int
	Parent int
}

// Node Node of a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build Builds a tree structure from array of records
func Build(records []Record) (*Node, error) {
	var tree *Node
	for _, record := range records {
		if record.ID == 0 && tree == nil {
			tree = &Node{ID: record.ID}
		} else if record.ID == 0 && tree != nil {
			tree.Children = append(tree.Children, &Node{ID: record.ID})
		} else if tree != nil {
			addBranch(tree, record)
		} else {
			return tree, nil
		}
	}

	return tree, nil
}

func addBranch(tree *Node, record Record) {
	if tree.ID == record.Parent {
		tree.Children = append(tree.Children, &Node{ID: record.ID})
	} else {
		for _, children := range tree.Children {
			addBranch(children, record)
		}
	}
}
