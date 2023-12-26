package challenge

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Input struct {
	lines chan string
}

func (input *Input) Lines() <-chan string {
	return input.lines
}

func (input *Input) AllLines() []string {
	var lines []string
	for line := range input.lines {
		lines = append(lines, line)
	}
	return lines
}

func GetInput(cmd *cobra.Command) *Input {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", cmd.Parent().Use))
	if err != nil {
		panic(err)
	}
	input := &Input{lines: make(chan string)}
	go readFile(file, input)
	return input
}

func readFile(file *os.File, input *Input) {
	defer file.Close()
	defer close(input.lines)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input.lines <- scanner.Text()
	}
}
