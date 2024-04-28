package cmddiffunctions_tree

import (
	"slices"
)

type CodeTreeNode struct {
	content     string
	weight      int
	left, right *CodeTreeNode
}

func makeArrayNodes(frequincies map[string]int) ([]CodeTreeNode, error) {
	arr := make([]CodeTreeNode, 0)
	for key, v := range frequincies {
		arr = append(arr, CodeTreeNode{key, v, nil, nil})
	}
	slices.SortFunc(arr, func(i CodeTreeNode, j CodeTreeNode) int {
		if i.weight >= j.weight {
			return 1
		} else {
			return -1
		}
	})
	return arr, nil
}
