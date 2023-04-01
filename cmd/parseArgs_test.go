package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args   []string
	config *configGreeter
	err    error
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args: []string{"-h"},
			config: &configGreeter{
				timesPrinted: 0,
				isHelp:       true,
			},
			err: nil,
		},
		{
			args: []string{"5"},
			config: &configGreeter{
				timesPrinted: 5,
				isHelp:       false,
			},
			err: nil,
		},
		{
			args: []string{"abc"},
			config: &configGreeter{
				timesPrinted: 0,
				isHelp:       false,
			},
			err: errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
		},
		{
			args: []string{"1", "foo"},
			config: &configGreeter{
				timesPrinted: 0,
				isHelp:       false,
			},
			err: errors.New(ERR_INVALID_NUM_ARGS),
		},
		{
			args: []string{"0"},
			config: &configGreeter{
				timesPrinted: 0,
				isHelp:       false,
			},
			err: errors.New(ERR_GREATER_THAN_ZERO),
		},
	}

	for _, test := range tests {
		config, err := parseArgs(test.args)
		if test.err != nil && err != test.err {
			t.Fatalf("Expected error to be: %q,\n got: %q\n", test.err, err)
		} else if test.err == nil && err != nil {
			t.Errorf("Expected nil error. Got: %q\n", err)
		} else if test.config.isHelp != config.isHelp {
			t.Errorf("Expected isHelp to be: %t, got: %t\n", test.config.isHelp, config.isHelp)
		} else if test.config.timesPrinted != config.timesPrinted {
			t.Errorf("Expected timesPrinted to be: %d, got: %d\n", test.config.timesPrinted, config.timesPrinted)
		}
	}
}
