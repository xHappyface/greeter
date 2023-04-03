package main

import (
	"bytes"
	"errors"
	"flag"
	"testing"
)

type testConfigParseArgs struct {
	args   []string
	config *configGreeter
	err    error
}

func TestParseArgs(t *testing.T) {
	tests := []testConfigParseArgs{
		{
			args:   []string{"-h"},
			config: &configGreeter{},
			err:    flag.ErrHelp,
		},
		{
			args: []string{"-n", "5"},
			config: &configGreeter{
				timesPrinted: 5,
			},
			err: nil,
		},
		{
			args:   []string{"-n", "abc"},
			config: &configGreeter{},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
		},
		{
			args: []string{"-n", "1", "foo"},
			config: &configGreeter{
				timesPrinted: 1,
			},
			err: ERR_POS_ARG_SPECIFIED,
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, test := range tests {
		config, err := parseArgs(byteBuf, test.args)
		if test.err != nil && errors.Unwrap(err) != errors.Unwrap(test.err) {
			t.Errorf("Expected error to be: %q,\ngot: %q\n", test.err, err)
		} else if test.err == nil && err != nil {
			t.Errorf("Expected error to be nil, got: %q\n", err)
		} else if config.timesPrinted != test.config.timesPrinted {
			t.Errorf("Expected timesPrinted to be: %d, got %d\n", test.config.timesPrinted, config.timesPrinted)
		}
	}
	byteBuf.Reset()
}
