/**
 * Package repository provides interfaces and implementations for managing users.
 * 
 * Interfaces:
 * 
 * - UserRepository: Interface defining methods for user data manipulation.
 *   Methods:
 *   - GetUserByEmail: Method to retrieve a user by email.
 *   - CreateUser: Method to create a new user.
 *   - GetUserTaskCategory: Method to retrieve user task categories.
 * 
 * Structs:
 * 
 * - userRepository: Struct implementing the UserRepository interface.
 *   Fields:
 *   - filebasedDb: Instance of filebased.Data for file-based database operations.
 *   Methods:
 *   - NewUserRepo: Function to create a new instance of userRepository.
 *   - GetUserByEmail: Method to retrieve a user by email using file-based database operations.
 *   - CreateUser: Method to create a new user using file-based database operations.
 *   - GetUserTaskCategory: Method to retrieve user task categories using file-based database operations.
 */

package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	filebasedDb *filebased.Data
}

func NewUserRepo(filebasedDb *filebased.Data) *userRepository {
	return &userRepository{filebasedDb}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	return r.filebasedDb.GetUserByEmail(email)
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	return r.filebasedDb.CreateUser(user)
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	return r.filebasedDb.GetUserTaskCategory()
}
