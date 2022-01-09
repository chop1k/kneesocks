package output

import (
	"fmt"
	"strings"
)

type ConsoleOutput struct {
	replacer string
}

func NewConsoleOutput(replacer string) ConsoleOutput {
	return ConsoleOutput{
		replacer: replacer,
	}
}

func (c ConsoleOutput) Log(format string, parameters map[string]string) {
	for k, v := range parameters {
		format = strings.ReplaceAll(format, fmt.Sprintf(c.replacer, k), v)
	}

	fmt.Println(format)
}
