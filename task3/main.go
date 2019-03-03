package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

type MyError struct {
	msg string
}

const (
	idArg        = "id"
	itemArg      = "item"
	fileNameArg  = "fileName"
	operationArg = "operation"

	addOps      = "add"
	findByIdOps = "findById"
	removeOps   = "remove"
	listOps     = "list"
)

var operationArguments = map[string][]string{
	addOps:      {itemArg, fileNameArg},
	findByIdOps: {idArg, fileNameArg},
	removeOps:   {idArg, fileNameArg},
	listOps:     {fileNameArg},
}

func (error *MyError) Error() string {
	return error.msg
}

func checkArguments(args Arguments) error {
	operation, operationOk := args[operationArg]

	if operationOk && operation != "" {
		arguments, argumentsOk := operationArguments[operation]

		if argumentsOk {
			for _, argument := range arguments {
				value, argumentOk := args[argument]
				if !argumentOk || value == "" {
					return &MyError{fmt.Sprintf("-%s flag has to be specified", argument)}
				}

				return nil
			}
		}

		return &MyError{fmt.Sprintf("Operation %s not allowed!", operation)}
	}

	return &MyError{"-operation flag has to be specified"}
}

func parseArgs() Arguments {
	return nil
}

func add(item, fileName string) {

}

func findById(id, fileName string) string {
	return ""
}

func remove(id, fileName string) {

}

func list(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	return content
}

func Perform(args Arguments, writer io.Writer) error {
	err := checkArguments(args)
	if err != nil {
		return err
	}

	switch args[operationArg] {
	case addOps:
		add(args[itemArg], args[fileNameArg])
	case findByIdOps:
		findById(args[idArg], args[fileNameArg])
	case removeOps:
		remove(args[idArg], args[fileNameArg])
	case listOps:
		writer.Write(list(args[fileNameArg]))
	}

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
