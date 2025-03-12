package logging

import (
	"io"

	"github.com/charmbracelet/log"
)

func NewCharmLogger(w io.Writer) Logger {
	return log.NewWithOptions(w, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}
