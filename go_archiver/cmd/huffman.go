package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var huffmanCmd = &cobra.Command{
	Use:   "huffman",
	Short: "pack by using huffman method",
	Run:   pack,
}

const packExt = "huf"

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		handleError(errors.New("EMPTY PATH TO FILE"))
		os.Exit(1)
	}
	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		handleError(err) //TODO:REFACT
	}
	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err) //TODO:REFACTOR
	}
	//data -> Encode(data)
	packed := compressTxt(data) // TODO:WRITE ENCODE FUNCTION
	//packedFileName = func()
	if err := os.WriteFile(packedFileName(filePath), []byte(packed), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
}
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	return (strings.TrimSuffix(fileName, ext) + string('.') + packExt)
}

func init() {
	packCmd.AddCommand(huffmanCmd)
}
func compressTxt(data []byte) string {
	//minil_codes := makeMinimalCodes(string(data)) //TODO:DELETE THIS SHIT
	//compressed_data = compressFile()      //TODO:WRITE
	return string(data) //TODO:
}
func makeMinimalCodes(data string) map[string]string {
	map_counts := make(map[string]int)
	for _, v := range data {
		map_counts[string(v)]++
	}

	result := make(map[string]string)
	return result
}

type TreeForCodes struct {
	left, right *TreeForCodes
	val         node
} //TODO:DODELAT
type node struct {
	key  string
	code string
}

//func (tree *TreeForCodes) append() //TODO:DODELAT
