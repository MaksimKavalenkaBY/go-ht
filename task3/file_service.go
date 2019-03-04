package main

import (
	"encoding/json"
	"fmt"
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

func add(item, fileName string) []byte {
	var user User
	json.Unmarshal([]byte(item), &user)

	if string(findById(user.Id, fileName)) == "" {

	}

	return []byte(fmt.Sprintf("Item with id %s already exists", user.Id))
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
