/**
 * Package api provides HTTP handlers for category-related operations.
 * 
 * Interfaces:
 * 
 * - CategoryAPI: Interface defining methods for handling category-related HTTP requests.
 *   Methods:
 *   - AddCategory: HTTP handler for adding a new category.
 *   - UpdateCategory: HTTP handler for updating an existing category.
 *   - DeleteCategory: HTTP handler for deleting a category.
 *   - GetCategoryByID: HTTP handler for retrieving a category by its ID.
 *   - GetCategoryList: HTTP handler for retrieving a list of all categories.
 * 
 * Structs:
 * 
 * - categoryAPI: Implements the CategoryAPI interface. It provides HTTP handlers for category-related operations.
 *   Fields:
 *   - categoryService: Instance of the CategoryService interface to interact with the category service.
 *   Methods:
 *   - NewCategoryAPI: Function to create a new instance of the categoryAPI struct.
 *     Parameters:
 *     - categoryRepo: Instance of the CategoryService interface.
 *     Returns:
 *     - *categoryAPI: A new instance of the categoryAPI struct.
 *   - AddCategory: HTTP handler for adding a new category.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - UpdateCategory: HTTP handler for updating an existing category.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - DeleteCategory: HTTP handler for deleting a category.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetCategoryByID: HTTP handler for retrieving a category by its ID.
 *     Parameters:
 *     - c: Context object representing the HTTP request.
 *   - GetCategoryList: HTTP handler for retrieving a list of all categories.
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

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	category.ID = categoryID
	err = ct.categoryService.Update(categoryID, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category update success"})
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	err = ct.categoryService.Delete(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category delete success"})
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	categories, err := ct.categoryService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
