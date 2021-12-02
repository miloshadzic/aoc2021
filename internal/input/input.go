package input

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCommands(input string) [][]string {
	var commands [][]string

	inputString, err := os.ReadFile(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(inputString), "\n") {
		command := strings.Split(line, " ")
		if len(command) == 2 {
			commands = append(commands, command)
		}
	}

	return commands
}

func GetLines(input string) []string {
	var lines []string

	inputString, err := os.ReadFile(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(inputString), "\n") {
		if err == nil {
			lines = append(lines, line)
		}
	}

	return lines
}

func GetInt64s(input string) []int64 {
	var ints []int64

	inputString, err := os.ReadFile(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(inputString), "\n") {
		i, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			ints = append(ints, i)
		}
	}

	return ints
}
