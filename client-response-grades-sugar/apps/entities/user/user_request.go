package user

import (
	"time"
)

type UserRequest struct {
	ID int
	Name string `json:"name"`
	Email string `json:"email"`
	Phone_Number string `json:"phone_number"`
	Role string `json:"role"`
	Password string `json:"password"`
	Created_At time.Time
	Updated_At time.Time
}

