package tree

import (
	"errors"
	"sort"
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
	sort.Slice(records, func(i int, j int) bool { return records[i].ID < records[j].ID })
	var tree *Node

	if len(records) > 0 && records[0].ID != 0 {
		return tree, errors.New("no root node")
	}

	mRecords := make(map[int]*Node)

	for _, record := range records {
		mRecords[record.ID] = &Node{ID: record.ID}
	}

	for index, record := range records {
		if record.ID == 0 && record.Parent != 0 {
			return tree, errors.New("root node has parent")
		}
		if index+1 < len(records) && records[index].ID == records[index+1].ID {
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
			tree = mRecords[record.ID]
		} else if tree != nil {
			mRecords[record.Parent].Children = append(mRecords[record.Parent].Children, mRecords[record.ID])
		} else {
			return tree, nil
		}
	}

	return tree, nil
}
