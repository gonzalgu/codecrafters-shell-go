package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from input: %v", err)
		}
		line = line[:len(line)-1]
		parts := strings.Split(line, " ")
		cmd := parts[0]
		//args := parts[1:]

		switch cmd {
		case "exit":
			args := parts[1:]
			n, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("error parsing argument: %v", err)
				os.Exit(-1)
			}
			os.Exit(n)
		case "echo":
			args := parts[1:]
			echoLine := strings.Join(args, " ")
			fmt.Printf("%s\n", echoLine)
		case "cd":
			fmt.Printf("cd received\n")
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
