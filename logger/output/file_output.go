package output

import (
	"fmt"
	"os"
	"strings"
)

type FileOutput struct {
	files        map[string]*os.File
	pathTemplate string
	replacer     string
}

func NewFileOutput(pathTemplate string, replacer string) FileOutput {
	return FileOutput{
		files:        make(map[string]*os.File),
		pathTemplate: pathTemplate,
		replacer:     replacer,
	}
}

func (f FileOutput) Log(format string, parameters map[string]string) {
	path := f.pathTemplate

	for k, v := range parameters {
		path = strings.ReplaceAll(path, fmt.Sprintf(f.replacer, k), v)
	}

	file, ok := f.files[path]

	if !ok {
		file, err := os.Create(path)

		if err != nil {
			return
		}

		f.files[path] = file

		f.log(file, format, parameters)
	} else {
		f.log(file, format, parameters)
	}
}

func (f FileOutput) log(file *os.File, format string, parameters map[string]string) {
	for k, v := range parameters {
		format = strings.Replace(format, fmt.Sprintf(f.replacer, k), v, -1)
	}

	_, err := fmt.Fprintln(file, format)

	if err != nil {
		fmt.Println(err)
	}
}
