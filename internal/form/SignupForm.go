package form

type SignupErrors struct {
	Email    string
	Password string
	Confirm  string
}

type SignupForm struct {
	Email       string `form:"email"`
	Password    string `form:"password"`
	Confirm     string `form:"confirm"`
	FieldErrors SignupErrors
}
