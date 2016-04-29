package boontling

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type Logger struct{}

func (l Logger) Debug(r *http.Request, format string, args ...interface{}) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, format, args...)
}
