package appcontext

import (
	"context"
	"gorestapi/config"
	"net/http"

	"github.com/jinzhu/gorm"
)

type AppContext struct {
	env *config.Config
	db  *gorm.DB
}

func New(env *config.Config, db *gorm.DB) *AppContext {
	if db == nil {
		panic("Database not set")
	}
	return &AppContext{env: env, db: db}
}

func (ctx *AppContext) GetDB() *gorm.DB {
	return ctx.db
}

func (ctx *AppContext) GetEnv() *config.Config {
	return ctx.env
}

func (ctx *AppContext) Get(r *http.Request) map[string]interface{} {
	m, ok := r.Context().Value("data").(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
	}
	return m
}

func (ctx *AppContext) Set(r *http.Request, key string, value interface{}) {
	m := ctx.Get(r)
	m[key] = value
	*r = *r.WithContext(context.WithValue(r.Context(), "data", m))
}

type ContextHandlerFunc func(*AppContext, http.ResponseWriter, *http.Request)

type ValidatorFunc func() interface{}
