package main

import (
	"fmt"
)

func main() {
	files, err := FileWalker("./")
	if err != nil {
		fmt.Println("Error walking the directory:", err)
		return
	}

	for _, file := range files {
		result, err := ParserWrapper(file)
		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", file, err)
			continue
		}

		if result == nil {
			continue
		}
		
	  fmt.Printf("result for file %s: %+v\n", file, result)	
	}
}	
