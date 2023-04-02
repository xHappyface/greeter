package main

import (
	"errors"
	"testing"
)

type testConfigValidateArgs struct {
	config *configGreeter
	err    error
}

func TestValidateArgs(t *testing.T) {
	tests := []testConfigValidateArgs{
		{
			config: &configGreeter{},
			err:    errors.New(ERR_GREATER_THAN_ZERO),
		},
		{
			config: &configGreeter{
				timesPrinted: -5,
			},
			err: errors.New(ERR_GREATER_THAN_ZERO),
		},
		{
			config: &configGreeter{
				timesPrinted: 10,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		err := validateArgs(test.config)
		if test.err != nil && errors.Is(err, test.err) {
			t.Errorf("Expected error to be: %q,\ngot: %q\n", test.err, err)
		} else if test.err == nil && err != nil {
			t.Errorf("Expected error to be nil, got: %q\n", err)
		}
	}
}
