package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var usage = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name your entered <integer> number of times.
`, os.Args[0])

const (
	ERR_INVALID_NUM_ARGS  = "invalid number of arguments"
	ERR_GREATER_THAN_ZERO = "must specify a number greater than 0"
	ERR_EMPTY_NAME_STRING = "empty name string"

	STR_ASK_FOR_NAME = "Your name, please? Press ENTER when done.\n"
)

type configGreeter struct {
	timesPrinted int
	isHelp       bool
}

func main() {
	r, w := os.Stdin, os.Stdout
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(w, "ERR: %v\n", err)
		fmt.Fprint(w, usage)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintf(w, "ERR: %v\n", err)
		fmt.Fprint(w, usage)
		os.Exit(1)
	}
	err = runCmd(r, w, c)
	if err != nil {
		fmt.Fprintf(w, "ERR; %v\n", err)
		fmt.Fprint(w, usage)
		os.Exit(1)
	}
}

func parseArgs(args []string) (*configGreeter, error) {
	if len(args) != 1 {
		return &configGreeter{}, errors.New(ERR_INVALID_NUM_ARGS)
	}
	if args[0] == "-h" || args[0] == "--help" {
		return &configGreeter{isHelp: true}, nil
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return &configGreeter{}, err
	}
	return &configGreeter{timesPrinted: num}, nil
}

func validateArgs(c *configGreeter) error {
	if !(c.timesPrinted > 0) {
		return errors.New(ERR_GREATER_THAN_ZERO)
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c *configGreeter) error {
	if c.isHelp {
		fmt.Fprint(w, usage)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(w, c, name)
	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	scanner := bufio.NewScanner(r)
	fmt.Fprintf(w, STR_ASK_FOR_NAME)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if !(len(name) > 0) {
		return "", errors.New(ERR_EMPTY_NAME_STRING)
	}
	return name, nil
}

func greetUser(w io.Writer, c *configGreeter, name string) {
	msg := fmt.Sprintf("Nice to meet you, %s!\n", name)
	for i := 0; i < c.timesPrinted; i++ {
		fmt.Fprint(w, msg)
	}
}
