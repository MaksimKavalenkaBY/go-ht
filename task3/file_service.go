package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Arguments map[string]string

type User struct {
	Id    string
	Email string
	Age   int
}

func (user User) String() string {
	return fmt.Sprintf("{\"id\":\"%s\",\"email\":\"%s\",\"age\":%d}", user.Id, user.Email, user.Age)
}

func add(item, fileName string) ([]byte, error) {
	if !exists(fileName) {
		createFile(fileName)
	}

	data, err := list(fileName)
	if err != nil {
		return nil, err
	}

	var user User
	json.Unmarshal([]byte(item), &user)

	data, err = findByIdData(user.Id, data)
	if err != nil {
		return nil, err
	}

	if string(data) != "" {
		return []byte(fmt.Sprintf("Item with id %s already exists", user.Id)), nil
	}

	var users []User
	json.Unmarshal(data, &users)
	users = append(users, user)

	var s []string
	for _, user := range users {
		s = append(s, user.String())
	}

	data = []byte(fmt.Sprintf("[%s]", strings.Join(s, ",")))
	return nil, writeToFile(fileName, data)
}

func findById(id, fileName string) ([]byte, error) {
	data, err := readFile(fileName)

	if err != nil {
		return nil, err
	}

	return findByIdData(id, data)
}

func findByIdData(id string, data []byte) ([]byte, error) {
	var users []User
	json.Unmarshal(data, &users)

	for _, user := range users {
		if user.Id == id {
			return []byte(user.String()), nil
		}
	}
	return []byte(""), nil
}

func remove(id, fileName string) {

}

func list(fileName string) ([]byte, error) {
	return readFile(fileName)
}
