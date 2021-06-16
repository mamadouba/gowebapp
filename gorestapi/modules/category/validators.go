package category

type createCategoryFields struct {
	Name string `validate:"required;regexp=^[a-zA-Z0-9]{2,20}$"`
}

func createCategoryValidator() interface{} {
	return &createCategoryFields{}
}
