package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Print(err)
	}

	return content
}
