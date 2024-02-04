package user

import (
	"time"
)

type UserResponseBody struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}