package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.stdin unbuffered
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "q" || t == "Q" {
			os.Exit(0)
		}
		println(strings.ToUpper(t))

	}
	err := scanner.Err()
	if err != nil {
		// EOF is not an error
		fmt.Fprint(os.Stderr, "error", err)
		os.Exit(1)
	}
}
