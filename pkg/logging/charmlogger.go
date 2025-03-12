package logging

import (
	"io"

	"github.com/charmbracelet/log"
)

func NewCharmLogger(w io.Writer, debug bool) Logger {
	level := log.InfoLevel

	if debug {
		level = log.DebugLevel
	}

	return log.NewWithOptions(w, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
		Level:           level,
	})
}
