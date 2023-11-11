package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Printf("%s$ ", wd)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = Execute(command, &wd)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
