/** 
 * Package repository provides interfaces and implementations for data access operations related to categories.
 * 
 * Interfaces:
 * 
 * - CategoryRepository: Interface defining methods for category data manipulation.
 *   Methods:
 *   - Store: Method to store a new category.
 *   - Update: Method to update an existing category.
 *   - Delete: Method to delete a category.
 *   - GetByID: Method to retrieve a category by its ID.
 *   - GetList: Method to retrieve a list of all categories.
 * 
 * Structs:
 * 
 * - categoryRepository: Struct implementing the CategoryRepository interface.
 *   Fields:
 *   - filebasedDb: Instance of filebased.Data for file-based database operations.
 *   Methods:
 *   - NewCategoryRepo: Function to create a new instance of categoryRepository.
 *   - Store: Method to store a new category using file-based database operations.
 *   - Update: Method to update an existing category using file-based database operations.
 *   - Delete: Method to delete a category using file-based database operations.
 *   - GetByID: Method to retrieve a category by its ID using file-based database operations.
 *   - GetList: Method to retrieve a list of all categories using file-based database operations.
 */

package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	return c.filebasedDb.StoreCategory(*Category)
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	return c.filebasedDb.UpdateCategory(id, category)
}

func (c *categoryRepository) Delete(id int) error {
	return c.filebasedDb.DeleteCategory(id)
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	return c.filebasedDb.GetCategoryByID(id)
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	return c.filebasedDb.GetCategories()
}
