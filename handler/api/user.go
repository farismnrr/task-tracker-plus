/**
 * Package api provides HTTP handlers for user-related operations such as registration, login, and retrieval of user task categories.
 * 
 * Interfaces:
 * 
 * - UserAPI: Interface defining methods for handling user-related HTTP requests.
 *   Methods:
 *   - Register: HTTP handler for user registration.
 *   - Login: HTTP handler for user login.
 *   - GetUserTaskCategory: HTTP handler for retrieving user task categories.
 * 
 * Structs:
 * 
 * - userAPI: Implements the UserAPI interface. It provides HTTP handlers for user-related operations.
 *   Fields:
 *   - userService: Instance of the UserService interface to interact with the user service.
 *   Methods:
 *   - NewUserAPI: Function to create a new instance of the userAPI struct.
 *     Parameters:
 *     - userService: Instance of the UserService interface.
 *     Returns:
 *     - *userAPI: A new instance of the userAPI struct.
 *   - Register: HTTP handler for user registration.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - Login: HTTP handler for user login.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetUserTaskCategory: HTTP handler for retrieving user task categories.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 */

package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var user model.UserLogin

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("login data is empty"))
		return
	}

	var recordUser = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	tokenString, err := u.userService.Login(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	token, err := jwt.ParseWithClaims(*tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid token"))
		return
	}
	claims.ExpiresAt = expirationTime.Unix()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   *tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, model.NewSuccessResponse("login success"))
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	categories, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, categories)
}
