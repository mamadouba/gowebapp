package middleware

import (
	"fmt"
	"gorestapi/appcontext"
	"gorestapi/config"
	"gorestapi/modules/user"
	"gorestapi/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Auth(next appcontext.ContextHandlerFunc, permission string) appcontext.ContextHandlerFunc {
	return func(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
		db := ctx.GetDB()
		authz := strings.Split(r.Header.Get("Authorization"), " ")[0]
		if authz != "" {
			token, err := jwt.Parse(authz, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.Configuration.SecretKey), nil
			})
			if err != nil {
				utils.RespondJSON(w, http.StatusUnauthorized, err.Error())
				return
			}
			if token.Valid {
				claims, _ := token.Claims.(jwt.MapClaims)
				var usr user.User
				db.Find(&usr, "id = ?", claims["sub"].(string))
				if usr.Id == "" {
					utils.RespondJSON(w, http.StatusUnauthorized, "User does not exist or migth be disabled")
					return
				}
				if permission != "" {

				}
				ctx.Set(r, "user", usr)
				next(ctx, w, r)
				return
			}
		}
		utils.RespondJSON(w, http.StatusBadRequest, "No Authorization header found")
	}
}
