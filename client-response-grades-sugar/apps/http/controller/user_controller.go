package controller

import (
	config "client-response-grades-sugar/config"
	infra "client-response-grades-sugar/infra"
	jwt "github.com/golang-jwt/jwt/v5"
	userEntity "client-response-grades-sugar/apps/entities/user"
	loginEntity "client-response-grades-sugar/apps/entities/login"

	"strconv"
	"time"
	// "reflect"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []userEntity.UserRequest
	data, err := config.DB.Query("SELECT * FROM users")

	if err != nil {
		return err
	}

	defer data.Close()

	for data.Next() {
		var user userEntity.UserRequest
		if err := data.Scan(&user.ID, &user.Name, &user.Email, &user.Phone_Number, &user.Role, &user.Password, &user.Created_At, &user.Updated_At); err != nil {
			return err
		}

		users = append(users, user)
	}

	var usersResponse []userEntity.UserResponseBody
	for _, user := range users {
		userResponse := userEntity.UserResponseBody{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Role:         user.Role,
			Created_At:   user.Created_At,
			Updated_At:   user.Updated_At,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	response := infra.ResponseAPIList("Success", "success", fiber.StatusOK, usersResponse, len(usersResponse))	
	return c.JSON(response)
}

func GetUserByRole(c *fiber.Ctx) error {
	var users []userEntity.UserRequest
	roleUser := c.Params("role")

	_, err := strconv.Atoi(roleUser)
	if err == nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "Bad request: Role must be a word", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}

	data, err := config.DB.Query("SELECT * FROM users WHERE role = ?", roleUser)

	if err != nil {
		return err
	}

	defer data.Close()

	for data.Next() {
		var user userEntity.UserRequest
		if err := data.Scan(&user.ID, &user.Name, &user.Email, &user.Phone_Number, &user.Role, &user.Password, &user.Created_At, &user.Updated_At); err != nil {
			return err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		responseFailed := infra.ResponseAPIFailed("Failed", "Not found", fiber.StatusNotFound)
    	return c.JSON(responseFailed)
	}

	var usersResponse []userEntity.UserResponseBody
	for _, user := range users {
		userResponse := userEntity.UserResponseBody{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Role:         user.Role,
			Created_At:   user.Created_At,
			Updated_At:   user.Updated_At,
		}
		usersResponse = append(usersResponse, userResponse)
	}
	
	response := infra.ResponseAPI("Success", "Succes to get user by role", fiber.StatusOK, usersResponse)
	return c.JSON(response)
}

func GetUserByName(c *fiber.Ctx) error {
	var users []userEntity.UserRequest
	nameUser := c.Params("name")

	_, err := strconv.Atoi(nameUser)
	if err == nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "Bad request: Role must be a word", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}

	data, err := config.DB.Query("SELECT * FROM users WHERE name LIKE ?", "%"+nameUser+"%")

	if err != nil {
		return err
	}

	defer data.Close()

	for data.Next() {
		var user userEntity.UserRequest
		if err := data.Scan(&user.ID, &user.Name, &user.Email, &user.Phone_Number, &user.Role, &user.Password, &user.Created_At, &user.Updated_At); err != nil {
			return err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		responseFailed := infra.ResponseAPIFailed("Failed", "Not found", fiber.StatusNotFound)
    	return c.JSON(responseFailed)
	}

	var usersResponse []userEntity.UserResponseBody
	for _, user := range users {
		userResponse := userEntity.UserResponseBody{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Role:         user.Role,
			Created_At:   user.Created_At,
			Updated_At:   user.Updated_At,
		}
		usersResponse = append(usersResponse, userResponse)
	}
	
	response := infra.ResponseAPI("Success", "Succes to get user by name", fiber.StatusOK, usersResponse)
	return c.JSON(response)
}

func GetUserByID(c *fiber.Ctx) error {
	var user userEntity.UserRequest
	idUser := c.Params("id")
	id, err := strconv.Atoi(idUser)

	if err != nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "Bad request: Id must be an number", fiber.StatusBadRequest)
    	return c.JSON(responseFailed)
	}

	data, err := config.DB.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer data.Close()

	if data.Next() {
		if err := data.Scan(&user.ID, &user.Name, &user.Email, &user.Phone_Number, &user.Role, &user.Password, &user.Created_At, &user.Updated_At); err != nil {
			return err
		}
	} else {
		responseFailed := infra.ResponseAPIFailed("Failed", "Not found", fiber.StatusNotFound)
		return c.JSON(responseFailed)
	}

	var userResponse userEntity.UserResponseBody
	userResponse = userEntity.UserResponseBody {
		ID: 			user.ID,
		Name:        	user.Name,
		Email:        	user.Email,
		Role:         	user.Role,
		Created_At:   	user.Created_At,
		Updated_At:   	user.Updated_At,
	}
	
	response := infra.ResponseAPI("Success", "Succes to get user by id", fiber.StatusOK, userResponse)
	return c.JSON(response)
}

