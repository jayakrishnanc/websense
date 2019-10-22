package main

import (
	"fmt"
	"log"
	"os"
        "idfile"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	fileName := os.Args[1]
	_, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	fileType := idfile.FindFileType(fileName)

	fmt.Printf("File is [%v] file \n", fileType)
}
