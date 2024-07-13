/**
 * Package service provides interfaces and implementations for managing categories.
 * 
 * Interfaces:
 * 
 * - CategoryService: Interface defining methods for category management.
 *   Methods:
 *   - Store: Method to store a category.
 *   - Update: Method to update a category.
 *   - Delete: Method to delete a category.
 *   - GetByID: Method to retrieve a category by ID.
 *   - GetList: Method to retrieve a list of categories.
 * 
 * Structs:
 * 
 * - categoryService: Struct implementing the CategoryService interface.
 *   Fields:
 *   - categoryRepository: Instance of repo.CategoryRepository for category repository operations.
 *   Methods:
 *   - NewCategoryService: Function to create a new instance of categoryService.
 *   - Store: Method to store a category using the category repository.
 *   - Update: Method to update a category using the category repository.
 *   - Delete: Method to delete a category using the category repository.
 *   - GetByID: Method to retrieve a category by ID using the category repository.
 *   - GetList: Method to retrieve a list of categories using the category repository.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	return c.categoryRepository.Store(category)
}

func (c *categoryService) Update(id int, category model.Category) error {
	return c.categoryRepository.Update(id, category)
}

func (c *categoryService) Delete(id int) error {
	return c.categoryRepository.Delete(id)
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	return c.categoryRepository.GetByID(id)
}

func (c *categoryService) GetList() ([]model.Category, error) {
	return c.categoryRepository.GetList()
}
