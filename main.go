package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No such file!")
		return
	}

	total := 0

	for i := 0; i < len(args); i++ {
		file, error := ioutil.ReadFile(args[i])

		if error != nil {
			return
		}

		current := bytes.Count(file, []byte{'\n'})
		total += current

		fmt.Println(current, args[i])
	}

	if (len(args) > 1) {
		fmt.Println(total, "total")
	}
}
