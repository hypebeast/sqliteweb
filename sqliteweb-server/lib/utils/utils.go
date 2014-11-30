package utils

import (
	"log"
	"os"
	"strings"
)

func FileSize(fileName string) (int64, error) {
	fi, err := os.Stat(fileName)
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}

func ExitOnError(err error, format string, args ...interface{}) {
	if err != nil {
		if strings.HasSuffix(format, "\n") {
			format += "\n"
		}

		log.Printf(format, args)
		os.Exit(1)
	}
}
