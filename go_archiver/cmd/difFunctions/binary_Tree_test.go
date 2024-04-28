package cmddiffunctions_tree

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func TestMakeArrayNodes(t *testing.T) {
	frequincies := make(map[string]int)
	for _, v := range "ABRACADABRA" {
		frequincies[string(v)]++
	}
	want := []CodeTreeNode{{"A", 5, nil, nil}, {"B", 2, nil, nil}, {"R", 2, nil, nil}, {"C", 1, nil, nil}, {"D", 1, nil, nil}}
	slices.SortFunc(want, func(i CodeTreeNode, j CodeTreeNode) int {
		if i.weight >= j.weight {
			return 1
		} else {
			return -1
		}
	})
	res, err := makeArrayNodes(frequincies)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// if (want != res) {
	// 	log.Fatalf("want = %v, res = %v",want,res)
	// }
	for i := range res {
		if res[i] != want[i] {
			log.Fatalf("want[i] = %v, res[i] = %v,i=%v", want[i], res[i], i)
		} else {
			fmt.Println("working")
		}
	}
}
