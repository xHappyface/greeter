package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

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
	name         string
}

func main() {
	r, w := os.Stdin, os.Stdout
	c, err := parseArgs(w, os.Args[1:])
	if err != nil {
		if !(errors.Is(err, ERR_POS_ARG_SPECIFIED)) {
			os.Exit(1)
		}
		fmt.Fprintf(w, "ERR: %v\n", err)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintf(w, "ERR: %v\n", err)
		os.Exit(1)
	}
	err = runCmd(r, w, c)
	if err != nil {
		fmt.Fprintf(w, "ERR; %v\n", err)
		os.Exit(1)
	}
}

func parseArgs(w io.Writer, args []string) (*configGreeter, error) {
	config := new(configGreeter)
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		strUsage := "A greeter application which prints the name your entered a specified number of times.\n\nUsage of %s: <options> [name]\n"
		fmt.Fprintf(w, strUsage, fs.Name())
		fmt.Fprint(w, "Options:\n\n")
		fs.PrintDefaults()
		fmt.Fprintln(w)
	}
	fs.IntVar(&config.timesPrinted, "n", 0, "Number of times to greet.")
	err := fs.Parse(args)
	if err != nil {
		return config, err
	} else if fs.NArg() > 1 {
		return config, ERR_POS_ARG_SPECIFIED
	}
	config.name = fs.Arg(0)
	return config, nil
}

func validateArgs(c *configGreeter) error {
	if !(c.timesPrinted > 0) {
		return ERR_GREATER_THAN_ZERO
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c *configGreeter) error {
	var err error
	if !(len(c.name) > 0) {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(w, c)
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

func greetUser(w io.Writer, c *configGreeter) {
	msg := fmt.Sprintf("Nice to meet you, %s!\n", c.name)
	for i := 0; i < c.timesPrinted; i++ {
		fmt.Fprint(w, msg)
	}
}
