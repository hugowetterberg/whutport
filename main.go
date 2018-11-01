package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: whutport appname [appname2 appname3]")
		os.Exit(1)
	}

	for _, name := range os.Args[1:] {
		s := []byte(name)
		port := make([]byte, 4)
		for i := 0; i < len(s); i++ {
			port[i%4] += s[i]
		}

		if port[0]%10 == 0 {
			port[0] = 1
		}

		fmt.Printf("%s: ", name)
		for i := 0; i < len(port); i++ {
			fmt.Printf("%d", port[i]%10)
		}
		fmt.Println()
	}
}
