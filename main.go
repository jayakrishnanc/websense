package main

import (
	"fmt"
	"log"
	"os"
	"path"
        "websense/idfile"
        "websense/codeanalyzer"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	file_name := os.Args[1]
	_, err := os.Stat(file_name)

	if err != nil {
		log.Fatal(err)
		return
	}

	fileType := idfile.FindFileType(file_name)

	fmt.Printf("File is [%v] file \n", fileType)

        dir_name := path.Dir(os.Args[0]) + "/codeanalyzer"
        var sad *codeanalyzer.Source_analysis_detail
        
        sad = codeanalyzer.Source_analyze(dir_name,file_name)
        fmt.Println(sad)
}
