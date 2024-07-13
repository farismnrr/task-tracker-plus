/**
 * Package service provides interfaces and implementations for managing tasks.
 * 
 * Interfaces:
 * 
 * - TaskService: Interface defining methods for task management.
 *   Methods:
 *   - Store: Method to store a task.
 *   - Update: Method to update a task.
 *   - Delete: Method to delete a task.
 *   - GetByID: Method to retrieve a task by ID.
 *   - GetList: Method to retrieve a list of tasks.
 *   - GetTaskCategory: Method to retrieve tasks by category.
 * 
 * Structs:
 * 
 * - taskService: Struct implementing the TaskService interface.
 *   Fields:
 *   - taskRepository: Instance of repo.TaskRepository for task repository operations.
 *   Methods:
 *   - NewTaskService: Function to create a new instance of taskService.
 *   - Store: Method to store a task using the task repository.
 *   - Update: Method to update a task using the task repository.
 *   - Delete: Method to delete a task using the task repository.
 *   - GetByID: Method to retrieve a task by ID using the task repository.
 *   - GetList: Method to retrieve a list of tasks using the task repository.
 *   - GetTaskCategory: Method to retrieve tasks by category using the task repository.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (c *taskService) Store(task *model.Task) error {
	return c.taskRepository.Store(task)
}

func (s *taskService) Update(id int, task *model.Task) error {
	return s.taskRepository.Update(id, task)
}

func (s *taskService) Delete(id int) error {
	return s.taskRepository.Delete(id)
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	return s.taskRepository.GetByID(id)
}

func (s *taskService) GetList() ([]model.Task, error) {
	return s.taskRepository.GetList()
}

func (s *taskService) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	return s.taskRepository.GetTaskCategory(id)
}
