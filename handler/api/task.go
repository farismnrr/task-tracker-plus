/**
 * Package api provides HTTP handlers for task-related operations.
 * 
 * Interfaces:
 * 
 * - TaskAPI: Interface defining methods for handling task-related HTTP requests.
 *   Methods:
 *   - AddTask: HTTP handler for adding a new task.
 *   - UpdateTask: HTTP handler for updating an existing task.
 *   - DeleteTask: HTTP handler for deleting a task.
 *   - GetTaskByID: HTTP handler for retrieving a task by its ID.
 *   - GetTaskList: HTTP handler for retrieving a list of all tasks.
 *   - GetTaskListByCategory: HTTP handler for retrieving a list of tasks by category.
 * 
 * Structs:
 * 
 * - taskAPI: Implements the TaskAPI interface. It provides HTTP handlers for task-related operations.
 *   Fields:
 *   - taskService: Instance of the TaskService interface to interact with the task service.
 *   Methods:
 *   - NewTaskAPI: Function to create a new instance of the taskAPI struct.
 *     Parameters:
 *     - taskRepo: Instance of the TaskService interface.
 *     Returns:
 *     - *taskAPI: A new instance of the taskAPI struct.
 *   - AddTask: HTTP handler for adding a new task.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - UpdateTask: HTTP handler for updating an existing task.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - DeleteTask: HTTP handler for deleting a task.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetTaskByID: HTTP handler for retrieving a task by its ID.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetTaskList: HTTP handler for retrieving a list of all tasks.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetTaskListByCategory: HTTP handler for retrieving a list of tasks by category.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 */

package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	task.ID = taskID
	err = t.taskService.Update(taskID, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	err = t.taskService.Delete(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	tasks, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	tasks, err := t.taskService.GetTaskCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}