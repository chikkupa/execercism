package tree

import(
	"errors"
)

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
	records = sortRecords(records)
	var tree *Node

	if(len(records) > 0 && records[0].ID != 0){
		return tree, errors.New("no root node")
	}

	for index, record := range records {
		if record.ID == 0 && record.Parent != 0 {
			return tree, errors.New("root node has parent")
		}
		if index + 1 < len(records) && records[index].ID == records[index + 1].ID {
			return tree, errors.New("duplicate node")
		}
		if record.ID != index {
			return tree, errors.New("non-continuous")
		}
		if record.ID != 0 && record.ID == record.Parent {
			return tree, errors.New("cycle directly")
		}
		if record.ID < record.Parent {
			return tree, errors.New("higher id parent of lower id")
		}
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

func sortRecords(records []Record) []Record {
	for i := 0; i < len(records) - 1; i++ {
		for j :=0; j < len(records) - i - 1; j++ {
			if(records[j].ID > records[j+1].ID){
				records[j], records[j+1] = records[j+1], records[j]
			}
		}
	}
	return records
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
