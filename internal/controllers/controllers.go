package controllers

import (
	"net/http"

	"reflect"
	"runtime"
	"strings"

	"github.com/CaptainFallaway/realgofile/pkg/logging"
	"github.com/go-chi/chi/v5"
)

type Controller interface {
	SetupRoutes(router chi.Router)
}

type LoggingDecorator struct {
	logger logging.Logger
}

func NewLoggingDecorator(logger logging.Logger) *LoggingDecorator {
	return &LoggingDecorator{logger}
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func getFunctionName(unc any) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(unc).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func (ld *LoggingDecorator) Decorate(unc Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We could use a pool here so we don't reallocate for each error
		err := unc(w, r)
		if err != nil {
			handlerName := getFunctionName(unc)
			ld.logger.Error("handler returned error", "handler", handlerName, "err", err)
		}
	}
}
