package main

import (
	"fmt"
	"gorestapi/appcontext"
	"gorestapi/config"
	"gorestapi/db"
	"gorestapi/logger"
	"gorestapi/middleware"
	"gorestapi/modules/auth"
	"gorestapi/modules/category"
	"gorestapi/modules/product"
	"gorestapi/modules/user"
	"gorestapi/router"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, route router.RoutePrefix, ctx *appcontext.AppContext) {
	s := r.PathPrefix(route.Prefix).Subrouter()
	for _, sr := range route.SubRoutes {
		var handler appcontext.ContextHandlerFunc
		handler = sr.HandlerFunc
		if sr.Protected {
			handler = middleware.Auth(handler, sr.Permission)

		}
		if sr.Validator != nil {
			handler = middleware.Validate(handler, sr.Validator)
		}
		s.Methods(sr.Method).
			Path(sr.Path).
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				middleware.Logger(handler)(ctx, w, r)
			})
	}
}

func main() {
	logger.Init("app", os.Getenv("DEBUG") == "1")
	logger.Info("Start Application")
	cnfg := config.Configuration
	db.Connect(cnfg)
	db.DB.AutoMigrate(&user.User{})
	db.DB.AutoMigrate(&auth.Token{})
	db.DB.AutoMigrate(&auth.PwdResetToken{})
	db.DB.AutoMigrate(&category.Category{})
	db.DB.AutoMigrate(&product.Product{})
	ctx := appcontext.New(cnfg, db.DB)

	address := fmt.Sprintf(":%s", cnfg.Port)
	r := mux.NewRouter().StrictSlash(false)
	RegisterRoutes(r, user.Routes, ctx)
	RegisterRoutes(r, auth.Routes, ctx)
	RegisterRoutes(r, category.Routes, ctx)
	RegisterRoutes(r, product.Routes, ctx)
	go func() {
		if err := http.ListenAndServe(address, r); err != nil {
			logger.Error(err.Error())
		}
	}()
	logger.Info("Running Server :%s", cnfg.Port)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	logger.Info("Close database")
	db.DB.Close()
	logger.Info("Shutdown server")
}
