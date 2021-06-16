package category

import (
	"gorestapi/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/categories",
	SubRoutes: []router.Route{
		router.Route{
			Path:        "/",
			Method:      "GET",
			Protected:   true,
			Permission:  "category:read",
			HandlerFunc: listCategory,
		},
		router.Route{
			Path:        "/",
			Method:      "POST",
			Protected:   true,
			Permission:  "category:read",
			HandlerFunc: createCategory,
			Validator:   createCategoryValidator,
		},
	},
}
