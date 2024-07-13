/** 
 * Package web provides functionality for serving web pages related to tasks using the Gin web framework.
 * 
 * Interfaces:
 * 
 * - TaskWeb: Interface defining methods for handling task-related web functionalities.
 *   Methods:
 *   - TaskPage: Method for rendering the task page.
 *   - TaskAddProcess: Method for processing task addition requests.
 * 
 * Structs:
 * 
 * - taskWeb: Implements the TaskWeb interface and contains dependencies for handling task-related web functionalities.
 *   Fields:
 *   - taskClient: Instance of the TaskClient interface for communicating with the task service.
 *   - sessionService: Instance of the SessionService interface for managing user sessions.
 *   - embed: Embed.FS for embedding static files.
 *   Methods:
 *   - NewTaskWeb: Function to create a new instance of the taskWeb struct.
 *     Parameters:
 *     - taskClient: Instance of the TaskClient interface.
 *     - sessionService: Instance of the SessionService interface.
 *     - embed: Embed.FS for embedding static files.
 *     Returns:
 *     - *taskWeb: A new instance of the taskWeb struct.
 * 
 * Functions:
 * 
 * - TaskPage: HTTP handler function for rendering the task page.
 *   Parameters:
 *   - c: Context provided by Gin framework.
 *   Description: This function retrieves the user's email from the context, fetches the user's session, 
 *     retrieves the user's tasks, and renders the task page using a template, passing the retrieved tasks and user email as data. 
 *     If an error occurs during template execution, it redirects the user to a modal page with the error message.
 * 
 * - TaskAddProcess: HTTP handler function for processing task addition requests.
 *   Parameters:
 *   - c: Context provided by Gin framework.
 *   Description: This function retrieves the user's email from the context, fetches the user's session, 
 *     parses form data to create a new task, and adds the task using the task client. 
 *     It then redirects the user to the login page if the task addition is successful, 
 *     otherwise redirects to a modal page with an error message.
 */

package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

type TaskWeb interface {
	TaskPage(c *gin.Context)
	TaskAddProcess(c *gin.Context)
}

type taskWeb struct {
	taskClient     client.TaskClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewTaskWeb(taskClient client.TaskClient, sessionService service.SessionService, embed embed.FS) *taskWeb {
	return &taskWeb{taskClient, sessionService, embed}
}

func (t *taskWeb) TaskPage(c *gin.Context) {
	var email string
	if temp, ok := c.Get("email"); ok {
		if contextData, ok := temp.(string); ok {
			email = contextData
		}
	}

	session, err := t.sessionService.GetSessionByEmail(email)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	tasks, err := t.taskClient.TaskList(session.Token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	var dataTemplate = map[string]interface{}{
		"email": email,
		"tasks": tasks,
	}

	var funcMap = template.FuncMap{
		"exampleFunc": func() int {
			return 0
		},
	}

	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "main", "task.html")

	temp, err := template.New("task.html").Funcs(funcMap).ParseFS(t.embed, filepath, header)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	err = temp.Execute(c.Writer, dataTemplate)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
	}
}

func (t *taskWeb) TaskAddProcess(c *gin.Context) {
	var email string
	if temp, ok := c.Get("email"); ok {
		if contextData, ok := temp.(string); ok {
			email = contextData
		}
	}

	session, err := t.sessionService.GetSessionByEmail(email)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	priority, _ := strconv.Atoi(c.Request.FormValue("priority"))
	categoryID, _ := strconv.Atoi(c.Request.FormValue("category_id"))
	userID, _ := strconv.Atoi(c.Request.FormValue("user_id"))
	task := model.Task{
		Title:      c.Request.FormValue("title"),
		Deadline:   c.Request.FormValue("deadline"),
		Priority:   priority,
		Status:     c.Request.FormValue("status"),
		CategoryID: categoryID,
		UserID:     userID,
	}

	status, err := t.taskClient.AddTask(session.Token, task)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	if status == 201 {
		c.Redirect(http.StatusSeeOther, "/client/login")
	} else {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message=Add Task Failed!")
	}
}
