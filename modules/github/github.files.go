package github

import (
	"log"
	"os"
)

func CreateFile(filePath string, data string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Println(err)
	}
}
