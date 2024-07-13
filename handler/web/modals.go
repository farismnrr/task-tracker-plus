/** 
 * Package web provides functionality for serving web pages using the Gin web framework.
 * 
 * Interfaces:
 * 
 * - ModalWeb: Interface defining methods for serving modal pages.
 *   Methods:
 *   - Modal: Method for rendering modal pages.
 * 
 * Structs:
 * 
 * - modalWeb: Implements the ModalWeb interface and contains an embedded file system for serving static files.
 *   Fields:
 *   - embed: Embed.FS for embedding static files.
 *   Methods:
 *   - NewModalWeb: Function to initialize a new instance of the modalWeb struct.
 *     Parameters:
 *     - embed: Embed.FS for embedding static files.
 *     Returns:
 *     - *modalWeb: A new instance of the modalWeb struct.
 * 
 * Functions:
 * 
 * - Modal: HTTP handler function for rendering modal pages.
 *   Parameters:
 *   - c: Context provided by Gin framework.
 *   Description: This function retrieves the status and message query parameters from the context, 
 *     parses the modals.html and header.html templates, executes them, and renders the modal page. 
 *     If an error occurs during template execution, it returns the error message as a JSON response.
 */

package web

import (
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

type ModalWeb interface {
	Modal(c *gin.Context)
}

type modalWeb struct {
	embed embed.FS
}

func NewModalWeb(embed embed.FS) *modalWeb {
	return &modalWeb{embed}
}

func (m *modalWeb) Modal(c *gin.Context) {
	status := c.Query("status")
	message := c.Query("message")

	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "modals", "modals.html")

	var tmpl, err = template.ParseFS(m.embed, filepath, header)
	if err != nil {
		c.JSON(http.StatusSeeOther, err.Error())
		return
	}

	var dataTemplate = map[string]interface{}{
		"status":  status,
		"message": message,
	}

	err = tmpl.Execute(c.Writer, dataTemplate)
	if err != nil {
		c.JSON(http.StatusSeeOther, err.Error())
	}
}
