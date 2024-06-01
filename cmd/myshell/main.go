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
		case "type":
			arg := parts[1]
			if isBuiltin(arg) {
				fmt.Printf("%s is a shell builtin\n", arg)
			} else if path, ok := isExecutable(arg, os.Getenv("PATH")); ok {
				fmt.Printf("%s is %s/%s\n", arg, path, arg)
			} else {
				fmt.Printf("%s: not found\n", arg)
			}
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

func isBuiltin(cmd string) bool {
	switch cmd {
	case "echo":
		return true
	case "exit":
		return true
	case "type":
		return true
	default:
		return false
	}
}

func isExecutable(cmd string, path string) (string, bool) {
	//fmt.Printf("PATH: %s\n", path)
	dirs := strings.Split(path, ":")
	for _, d := range dirs {
		//fmt.Printf("dir: %s\n", d)
		items, _ := os.ReadDir(d)
		for _, item := range items {
			//fmt.Printf("item: %v\n", item)
			if item.Name() == cmd {
				return d, true
			}
		}
	}
	return "", false
}
