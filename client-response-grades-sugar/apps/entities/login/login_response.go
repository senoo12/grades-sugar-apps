package login

type LoginResponseBody struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Exp string `json:"exp"`
	Token string `json:"token"`
}