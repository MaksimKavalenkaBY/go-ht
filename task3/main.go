package main

import (
	"fmt"
	"io"
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
)

var operationArguments = map[string][]string{
	"add":      {itemArg, fileNameArg},
	"findById": {idArg, fileNameArg},
	"remove":   {idArg, fileNameArg},
	"list":     {fileNameArg},
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

func Perform(args Arguments, writer io.Writer) error {
	return checkArguments(args)
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
