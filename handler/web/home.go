/** 
 * Package web provides functionality for serving web pages using the Gin web framework.
 * 
 * Interfaces:
 * 
 * - HomeWeb: Interface defining methods for serving home pages.
 *   Methods:
 *   - Index: Method for rendering the index page.
 * 
 * Structs:
 * 
 * - homeWeb: Implements the HomeWeb interface and contains an embedded file system for serving static files.
 *   Fields:
 *   - embed: Embed.FS for embedding static files.
 *   Methods:
 *   - NewHomeWeb: Function to initialize a new instance of the homeWeb struct.
 *     Parameters:
 *     - embed: Embed.FS for embedding static files.
 *     Returns:
 *     - *homeWeb: A new instance of the homeWeb struct.
 * 
 * Functions:
 * 
 * - Index: HTTP handler function for rendering the index page.
 *   Parameters:
 *   - c: Context provided by Gin framework.
 *   Description: This function parses the index.html and header.html templates, executes them, and renders the index page. 
 *     If an error occurs during template execution, it redirects the user to a modal page with the error message.
 */

package web

import (
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

type HomeWeb interface {
	Index(c *gin.Context)
}

type homeWeb struct {
	embed embed.FS
}

func NewHomeWeb(embed embed.FS) *homeWeb {
	return &homeWeb{embed}
}

func (h *homeWeb) Index(c *gin.Context) {
	var filepath = path.Join("views", "main", "index.html")
	var header = path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(h.embed, filepath, header))

	err := tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}
}
