package main

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"bufio"
)

func execute(command string) {
	arr := strings.Split(command, " ")
	var args []string
	if len(arr) > 0 {
		for index, value := range arr {
			if index > 0 {
				args = append(args, value)
			}
		}
	}
	fmt.Println(">>", command)
	out, err := exec.Command(arr[0], args...).Output()
	if err != nil {
		fmt.Println("Cannot execute command:", command)
	} else {
		if len(string(out)) > 0 {
			fmt.Println(string(out))
		}
	}
}

func main() {
    path, err := os.Getwd()
	if err != nil {
	    fmt.Println(err)
	    return
	}

	file, err := os.Open(path + "\\Makefile")
	if err != nil {
		fmt.Println("Makefile not found")
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	for _, line := range lines {
		execute(line)
	}
}