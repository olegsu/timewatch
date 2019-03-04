package cmd

import (
	"os"

	"github.com/olegsu/timewatch/pkg/logger"
)

func buildLogger(command string) logger.Logger {
	return logger.New(&logger.Options{
		Verbose: rootCmdOptions.Verbose,
		Command: command,
	})
}

func dieOnError(err error, log logger.Logger) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
