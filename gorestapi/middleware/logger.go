package middleware

import (
	"gorestapi/appcontext"
	"gorestapi/logger"
	"gorestapi/modules/user"
	"gorestapi/utils"
	"net/http"
	"time"
)

func Logger(next appcontext.ContextHandlerFunc) appcontext.ContextHandlerFunc {
	return func(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := utils.NewResponseWriter(w)
		next(ctx, lrw, r)
		usr, _ := ctx.Get(r)["user"].(user.User)
		host := r.Host
		method := r.Method
		url := r.URL.String()
		end := time.Since(start)
		format := "user=%s host=%s method=%s url=%s status=%d time=%s bytes=%d\n"
		logger.Info(format, usr.Email, host, method, url, lrw.Status(), end, lrw.Size())
	}
}
