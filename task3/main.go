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

	var content []byte
	switch args[operationArg] {
	case addOps:
		content = add(args[itemArg], args[fileNameArg])
	case findByIdOps:
		content = findById(args[idArg], args[fileNameArg])
	case removeOps:
		remove(args[idArg], args[fileNameArg])
	case listOps:
		content = list(args[fileNameArg])
	}

	fmt.Println(string(content))
	writer.Write(content)

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
