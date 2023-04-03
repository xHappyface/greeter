package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var usage = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name your entered <integer> number of times.
`, os.Args[0])

var (
	ERR_POS_ARG_SPECIFIED = errors.New("positional arguments specified")
	ERR_GREATER_THAN_ZERO = errors.New("must specify a number greater than 0")
	ERR_EMPTY_NAME_STRING = errors.New("empty name string")
)

const (
	STR_ASK_FOR_NAME = "Your name, please? Press ENTER when done.\n"
)

type configGreeter struct {
	timesPrinted int
}

func main() {
	r, w := os.Stdin, os.Stdout
	c, err := parseArgs(w, os.Args[1:])
	if err != nil {
		if !(errors.Is(err, ERR_POS_ARG_SPECIFIED)) {
			os.Exit(1)
		}
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

func parseArgs(w io.Writer, args []string) (*configGreeter, error) {
	config := new(configGreeter)
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&config.timesPrinted, "n", 0, "Number of times to greet.")
	err := fs.Parse(args)
	if err != nil {
		return config, err
	} else if fs.NArg() != 0 {
		return config, ERR_POS_ARG_SPECIFIED
	}
	return config, nil
}

func validateArgs(c *configGreeter) error {
	if !(c.timesPrinted > 0) {
		return ERR_GREATER_THAN_ZERO
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c *configGreeter) error {
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
		return "", ERR_EMPTY_NAME_STRING
	}
	return name, nil
}

func greetUser(w io.Writer, c *configGreeter, name string) {
	msg := fmt.Sprintf("Nice to meet you, %s!\n", name)
	for i := 0; i < c.timesPrinted; i++ {
		fmt.Fprint(w, msg)
	}
}
