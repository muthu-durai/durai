package main

import (
	"errors"
	"fmt"
	"os"
)

func checkForFile(fPath string) error {
	if fPath == "" {
		return errors.New("The Given Path Is Empty")
	}

	f, err := os.Open(fPath) // Trying to Open a file

	if err != nil {
		return fmt.Errorf("InValid Path %v", f)
	}

	// if f.Name() == "unexpected" {
	// 	return fmt.Errorf("The Path Of the File Is Wrong")
	// }

	return nil
}

func main() {
	err := checkForFile("../Clone")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The Entered File Is Present")
	}
}