package auth

type registerFields struct {
	Email    string `validate:"required;email"`
	Password string `validate:"required;length=5"`
}

func registerValidator() interface{} {
	return &registerFields{}
}

type loginFields struct {
	Email    string `validate:"required;email"`
	Password string `validate:"required"`
}

func loginValidator() interface{} {
	return &loginFields{}
}
