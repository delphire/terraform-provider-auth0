package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/logging"
)

const LogPrefix = "AUTH0"

func Logger() (*log.Logger, error) {

	if output, err := logging.LogOutput(); err != nil {
		return nil, err
	} else {
		flags := log.Ldate | log.Ltime
		return log.New(output, LogPrefix, flags), nil
	}
}

func LogPrintf(s string, v ...interface{}) {
	if logger, err := Logger(); err != nil {
		panic(err)
	} else {
		logger.Printf(s, v...)
	}
}