func Register(c *fiber.Ctx) error {
	var user userEntity.UserRequest

	if err := c.BodyParser(&user); err != nil {
		responseFailed := infra.ResponseAPI("Failed", "error parsing request body", fiber.StatusBadRequest, nil)
		return c.JSON(responseFailed)
	}

	if user.Name == "" || user.Email == "" || user.Phone_Number == "" || user.Password == "" {
		responseFailed := infra.ResponseAPIFailed("Failed", "Please fill all field!", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	} 

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "error hash", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}
	user.Password = string(hashedPassword)

	if user.Role != "super user" && user.Role != "admin" && user.Role != "user" && user.Role != "" {
		responseFailed := infra.ResponseAPIFailed("Failed", "Role must be super user, admin, or user!", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}	

	if user.Role == "" {
		dataDefault, err := config.DB.Exec("INSERT INTO users(name, email, phone_number, role, password) VALUES(?, ?, ?, ?, ?)", user.Name, user.Email, user.Phone_Number, "user", user.Password)
		if err != nil {
			responseFailed := infra.ResponseAPIFailed("Failed", "Failed to request", fiber.StatusBadRequest)
			return c.JSON(responseFailed)
		}

		lastInsertIDDefault, err := dataDefault.LastInsertId()
			if err != nil {
				return err
			}

		user.ID = int (lastInsertIDDefault)
	} 	else {
			data, err := config.DB.Exec("INSERT INTO users(name, email, phone_number, role, password) VALUES(?, ?, ?, ?, ?)", user.Name, user.Email, user.Phone_Number, user.Role, user.Password)
			if err != nil {
				responseFailed := infra.ResponseAPIFailed("Failed", "Email or phone number already to register!", fiber.StatusBadRequest)
				return c.JSON(responseFailed)
			}

			lastInsertID, err := data.LastInsertId()
			if err != nil {
					return err
			}

			user.ID = int (lastInsertID)
	}

	var userResponse userEntity.UserResponseBody
	
	userResponse = userEntity.UserResponseBody {
		ID: 			user.ID,
		Name:        	user.Name,
		Email:        	user.Email,
		Role:       func() string {
						if user.Role == "" {
							return "user"
						}
						return user.Role
					}(),
		Created_At:   	time.Now(),
		Updated_At:   	time.Now(),
	}
	
	response := infra.ResponseAPI("Success", "Succes to created account", fiber.StatusOK, userResponse)
	return c.JSON(response)
}

func Login(c *fiber.Ctx) error  {
	var loginRequest loginEntity.LoginRequestParams
	if err := c.BodyParser(&loginRequest); err != nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "Cannot get request params", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}

	var user userEntity.UserRequest

	data, err := config.DB.Query("SELECT * FROM `users` WHERE email = ?", loginRequest.Email)
	if err != nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "Password or Email wrong!", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}

	defer data.Close()

	if data.Next(){
		if err := data.Scan(&user.ID, &user.Name, &user.Email, &user.Phone_Number, &user.Role, &user.Password, &user.Created_At, &user.Updated_At); err != nil {
			return err
		}
	} 

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		responseFailed := infra.ResponseAPIFailed("Failed", "password wrong", fiber.StatusBadRequest)
		return c.JSON(responseFailed)
	}

	claims := jwt.MapClaims{
		"email": loginRequest.Email,
		"role": loginRequest.Role,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	}

	expTime := time.Unix(claims["exp"].(int64), 0)
	expFormat := expTime.Format(time.RFC3339)

	tokenAdmin := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAdmin.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var loginResponse loginEntity.LoginResponseBody
	loginResponse = loginEntity.LoginResponseBody {
		ID: 			user.ID,
		Name:        	user.Name,
		Email:        	user.Email,
		Role:         	user.Role,
		Exp: 			expFormat,
		Token: 			token,
	}

	response := infra.ResponseAPI("Success", "Succes to login", fiber.StatusOK, loginResponse)
	return c.JSON(response)	
}

func SuperUserDashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "superuser" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	fmt.Println(role)

	return c.SendString("Welcome Superuser")
}

func PageUser(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	return c.Render("views/index.html", data)
}