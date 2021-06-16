package product

import (
	"gorestapi/appcontext"
	"gorestapi/modules/category"
	"gorestapi/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func listProduct(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var products []Product
	db.Find(&products)
	utils.RespondJSON(w, http.StatusOK, &products)
}

func createProduct(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	data := ctx.Get(r)["payload"].(*createProductFields)
	var prod Product
	if db.Where("title = ?", data.Title).First(&prod); prod.Id != "" {
		utils.RespondJSON(w, http.StatusConflict, "Product title already used")
		return
	}
	var category category.Category
	if db.Where("id = ?", data.CategoryId).First(&category); category.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "Product category does not exist")
		return
	}
	prod.Id = utils.GetUid()
	prod.Title = data.Title
	prod.Description = data.Description
	prod.Price = data.Price
	prod.CategoryId = data.CategoryId
	db.Create(&prod)
	utils.RespondJSON(w, http.StatusOK, &prod)
}

func fetchProduct(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var prod Product
	if db.Where("id = ?", mux.Vars(r)["id"]).First(&prod); prod.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "Product does not exist")
		return
	}
	utils.RespondJSON(w, http.StatusOK, &prod)
}

func updateProduct(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var prod Product
	if db.Where("id = ?", mux.Vars(r)["id"]).First(&prod); prod.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "Product does not exist")
		return
	}
	data := ctx.Get(r)["payload"].(*updateProductFields)
	db.Model(&prod).Updates(data)
	utils.RespondJSON(w, http.StatusOK, &prod)
}

func deleteProduct(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var prod Product
	if db.Where("id = ?", mux.Vars(r)["id"]).First(&prod); prod.Id == "" {
		utils.RespondJSON(w, http.StatusNotFound, "Product does not exist")
		return
	}
	db.Delete(&prod)
	utils.RespondJSON(w, http.StatusOK, &prod)
}
