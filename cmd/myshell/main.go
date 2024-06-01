package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from input: %v", err)
		}
		cmd = cmd[:len(cmd)-1]
		switch cmd {
		case "echo":
			fmt.Printf("echo received\n")
		case "cd":
			fmt.Printf("cd received\n")
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
