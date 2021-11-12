package main

import (
	cmdrrel "cmdr-starter/cli/app/cli/app/cmdr"
)

func init() {
	// build.New(build.NewLoggerConfigWith(true, "logrus", "debug"))
}

func main() {
	cmdrrel.Entry()
	return
}
