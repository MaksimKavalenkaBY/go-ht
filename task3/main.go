package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

type User struct {
	Id    string
	Email string
	Age   int
}

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

func (user User) String() string {
	return fmt.Sprintf("{\"id\":\"%s\",\"email\":\"%s\",\"age\":%d}", user.Id, user.Email, user.Age)
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

func readFile(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Print(err)
	}

	return content
}

func add(item, fileName string) {
	content := readFile(fileName)
	var users interface{}
	json.Unmarshal(content, &users)

}

func findById(id, fileName string) []byte {
	content := readFile(fileName)
	var users []User
	json.Unmarshal(content, &users)

	for _, user := range users {
		if user.Id == id {
			return []byte(user.String())
		}
	}
	return []byte("")
}

func remove(id, fileName string) {

}

func list(fileName string) []byte {
	content := readFile(fileName)
	return content
}

func Perform(args Arguments, writer io.Writer) error {
	err := checkArguments(args)
	if err != nil {
		return err
	}

	var content []byte
	switch args[operationArg] {
	case addOps:
		add(args[itemArg], args[fileNameArg])
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
