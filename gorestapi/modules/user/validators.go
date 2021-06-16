package user

type updateUserFields struct {
	FirstName string `validate:"omitempty;max=3"`
	LastName  string `validate:"omitempty;length=2"`
	Role      int    `validate:"omitempty"`
}

func updateUserValidator() interface{} {
	return &updateUserFields{}
}
