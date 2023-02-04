package goconcepts

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type item struct {
	id    string
	name  string
	price int
	img   string
}

//  var _ io.Reader = (*os.File)(nil)

func ReadJson() {
	file, err := os.Open("data/menu.json")
	if err != nil {
		fmt.Println("error while opening the file", err)
	}
	// here file implements the io.Reader interface so you can pass it to the io.ReadAll
	// https://stackoverflow.com/questions/25677235/create-a-io-reader-from-a-local-file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error while readint he file", err)
	}
	var payload []item
	err = json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println("error while opening the file")
	}


}
