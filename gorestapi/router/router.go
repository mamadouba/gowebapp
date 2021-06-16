package router

import (
	"gorestapi/appcontext"
)

type Route struct {
	Path        string
	Method      string
	Protected   bool
	Permission  string
	Validator   appcontext.ValidatorFunc
	HandlerFunc appcontext.ContextHandlerFunc
}

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}
