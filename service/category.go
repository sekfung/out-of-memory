package service

import "outofmemory/models"

type Category struct {
	CategoryID uint32 `json:"category_id"`
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
}

func (c *Category) GetCategoryByID() (interface{}, error) {
	return models.GetCategoryById(c.CategoryID)
}

func (c *Category) GetCategories() (interface{}, error) {
	return models.GetCategories(c.Page, c.PerPage)
}

func (c *Category) AddCategory() (interface{}, error) {
	return models.AddCategory(c.Name, c.State)
}

func (c *Category) UpdateCategory() (interface{}, error) {
	categoryData := makeCategoryData(c)
	return models.UpdateCategoryById(categoryData)
}

func (c *Category) DeleteCategoryByID() error {
	return models.DeleteCategoryById(c.CategoryID)
}

func makeCategoryData(category *Category) map[string]interface{} {
	data := map[string]interface{}{
		"category_id": category.CategoryID,
		"name":        category.Name,
		"state":       category.State,
	}
	return data
}
