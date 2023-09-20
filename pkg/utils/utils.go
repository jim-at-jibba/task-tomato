package utils

import (
	"fmt"
	"log"
	"os"
)

func SetUpPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("%s/task-tomato", dirname)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Print(err.Error())
		}
	}

	return path
}
