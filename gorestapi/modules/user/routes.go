package user

import (
	"gorestapi/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/users",
	SubRoutes: []router.Route{
		router.Route{
			Path:        "/",
			Method:      "GET",
			Protected:   true,
			HandlerFunc: list,
			Permission:  "user:read",
		},
		router.Route{
			Path:        "/{id}",
			Method:      "GET",
			Protected:   true,
			HandlerFunc: fetch,
			Permission:  "user:read",
		},
		router.Route{
			Path:        "/{id}",
			Method:      "PUT",
			Protected:   true,
			HandlerFunc: update,
			Validator:   updateUserValidator,
			Permission:  "user:write",
		},
		router.Route{
			Path:        "/{id}",
			Method:      "DELETE",
			Protected:   true,
			HandlerFunc: delete,
			Permission:  "user:write",
		},
		router.Route{
			Path:        "/me",
			Method:      "GET",
			Protected:   true,
			HandlerFunc: me,
			Permission:  "user:read",
		},
	},
}
