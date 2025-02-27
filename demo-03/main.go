package main

import "errors"

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}

	return errors.New("file not found")
}

func main() {

}
