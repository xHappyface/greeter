package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

type testConfigRunCmd struct {
	config *configGreeter
	input  string
	output string
	err    error
}

func TestRunCmd(t *testing.T) {
	tests := []testConfigRunCmd{
		{
			config: &configGreeter{
				timesPrinted: 5,
			},
			input:  "",
			output: STR_ASK_FOR_NAME,
			err:    ERR_EMPTY_NAME_STRING,
		},
		{
			config: &configGreeter{
				timesPrinted: 5,
			},
			input:  "Bill Bryson",
			output: STR_ASK_FOR_NAME + strings.Repeat("Nice to meet you, Bill Bryson!\n", 5),
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, test := range tests {
		r := strings.NewReader(test.input)
		err := runCmd(r, byteBuf, test.config)
		if err != nil && test.err == nil {
			t.Fatalf("Expected error to be nil, got: %q\n", err)
		} else if test.err != nil && !errors.Is(err, test.err) {
			t.Fatalf("Expected error to be: %q,\n, got: %q\n", test.err, err)
		}
		gotMsg := byteBuf.String()
		if gotMsg != test.output {
			t.Errorf("Expected stdout message to be: %q\n, got: %q\n", test.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
