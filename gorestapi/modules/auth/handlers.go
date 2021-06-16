package auth

import (
	"encoding/json"
	"fmt"
	"gorestapi/appcontext"
	"gorestapi/logger"
	"gorestapi/modules/user"
	"gorestapi/utils"
	"gorestapi/validator"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func register(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	params := ctx.Get(r)["payload"].(*registerFields)
	db := ctx.GetDB()
	var usr user.User
	if db.Where("email = ?", params.Email).First(&usr); usr.Id != "" {
		utils.RespondJSON(w, http.StatusConflict, "Email address already used")
		return
	}
	usr.Id = utils.GetUid()
	usr.Email = params.Email
	usr.HashPassword(params.Password)
	db.Create(&usr)
	args := struct {
		ConfirmUrl string
	}{
		ConfirmUrl: "http://foo.com",
	}
	go utils.SendMail(
		"templates/register.html",
		"Confirmation inscription",
		"mamadoubo.barry@gmail.com",
		args,
	)
	utils.RespondJSON(w, http.StatusCreated, &usr)
}

func login(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	params := ctx.Get(r)["payload"].(*loginFields)
	db := ctx.GetDB()
	var usr user.User
	if db.Where("email = ?", params.Email).First(&usr); usr.Id != "" {
		if usr.CheckPassword(params.Password) {
			var token Token
			jwtToken, err := token.GenerateJWT(db, usr.Id)
			if err != nil {
				logger.Error(err.Error())
				utils.RespondJSON(w, http.StatusInternalServerError, "Token generation failed")
				return
			}
			utils.RespondJSON(w, http.StatusOK, jwtToken)
			return
		}
	}
	utils.RespondJSON(w, http.StatusUnauthorized, "Authentication failed")
}

func refresh(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	refTokenStr := r.Header.Get("RefreshToken")
	var token Token
	if refTokenStr != "" {
		if db.Where("refresh_token = ?", refTokenStr).First(&token); token.UserId != "" {
			if token.Valid() {
				var usr user.User
				if db.Find(&usr, "id = ?", token.UserId); usr.Id != "" {
					var token Token
					jwtToken, err := token.GenerateJWT(db, usr.Id)
					if err != nil {
						logger.Error(err.Error())
						utils.RespondJSON(w, http.StatusInternalServerError, "Token generation failed")
						return
					}
					utils.RespondJSON(w, http.StatusOK, jwtToken)
					return
				}
			}
		}
		utils.RespondJSON(w, http.StatusBadRequest, "RefreshToken invalid")
	} else {
		utils.RespondJSON(w, http.StatusBadRequest, "No RefreshToken header found")
	}
}

func resetPasswdReq(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	params := mux.Vars(r)
	var usr user.User
	if db.Where("email = ?", params["email"]).First(&usr); usr.Id == "" {
		utils.RespondJSON(w, http.StatusUnauthorized, "Invalid user")
		return
	}
	var pwdResetToken PwdResetToken
	token := pwdResetToken.Generate(db, usr.Email)
	args := struct {
		Name         string
		ResetPassUrl string
	}{
		Name:         fmt.Sprintf("%s %s", usr.FirstName, usr.LastName),
		ResetPassUrl: fmt.Sprintf("%s/auth/resetpasswd/%s", os.Getenv("HOST"), token),
	}
	go utils.SendMail(
		"templates/resetPasswd.html",
		"Changement mot de passe",
		usr.Email,
		args,
	)
	utils.RespondJSON(w, http.StatusOK, "Password reset email sent")
}

func resetPasswd(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	params := mux.Vars(r)
	var pwdToken PwdResetToken
	if db.Where("token = ?", params["token"]).First(&pwdToken); pwdToken.Email == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "invalid token")
		return
	}
	if pwdToken.ExpireAt < time.Now().Unix() {
		utils.RespondJSON(w, http.StatusBadRequest, "token expired")
		return
	}

	var pwdReset PwdReset
	json.NewDecoder(r.Body).Decode(&pwdReset)
	ok, validationErrors := validator.Struct(pwdReset)
	if !ok && len(validationErrors) != 0 {
		utils.RespondValidationError(w, validationErrors)
		return
	}
	var usr user.User
	db.Where("email = ?", pwdToken.Email).First(&usr)
	usr.HashPassword(pwdReset.Password)
	db.Save(&usr)
	db.Delete(&pwdToken)
	utils.RespondJSON(w, http.StatusOK, "Password changed")
}
