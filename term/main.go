package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var execCommandString []string
var history int

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		if args[1] == "~" || args[1] == "~/" {
			home := os.Getenv("HOME")
			fmt.Println("home:", home)
			if home != "" {
				return os.Chdir(home)
			}
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	username := os.Getenv("USER")
	for {
		pwd, _ := os.Getwd()
		pwds := strings.Split(pwd, "/")
		os.Stdin.WriteString(fmt.Sprintf("%s %s ~: ", username, pwds[len(pwds)-1]))
		if scanner.Scan() {
			bytes := scanner.Bytes()
			switch bytes[len(bytes)-1] {
			case 65:
				if history > 0 {
					history--
					execInput(execCommandString[history])
				} else {
					fmt.Fprintln(os.Stderr, "not fond shell exec")
				}
			case 66:
				if history < len(execCommandString) {
					history++
					execInput(execCommandString[history-1])
				} else {
					fmt.Fprintln(os.Stderr, "not fond shell exec")
				}
			default:
				input := scanner.Text()
				switch input {
				case "history":
					history := strings.Join(execCommandString, "\n")
					fmt.Fprintf(os.Stdout, history)
				default:
					execInput(input)
				}
				execCommandString = append(execCommandString, input)
				history = len(execCommandString)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standad input:", err)
		}
	}
}
