package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error: %+v",err)
		return
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Printf("Error: %+v",err)
		return
	}

	var versionString string

	for idx := range bytes {
		if bytes[idx] == 'g' {
			versionIdx := idx
			versionEnd := idx
			if bytes[versionIdx+1] == 'o' {
				versionEnd := versionEnd + 2
				for (bytes[versionEnd] >= 48 && bytes[versionEnd] <= 57) || bytes[versionEnd] == 46 {
					versionEnd++
				}
				if versionEnd > versionIdx+2 {
					versionString = string(bytes[versionIdx:versionEnd])
					break
				}
			}
		}
	}

	fmt.Printf("Go Version: %s\n",versionString)
}