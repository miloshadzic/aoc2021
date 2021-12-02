package input

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type Command struct {
	Direction Direction
	Amount    int64
}

func GetCommands(input string) []Command {
	var commands []Command

	inputString, err := os.ReadFile(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(inputString), "\n") {
		commandTokens := strings.Split(line, " ")
		if len(commandTokens) != 2 {
			continue
		}

		var dir Direction

		switch commandTokens[0] {
		case "forward":
			dir = Forward
		case "up":
			dir = Up
		case "down":
			dir = Down
		}

		amount, _ := strconv.ParseInt(commandTokens[1], 10, 0)

		commands = append(commands, Command{dir, amount})
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
