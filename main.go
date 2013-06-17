package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	b, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data map[string]interface{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(m))
}
