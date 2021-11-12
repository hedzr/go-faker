package main

import (
	cmdrrel "github.com/hedzr/go-faker/cli/app/cli/app/cmdr"
)

func init() {
	// build.New(build.NewLoggerConfigWith(true, "logrus", "debug"))
}

func main() {
	cmdrrel.Entry()
	return
}
