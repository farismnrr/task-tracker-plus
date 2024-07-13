/** 
 * Package repository provides interfaces and implementations for managing tasks.
 * 
 * Interfaces:
 * 
 * - TaskRepository: Interface defining methods for task data manipulation.
 *   Methods:
 *   - Store: Method to store a new task.
 *   - Update: Method to update an existing task.
 *   - Delete: Method to delete a task by ID.
 *   - GetByID: Method to retrieve a task by its ID.
 *   - GetList: Method to retrieve a list of all tasks.
 *   - GetTaskCategory: Method to retrieve a list of tasks by category.
 * 
 * Structs:
 * 
 * - taskRepository: Struct implementing the TaskRepository interface.
 *   Fields:
 *   - filebased: Instance of filebased.Data for file-based database operations.
 *   Methods:
 *   - NewTaskRepo: Function to create a new instance of taskRepository.
 *   - Store: Method to store a new task using file-based database operations.
 *   - Update: Method to update an existing task using file-based database operations.
 *   - Delete: Method to delete a task by ID using file-based database operations.
 *   - GetByID: Method to retrieve a task by its ID using file-based database operations.
 *   - GetList: Method to retrieve a list of all tasks using file-based database operations.
 *   - GetTaskCategory: Method to retrieve a list of tasks by category using file-based database operations.
 */

package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	return t.filebased.StoreTask(*task)
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	return t.filebased.UpdateTask(taskID, *task)
}

func (t *taskRepository) Delete(id int) error {
	return t.filebased.DeleteTask(id)
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	return t.filebased.GetTaskByID(id)
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	return t.filebased.GetTasks()
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	return t.filebased.GetTaskListByCategory(id)
}
