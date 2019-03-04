package main

import (
	"io/ioutil"
	"os"
)

func exists(fileName string) bool {
	_, err := os.Stat(fileName)

	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

func createFile(fileName string) error {
	newFile, err := os.Create(fileName)

	if err != nil {
		return err
	}

	newFile.Close()
	return nil
}

func readFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeToFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, 0666)

	if err != nil {
		return err
	}

	return nil
}
