package main

import "fmt"

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

func checkArguments(args Arguments) error {
	operation, operationOk := args[operationArg]

	if !operationOk || operation == "" {
		return &MyError{"-operation flag has to be specified"}
	}

	arguments, argumentsOk := operationArguments[operation]

	if !argumentsOk {
		return &MyError{fmt.Sprintf("Operation %s not allowed!", operation)}
	}

	for _, argument := range arguments {
		value, argumentOk := args[argument]
		if !argumentOk || value == "" {
			return &MyError{fmt.Sprintf("-%s flag has to be specified", argument)}
		}
	}

	return nil
}

func parseArgs() Arguments {
	return nil
}
