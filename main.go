package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	executable, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}
	executable = strings.TrimSuffix(filepath.Base(executable), filepath.Ext(executable))

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		os.Exit(1)
	}
	executable = reg.ReplaceAllString(executable, "")
	executable = strings.ToUpper(executable)

	var exitcode int

	exitFromEnv := os.Getenv(executable + "_EXIT")

	if exitFromEnv == "" {
		exitcode = 0
	} else {
		exitcode, err = strconv.Atoi(exitFromEnv)
		if err != nil {
			exitcode = 1
		}

	}

	os.Exit(exitcode)

}
