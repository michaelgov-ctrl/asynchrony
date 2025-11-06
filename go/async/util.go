package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func appendFileA() {
	appendFile("fileA.txt")
}

func appendFileB() {
	appendFile("fileB.txt")
}

func appendFile(path string) {
	perms := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	f, err := os.OpenFile(path, perms, 0644)
	if err != nil {
		log.Fatal("danger will robinson")
	}
	defer f.Close()

	now := time.Now().Format("2006-01-02 15:04:05.000")
	if _, err := f.WriteString(fmt.Sprintf("%s\n", now)); err != nil {
		log.Fatal("danger will robinson")
	}

	fmt.Printf("appended: %s, to: %s\n", now, path)
}
