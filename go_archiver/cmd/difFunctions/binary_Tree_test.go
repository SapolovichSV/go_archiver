package cmddiffunctions_tree

import (
	"log"
	"testing"
)

func TestMakeCodeTree(t *testing.T) {
	frequincies := make(map[string]int)
	for _, v := range "ABRACADABRA" {
		frequincies[string(v)]++
	}
	codeTree, err := MakeCodeTree(frequincies)
	if err != nil {
		log.Fatalf("%v,#%v,%v", codeTree, codeTree, err)
	}
}
