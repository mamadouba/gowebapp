package user

import (
	"gorestapi/appcontext"
	"gorestapi/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func list(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var users []User
	db.Find(&users)
	utils.RespondJSON(w, http.StatusOK, &users)
}

func fetch(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	params := mux.Vars(r)
	var usr User
	if db.Find(&usr, "id = ?", params["id"]); usr.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, &usr)
}

func update(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	params := ctx.Get(r)["payload"].(*updateUserFields)
	db := ctx.GetDB()
	var usr User
	if db.Find(&usr, "id = ?", mux.Vars(r)["id"]); usr.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}
	db.Model(&usr).Updates(params)
	utils.RespondJSON(w, http.StatusCreated, &usr)
}

func delete(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	params := mux.Vars(r)
	var usr User
	if db.Find(&usr, "id = ?", params["id"]); usr.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}
	db.Delete(&usr)
	utils.RespondJSON(w, http.StatusOK, &usr)
}

func me(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	usr := ctx.Get(r)["user"].(*User)
	utils.RespondJSON(w, http.StatusOK, &usr)
}
