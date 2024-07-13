/**
 * Package main is the entry point for the server application.
 * 
 * This package initializes and starts an HTTP server using the Gin web framework.
 * It sets up API routes for user authentication, task management, and category management.
 * It also serves web client pages with embedded resources.
 *
 * Structs:
 * 
 * - APIHandler: Contains the API handlers for users, categories, and tasks.
 *   Fields:
 *   - UserAPIHandler: Handles user-related API requests.
 *   - CategoryAPIHandler: Handles category-related API requests.
 *   - TaskAPIHandler: Handles task-related API requests.
 *
 * - ClientHandler: Contains the web client handlers for authentication, home, dashboard, tasks, categories, and modals.
 *   Fields:
 *   - AuthWeb: Handles web authentication requests.
 *   - HomeWeb: Handles requests for the home page.
 *   - DashboardWeb: Handles requests for the dashboard page.
 *   - TaskWeb: Handles requests for the task page.
 *   - CategoryWeb: Handles requests for the category page.
 *   - ModalWeb: Handles requests for modals.
 *
 * Embedded Files:
 *
 * - Resources: Embeds the views directory for serving web client pages.
 *
 * Functions:
 *
 * - main: The main function that sets up and starts the HTTP server. It initializes the file-based database and configures the routes for both API and web client.
 *
 * - RunServer: Sets up the API routes. It initializes the repositories and services for users, categories, and tasks, and registers the respective routes.
 *   Parameters:
 *   - gin: The Gin engine instance.
 *   - filebasedDb: The file-based database instance.
 *   Returns:
 *   - *gin.Engine: The configured Gin engine instance.
 *
 * - RunClient: Sets up the web client routes. It initializes the client handlers for authentication, home, dashboard, tasks, categories, and modals, and registers the respective routes.
 *   Parameters:
 *   - gin: The Gin engine instance.
 *   - embed: The embedded file system instance.
 *   - filebasedDb: The file-based database instance.
 *   Returns:
 *   - *gin.Engine: The configured Gin engine instance.
 *
 * API Routes:
 * 
 * User Routes:
 * - POST /api/v1/user/login: Endpoint to handle user login. Expects a JSON payload with username and password. Returns a JSON response with user details and authentication token.
 * - POST /api/v1/user/register: Endpoint to handle user registration. Expects a JSON payload with user details such as username, password, and email. Returns a JSON response with the registered user's details.
 * - GET /api/v1/user/tasks: Protected endpoint to retrieve tasks associated with the logged-in user. Requires a valid authentication token. Returns a JSON response with the list of tasks categorized.
 * 
 * Task Routes:
 * - POST /api/v1/task/add: Protected endpoint to add a new task. Expects a JSON payload with task details. Returns a JSON response with the added task's details.
 * - GET /api/v1/task/get/:id: Protected endpoint to get a task by its ID. Requires a valid authentication token. Returns a JSON response with the task details.
 * - PUT /api/v1/task/update/:id: Protected endpoint to update a task by its ID. Expects a JSON payload with updated task details. Returns a JSON response with the updated task's details.
 * - DELETE /api/v1/task/delete/:id: Protected endpoint to delete a task by its ID. Requires a valid authentication token. Returns a JSON response indicating the success of the operation.
 * - GET /api/v1/task/list: Protected endpoint to get the list of all tasks. Requires a valid authentication token. Returns a JSON response with the list of tasks.
 * - GET /api/v1/task/category/:id: Protected endpoint to get tasks by category ID. Requires a valid authentication token. Returns a JSON response with the list of tasks in the specified category.
 * 
 * Category Routes:
 * - POST /api/v1/category/add: Protected endpoint to add a new category. Expects a JSON payload with category details. Returns a JSON response with the added category's details.
 * - GET /api/v1/category/get/:id: Protected endpoint to get a category by its ID. Requires a valid authentication token. Returns a JSON response with the category details.
 * - PUT /api/v1/category/update/:id: Protected endpoint to update a category by its ID. Expects a JSON payload with updated category details. Returns a JSON response with the updated category's details.
 * - DELETE /api/v1/category/delete/:id: Protected endpoint to delete a category by its ID. Requires a valid authentication token. Returns a JSON response indicating the success of the operation.
 * - GET /api/v1/category/list: Protected endpoint to get the list of all categories. Requires a valid authentication token. Returns a JSON response with the list of categories.
 * 
 * Web Client Routes:
 * 
 * Static Files:
 * - Serves static files from the "frontend/public" directory at the "/static" path.
 * 
 * Home Route:
 * - GET /client: Route to serve the home page.
 * 
 * User Routes:
 * - GET /client/login: Route to display the login page.
 * - POST /client/login/process: Route to process the login form. Expects form data with username and password. Redirects to the appropriate page based on the success of the login.
 * - GET /client/register: Route to display the registration page.
 * - POST /client/register/process: Route to process the registration form. Expects form data with user details such as username, password, and email. Redirects to the appropriate page based on the success of the registration.
 * - GET /client/logout: Protected route to log out the user. Redirects to the home page after logging out.
 * 
 * Main Routes:
 * - GET /client/dashboard: Protected route to display the dashboard page.
 * - GET /client/task: Protected route to display the task page.
 * - POST /client/task/add/process: Protected route to process the task addition form. Expects form data with task details. Redirects to the task page based on the success of the task addition.
 * - GET /client/category: Protected route to display the category page.
 * 
 * Modal Routes:
 * - GET /client/modal: Route to display a modal page.
 */

