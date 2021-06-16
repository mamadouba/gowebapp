package auth

import (
	"gorestapi/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/auth",
	SubRoutes: []router.Route{
		router.Route{
			Path:        "/register",
			Method:      "POST",
			Protected:   false,
			HandlerFunc: register,
			Validator:   registerValidator,
		},
		router.Route{
			Path:        "/login",
			Method:      "POST",
			Protected:   false,
			HandlerFunc: login,
			Validator:   loginValidator,
		},
		router.Route{
			Path:        "/refresh",
			Method:      "GET",
			Protected:   false,
			HandlerFunc: refresh,
		},
		router.Route{
			Path:        "/resetpwdreq/{email}",
			Method:      "GET",
			Protected:   false,
			HandlerFunc: resetPasswdReq,
		},
		router.Route{
			Path:        "/resetpwd/{token}",
			Method:      "POST",
			Protected:   false,
			HandlerFunc: resetPasswd,
		},
	},
}
