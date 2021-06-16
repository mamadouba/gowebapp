package middleware

import (
	"gorestapi/appcontext"
	"gorestapi/utils"
	"gorestapi/validator"
	"net/http"
)

func Validate(next appcontext.ContextHandlerFunc, fn appcontext.ValidatorFunc) appcontext.ContextHandlerFunc {
	return func(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
		payload := fn()
		if err := utils.DecodeJSON(w, r, &payload); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		ok, validationErrors := validator.Struct(payload)
		if !ok && len(validationErrors) != 0 {
			utils.RespondValidationError(w, validationErrors)
			return
		}
		ctx.Set(r, "payload", payload)
		next(ctx, w, r)
	}
}
