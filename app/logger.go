package app

import (
	"flag"
	"github.com/whosonfirst/go-spatial/flags"
	"github.com/whosonfirst/go-whosonfirst-log"
	"io"
	"os"
)

func NewApplicationLogger(fl *flag.FlagSet) (*log.WOFLogger, error) {

	verbose, _ := flags.BoolVar(fl, "verbose")

	logger := log.SimpleWOFLogger()
	level := "status"

	if verbose {
		level = "debug"
	}

	stdout := io.Writer(os.Stdout)
	logger.AddLogger(stdout, level)

	return logger, nil
}
