package cmddiffunctions_tree

import (
	"log"
	"maps"
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
func TestGethufCode(t *testing.T) {
	want := map[string]string{"A": "0", "B": "11", "R": "101", "C": "1001", "D": "1000"}
	frequincies := make(map[string]int)
	for _, v := range "ABRACADABRA" {
		frequincies[string(v)]++
	}
	codeTree, err := MakeCodeTree(frequincies)
	if err != nil {
		log.Fatalf("%v", err)
	}
	have := make(map[string]string, len(want))
	for key := range frequincies {
		_, have[key] = GethufCode(codeTree, key, "")
	}
	if !maps.Equal(want, have) {
		log.Fatalf("want = %v, have = %v,", want, have)
	}
}
