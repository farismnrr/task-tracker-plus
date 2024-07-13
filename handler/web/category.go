/**
 * Package web provides HTTP handlers for web-related operations such as category management.
 * 
 * Interfaces:
 * 
 * - CategoryWeb: Interface defining methods for handling web-based category management.
 *   Methods:
 *   - Category: HTTP handler for rendering category page.
 * 
 * Structs:
 * 
 * - categoryWeb: Implements the CategoryWeb interface. It provides HTTP handlers for web-based category management.
 *   Fields:
 *   - categoryClient: Instance of the CategoryClient interface for communicating with the category service.
 *   - sessionService: Instance of the SessionService interface for managing user sessions.
 *   - embed: Embed.FS for embedding static files.
 *   Methods:
 *   - NewCategoryWeb: Function to create a new instance of the categoryWeb struct.
 *     Parameters:
 *     - categoryClient: Instance of the CategoryClient interface.
 *     - sessionService: Instance of the SessionService interface.
 *     - embed: Embed.FS for embedding static files.
 *     Returns:
 *     - *categoryWeb: A new instance of the categoryWeb struct.
 * 
 * Functions:
 * 
 * - Category: HTTP handler function for rendering the category page.
 *   Parameters:
 *   - ctx: Context provided by Gin framework.
 *   Description: This function retrieves the user's email from the context, fetches the user's session, and then fetches the categories associated with the user. It then renders the category page using a template, passing the retrieved categories and user email as data.
 */

package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

type CategoryWeb interface {
	Category(c *gin.Context)
}

type categoryWeb struct {
	categoryClient client.CategoryClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewCategoryWeb(categoryClient client.CategoryClient, sessionService service.SessionService, embed embed.FS) *categoryWeb {
	return &categoryWeb{categoryClient, sessionService, embed}
}

func (c *categoryWeb) Category(ctx *gin.Context) {
	var email string
	if temp, ok := ctx.Get("email"); ok {
		if contextData, ok := temp.(string); ok {
			email = contextData
		}
	}

	session, err := c.sessionService.GetSessionByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	categories, err := c.categoryClient.CategoryList(session.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var dataTemplate = map[string]interface{}{
		"email":      email,
		"categories": categories,
	}

	var funcMap = template.FuncMap{
		"exampleFunc": func() int {
			return 0
		},
	}

	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "main", "category.html")

	t, err := template.New("category.html").Funcs(funcMap).ParseFS(c.embed, filepath, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	err = t.Execute(ctx.Writer, dataTemplate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}
}
