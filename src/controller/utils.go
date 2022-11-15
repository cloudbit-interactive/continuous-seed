package controller

import (
	"bytes"
	"github.com/cloudbit-interactive/cuppago"
	"os/exec"
	"strings"
)

func ReplaceString(string string) string {
	for key := range YamlVars {
		string = cuppago.ReplaceNotCase(string, "{"+key+"}", cuppago.String(YamlVars[key]))
	}
	return string
}

func Command(app string, args []string, workingDirectory string) string {
	for i := 0; i < len(args); i++ {
		args[i] = strings.TrimSpace(args[i])
	}
	var output bytes.Buffer
	cmd := exec.Command(app, args...)
	cmd.Dir = workingDirectory
	cmd.Stdout = &output
	err := cmd.Run()
	outputString := ""
	if err != nil {
		outputString = cuppago.String(err)
	} else {
		outputString = output.String()
	}
	outputString = strings.TrimSpace(outputString)
	outputString = strings.Trim(outputString, "\n")
	return outputString
}
