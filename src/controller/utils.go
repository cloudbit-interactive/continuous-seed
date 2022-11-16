package controller

import (
	"bytes"
	"github.com/cloudbit-interactive/cuppago"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func ReplaceVariables(values map[string]interface{}) map[string]interface{} {
	for key := range values {
		values[key] = ReplaceToSystemValue(cuppago.String(values[key]))
	}
	return values
}

func ReplaceString(string string) string {
	for key := range YamlVars {
		string = cuppago.ReplaceNotCase(string, "\\${"+key+"}", cuppago.String(YamlVars[key]))
	}
	string = ReplaceToSystemValue(string)
	return string
}

func ReplaceToSystemValue(string string) string {
	cuppago.Log("ReplaceToSystemValue", string)
	os := runtime.GOOS
	date := time.Now().String()
	string = cuppago.ReplaceNotCase(string, "\\${DATE}", date[0:10])
	string = cuppago.ReplaceNotCase(string, "\\${DATETIME}", date[0:19])
	string = cuppago.ReplaceNotCase(string, "\\${OS}", os)
	return string
}

func BashCommand(command string) string {
	outputString := ""
	command = ReplaceString(command)
	Log("-- CMD: " + command)
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		outputString = cuppago.String(err)
	} else {
		outputString = string(output)
	}
	outputString = strings.TrimSpace(outputString)
	outputString = strings.Trim(outputString, "\n")
	Log("---- output: " + outputString)
	return outputString
}

func Command(app string, args []string, workingDirectory string) string {
	for i := 0; i < len(args); i++ {
		args[i] = strings.TrimSpace(ReplaceString(args[i]))
	}
	workingDirectory = strings.TrimSpace(ReplaceString(workingDirectory))
	Log("-- CMD: "+app, "-- args: ", args, "-- workingDirectory: "+workingDirectory)
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
	Log("---- output: " + outputString)
	return outputString
}

func Log(values ...interface{}) {
	if YamlData["log"] != true {
		return
	}
	cuppago.LogFile(values...)
}
