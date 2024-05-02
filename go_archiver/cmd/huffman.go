package cmd

import (
	cmddiffunctions_tree "archiver/cmd/difFunctions"
	"encoding/json"
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
	packed, dict := compressTxt(data) // TODO:WRITE ENCODE FUNCTION
	//packedFileName = func()
	if err := os.WriteFile(packedFileName(filePath), []byte(packed), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	json_dict, _ := json.Marshal(dict)
	os.WriteFile("codes.txt", json_dict, 0644)

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
func compressTxt(data []byte) ([]byte, map[string]string) {
	minil_codes := makeMinimalCodes(string(data))              //TODO:DELETE THIS SHIT
	compressed_data := compressFile(minil_codes, string(data)) //TODO:WRITE
	return compressed_data, minil_codes                        //TODO:
}
func makeMinimalCodes(data string) map[string]string {
	map_counts := make(map[string]int)
	for _, v := range data {
		map_counts[string(v)]++
	}
	binary_tree, err := cmddiffunctions_tree.MakeCodeTree(map_counts)
	if err != nil {
		handleError(err)
	}
	result := make(map[string]string, len(map_counts))
	for key := range map_counts {
		_, result[key] = cmddiffunctions_tree.GethufCode(binary_tree, key, "")
		if err != nil {
			handleError(err)
		}
	}
	return result
}
func compressFile(codes map[string]string, data string) []byte {
	bin_data := dataToBinary(codes, data)
	return binaryDataToByteSlices(bin_data)
}
func dataToBinary(codes map[string]string, data string) string {
	string_data_bin := ""
	for _, v := range data {
		string_data_bin += codes[string(v)]
	}
	return string_data_bin
}
func binaryDataToByteSlices(binaryData string) []byte {
	index := 0
	ByteSlice := make([]byte, 0)
	var new_byte byte = 0
	for _, v := range binaryData {
		if v == '1' {
			new_byte |= 1
			new_byte <<= 1

		} else {
			new_byte <<= 1
		}
		index++
		if index == 8 {
			index = 0
			ByteSlice = append(ByteSlice, new_byte)
			new_byte = 0
		}
	}
	return ByteSlice
}

//func (tree *TreeForCodes) append() //TODO:DODELAT
