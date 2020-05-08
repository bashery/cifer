package fmt24

import (
	"./fmt24"
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	for _, c := range cases {
		if c.output != fmt24.Format(c.input) {
			fmt.Errorf("Format of %s is not %s", fmt24.Format(c.input), c.output)
			break
		}

		fmt.Printf("Pass : %s  ==> %s\n", c.input, c.output)
	}
}

var cases = []struct {
	input  string
	output string
}{
	{"01:30:23PM", "13:30:23"},
	{"01:30:23AM", "01:30:23"},
	{"05:30:23PM", "17:30:23"},
}
