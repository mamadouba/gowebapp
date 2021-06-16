package product

import (
	"gorestapi/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/products",
	SubRoutes: []router.Route{
		router.Route{
			Path:        "/",
			Method:      "GET",
			Protected:   true,
			Permission:  "product:write",
			HandlerFunc: listProduct,
		},
		router.Route{
			Path:        "/",
			Method:      "POST",
			Protected:   true,
			Permission:  "product:write",
			HandlerFunc: createProduct,
			Validator:   createProductValidator,
		},
		router.Route{
			Path:        "/{id}",
			Method:      "GET",
			Protected:   true,
			Permission:  "product:read",
			HandlerFunc: fetchProduct,
		},
		router.Route{
			Path:        "/{id}",
			Method:      "PUT",
			Protected:   true,
			Permission:  "product:write",
			HandlerFunc: updateProduct,
			Validator:   updateProductValidator,
		},
		router.Route{
			Path:        "/{id}",
			Method:      "DELETE",
			Protected:   true,
			Permission:  "product:write",
			HandlerFunc: deleteProduct,
		},
	},
}
