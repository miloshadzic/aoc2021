package input

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/miloshadzic/aoc2021/sub"
)

func GetCommands(input string) []sub.Command {
	var commands []sub.Command

	inputString, err := os.ReadFile(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(inputString), "\n") {
		commandTokens := strings.Split(line, " ")
		if len(commandTokens) != 2 {
			continue
		}

		var dir sub.Direction

		switch commandTokens[0] {
		case "forward":
			dir = sub.Forward
		case "up":
			dir = sub.Up
		case "down":
			dir = sub.Down
		}

		amount, _ := strconv.ParseInt(commandTokens[1], 10, 0)

		commands = append(commands, sub.Command{Direction: dir, Amount: amount})
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
