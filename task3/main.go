package main

import (
	"fmt"
	"io"
	"os"
)

func Perform(args Arguments, writer io.Writer) error {
	err := checkArguments(args)
	if err != nil {
		return err
	}

	var data []byte
	var fileErr error
	switch args[operationArg] {
	case addOps:
		data, fileErr = add(args[itemArg], args[fileNameArg])
	case findByIdOps:
		data, fileErr = findById(args[idArg], args[fileNameArg])
	case removeOps:
		data, fileErr = remove(args[idArg], args[fileNameArg])
	case listOps:
		data, fileErr = list(args[fileNameArg])
	}

	fmt.Println(string(data))
	writer.Write(data)

	return fileErr
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
