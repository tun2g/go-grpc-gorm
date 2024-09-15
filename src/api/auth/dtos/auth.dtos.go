package auth

type LoginParamsDto struct {
	Email    string
	Password string
}

type RegisterParamsDto struct {
	Email    string
	Password string
	FullName string
}

type TokenResDto struct {
	AccessToken  string
	RefreshToken string
}

type UserResDto struct {
	ID       string
	FullName string
	Email    string
}

type AuthResDto struct {
	Tokens TokenResDto
	User   UserResDto
}
