package logging

import (
	"fmt"
	"log"
	"strings"
)

type stdLogger struct{}

func NewStdLogger() Logger {
	return &stdLogger{}
}

const (
	LvlDebug = "DEBUG"
	LvlInfo  = "INFO"
	LvlWarn  = "WARN"
	LvlError = "ERR"
	LvlFatal = "FATAL"
)

func argsToString(args ...any) string {
	builder := strings.Builder{}

	for i, arg := range args {
		if i%2 == 0 {
			builder.WriteString(arg.(string))
			builder.WriteString("=")
		} else {
			builder.WriteString(fmt.Sprintf("\"%s\"", arg.(string)))
			builder.WriteString(" ")
		}
	}

	return builder.String()
}

func print(level, msg any, args ...any) {
	log.Printf("%s %s %s\n", level, msg, argsToString(args))
}

func (sl *stdLogger) Debug(msg any, args ...any) {
	print(LvlDebug, msg, args...)
}

func (sl *stdLogger) Info(msg any, args ...any) {
	print(LvlInfo, msg, args...)
}

func (sl *stdLogger) Warn(msg any, args ...any) {
	print(LvlWarn, msg, args...)
}

func (sl *stdLogger) Error(msg any, args ...any) {
	print(LvlError, msg, args...)
}

func (sl *stdLogger) Fatal(msg any, args ...any) {
	log.Fatalf("%s %s %s\n", LvlFatal, msg, argsToString(args))
}