package main

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/handler/api"
	"a21hc3NpZ25tZW50/handler/web"
	"a21hc3NpZ25tZW50/middleware"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "embed"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type APIHandler struct {
	UserAPIHandler     api.UserAPI
	CategoryAPIHandler api.CategoryAPI
	TaskAPIHandler     api.TaskAPI
}

type ClientHandler struct {
	AuthWeb      web.AuthWeb
	HomeWeb      web.HomeWeb
	DashboardWeb web.DashboardWeb
	TaskWeb      web.TaskWeb
	CategoryWeb  web.CategoryWeb
	ModalWeb     web.ModalWeb
}

//go:embed views/*
var Resources embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode) //release

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())

		filebasedDb, err := filebased.InitDB()

		if err != nil {
			panic(err)
		}

		router = RunServer(router, filebasedDb)
		router = RunClient(router, Resources, filebasedDb)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	wg.Wait()
}

func RunServer(gin *gin.Engine, filebasedDb *filebased.Data) *gin.Engine {
	userRepo := repo.NewUserRepo(filebasedDb)
	sessionRepo := repo.NewSessionsRepo(filebasedDb)
	categoryRepo := repo.NewCategoryRepo(filebasedDb)
	taskRepo := repo.NewTaskRepo(filebasedDb)

	userService := service.NewUserService(userRepo, sessionRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	taskService := service.NewTaskService(taskRepo)

	userAPIHandler := api.NewUserAPI(userService)
	categoryAPIHandler := api.NewCategoryAPI(categoryService)
	taskAPIHandler := api.NewTaskAPI(taskService)

	apiHandler := APIHandler{
		UserAPIHandler:     userAPIHandler,
		CategoryAPIHandler: categoryAPIHandler,
		TaskAPIHandler:     taskAPIHandler,
	}

	version := gin.Group("/api/v1")
	{
		user := version.Group("/user")
		{
			user.POST("/login", apiHandler.UserAPIHandler.Login)
			user.POST("/register", apiHandler.UserAPIHandler.Register)

			user.Use(middleware.Auth())
			user.GET("/tasks", apiHandler.UserAPIHandler.GetUserTaskCategory)
		}

		task := version.Group("/task")
		{
			task.Use(middleware.Auth())
			task.POST("/add", apiHandler.TaskAPIHandler.AddTask)
			task.GET("/get/:id", apiHandler.TaskAPIHandler.GetTaskByID)
			task.PUT("/update/:id", apiHandler.TaskAPIHandler.UpdateTask)
			task.DELETE("/delete/:id", apiHandler.TaskAPIHandler.DeleteTask)
			task.GET("/list", apiHandler.TaskAPIHandler.GetTaskList)
			task.GET("/category/:id", apiHandler.TaskAPIHandler.GetTaskListByCategory)
		}

		category := version.Group("/category")
		{
			category.Use(middleware.Auth())
			category.POST("/add", apiHandler.CategoryAPIHandler.AddCategory)
			category.GET("/get/:id", apiHandler.CategoryAPIHandler.GetCategoryByID)
			category.PUT("/update/:id", apiHandler.CategoryAPIHandler.UpdateCategory)
			category.DELETE("/delete/:id", apiHandler.CategoryAPIHandler.DeleteCategory)
			category.GET("/list", apiHandler.CategoryAPIHandler.GetCategoryList)
		}
	}

	return gin
}

func RunClient(gin *gin.Engine, embed embed.FS, filebasedDb *filebased.Data) *gin.Engine {
	sessionRepo := repo.NewSessionsRepo(filebasedDb)
	sessionService := service.NewSessionService(sessionRepo)

	userClient := client.NewUserClient()
	taskClient := client.NewTaskClient()
	categoryClient := client.NewCategoryClient()

	authWeb := web.NewAuthWeb(userClient, sessionService, embed)
	modalWeb := web.NewModalWeb(embed)
	homeWeb := web.NewHomeWeb(embed)
	dashboardWeb := web.NewDashboardWeb(userClient, sessionService, embed)
	taskWeb := web.NewTaskWeb(taskClient, sessionService, embed)
	categoryWeb := web.NewCategoryWeb(categoryClient, sessionService, embed)

	client := ClientHandler{
		authWeb, homeWeb, dashboardWeb, taskWeb, categoryWeb, modalWeb,
	}

	gin.StaticFS("/static", http.Dir("frontend/public"))

	gin.GET("/", client.HomeWeb.Index)

	user := gin.Group("/client")
	{
		user.GET("/login", client.AuthWeb.Login)
		user.POST("/login/process", client.AuthWeb.LoginProcess)
		user.GET("/register", client.AuthWeb.Register)
		user.POST("/register/process", client.AuthWeb.RegisterProcess)

		user.Use(middleware.Auth())
		user.GET("/logout", client.AuthWeb.Logout)
	}

	main := gin.Group("/client")
	{
		main.Use(middleware.Auth())
		main.GET("/dashboard", client.DashboardWeb.Dashboard)
		main.GET("/task", client.TaskWeb.TaskPage)
		user.POST("/task/add/process", client.TaskWeb.TaskAddProcess)
		main.GET("/category", client.CategoryWeb.Category)
	}

	modal := gin.Group("/client")
	{
		modal.GET("/modal", client.ModalWeb.Modal)
	}

	return gin
}
