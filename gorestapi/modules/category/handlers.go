package category

import (
	"gorestapi/appcontext"
	"gorestapi/utils"
	"net/http"
)

func listCategory(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	db := ctx.GetDB()
	var categories []Category
	db.Find(&categories)
	utils.RespondJSON(w, http.StatusOK, &categories)
}

func createCategory(ctx *appcontext.AppContext, w http.ResponseWriter, r *http.Request) {
	data := ctx.Get(r)["payload"].(*createCategoryFields)
	db := ctx.GetDB()
	var category Category
	if db.Where("name = ?", data.Name).First(&category); category.Id != "" {
		utils.RespondJSON(w, http.StatusConflict, "Category name already used")
		return
	}
	category.Id = utils.GetUid()
	category.Name = data.Name
	db.Create(&category)
	utils.RespondJSON(w, http.StatusCreated, &category)
}
