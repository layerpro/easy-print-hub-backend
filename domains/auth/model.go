package auth

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SignIn struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type ResponseSignIn struct {
	AccessToken string `json:"access_token"`
}
