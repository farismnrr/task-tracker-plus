/**
 * Package web provides HTTP handlers for web-related operations such as user authentication (login, register, logout).
 * 
 * Interfaces:
 * 
 * - AuthWeb: Interface defining methods for handling web-based user authentication.
 *   Methods:
 *   - Login: HTTP handler for rendering the login page.
 *   - LoginProcess: HTTP handler for processing user login.
 *   - Register: HTTP handler for rendering the registration page.
 *   - RegisterProcess: HTTP handler for processing user registration.
 *   - Logout: HTTP handler for user logout.
 * 
 * Structs:
 * 
 * - authWeb: Implements the AuthWeb interface. It provides HTTP handlers for web-based user authentication.
 *   Fields:
 *   - userClient: Instance of the UserClient interface for communicating with the user service.
 *   - sessionService: Instance of the SessionService interface for managing user sessions.
 *   - embed: Embed.FS for embedding static files.
 *   Methods:
 *   - NewAuthWeb: Function to create a new instance of the authWeb struct.
 *     Parameters:
 *     - userClient: Instance of the UserClient interface.
 *     - sessionService: Instance of the SessionService interface.
 *     - embed: Embed.FS for embedding static files.
 *     Returns:
 *     - *authWeb: A new instance of the authWeb struct.
 * 
 * Routes:
 * 
 * - /client/login: 
 *   - Method: GET
 *   - Handler: Login
 *   - Description: Renders the login page.
 * 
 * - /client/login: 
 *   - Method: POST
 *   - Handler: LoginProcess
 *   - Description: Processes user login.
 * 
 * - /client/register: 
 *   - Method: GET
 *   - Handler: Register
 *   - Description: Renders the registration page.
 * 
 * - /client/register: 
 *   - Method: POST
 *   - Handler: RegisterProcess
 *   - Description: Processes user registration.
 * 
 * - /client/logout: 
 *   - Method: GET
 *   - Handler: Logout
 *   - Description: Handles user logout.
 */

package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

type AuthWeb interface {
	Login(c *gin.Context)
	LoginProcess(c *gin.Context)
	Register(c *gin.Context)
	RegisterProcess(c *gin.Context)
	Logout(c *gin.Context)
}

type authWeb struct {
	userClient     client.UserClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewAuthWeb(userClient client.UserClient, sessionService service.SessionService, embed embed.FS) *authWeb {
	return &authWeb{userClient, sessionService, embed}
}

func (a *authWeb) Login(c *gin.Context) {
	var filepath = path.Join("views", "auth", "login.html")
	var header = path.Join("views", "general", "header.html")

	var tmpl, err = template.ParseFS(a.embed, filepath, header)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
	}
}

func (a *authWeb) LoginProcess(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	status, err := a.userClient.Login(email, password)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	session, err := a.sessionService.GetSessionByEmail(email)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	if status == 200 {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "session_token",
			Value:  session.Token,
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})

		c.Redirect(http.StatusSeeOther, "/client/dashboard")
	} else {
		c.Redirect(http.StatusSeeOther, "/client/login")
	}
}

func (a *authWeb) Register(c *gin.Context) {
	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "auth", "register.html")

	var tmpl, err = template.ParseFS(a.embed, filepath, header)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
	}
}

func (a *authWeb) RegisterProcess(c *gin.Context) {
	fullname := c.Request.FormValue("fullname")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	status, err := a.userClient.Register(fullname, email, password)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	if status == 201 {
		c.Redirect(http.StatusSeeOther, "/client/login")
	} else {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message=Register Failed!")
	}
}

func (a *authWeb) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", false, false)
	c.Redirect(http.StatusSeeOther, "/client/dashboard")
}
