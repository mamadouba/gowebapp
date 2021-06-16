package middleware

import "gorestapi/modules/user"

type AppCtx struct {
	authUser user.User
}
