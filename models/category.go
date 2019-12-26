// Copyright 2019 sekfung
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"outofmemory/errors"
	"time"
)

var CategoryStateToUint = map[string]uint8{
	"disable": 0,
	"enable":  1,
	"deleted": 2,
}

type Category struct {
	Id         uint32 `json:"-"`
	CategoryId uint32 `json:"category_id"`
	Name       string `json:"name"`
	State      uint8  `json:"-"` // 0: disable; 1: enable 2: deleted
	CreatedAt  time.Time   `xorm:"created" json:"-"`
	UpdatedAt  time.Time   `xorm:"updated" json:"-"`
	DeletedAt  time.Time   `xorm:"deleted" json:"-"`
}

func AddCategory(name string, state uint8) (interface{}, error) {
	_, err := existCategoryByName(name)
	// category is exist
	if err != nil && err != errors.ErrCategoryNotExist {
		return nil, err
	}
	// generate a short unique id
	cid := uuid.New().ID()
	category := Category{
		CategoryId: cid,
		Name:       name,
		State:      state,
	}
	_, err = engine.Insert(&category)
	if err != nil {
		return nil, errors.ErrCreateCategoryFailed
	}
	return &category, nil
}

func GetCategoryById(categoryId uint32) (interface{}, error) {
	return existCategoryByID(categoryId)
}

func GetCategories(page, perPage int) (interface{}, error) {
	var categories []*Category
	err := engine.Where("state <= ?", CategoryStateToUint["enable"]).Limit((page-1)*perPage, perPage).Find(&categories)
	if err != nil {
		return nil, errors.ErrGetCategoryFailed
	}
	return categories, nil
}

func DeleteCategoryById(tagIds interface{}) error {
	var err error
	switch inst := tagIds.(type) {
	case uint32:
		err = deleteCategoryById(inst)
		break
	case []uint32:
		for _, categoryID := range inst {
			err = deleteCategoryById(categoryID)
		}
		break
	}
	return err
}

func deleteCategoryById(categoryID uint32) error {
	_, err := existCategoryByID(categoryID)
	if err != nil {
		return err
	}
	// soft delete
	_, err = engine.Table("category").Where("category_id= ?", categoryID).Update("state", CategoryStateToUint["deleted"])
	if err != nil {
		return errors.ErrDeleteCategoryFailed
	}
	return nil
}

func UpdateCategoryById(data map[string]interface{}) (interface{}, error) {
	categoryData := parseCategoryData(data)
	_, err := existCategoryByName(categoryData.Name)
	// make sure category name is not exist
	if err == nil {
		return nil, errors.ErrCategoryIsExist
	}
	// make sure category id is exist
	_, err = existCategoryByID(categoryData.CategoryId)
	if err != nil {
		return nil, err
	}
	_, err = engine.Table("category").Where("category_id = ?", categoryData.CategoryId).Update(categoryData)
	if err != nil {
		return nil, errors.ErrUpdateCategoryFailed
	}
	return &categoryData, nil
}

func existCategoryByID(categoryID uint32) (interface{}, error) {
	var category Category
	err := engine.Where("category_id = ? and state <= ?", categoryID, CategoryStateToUint["enable"]).Find(&category)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrCategoryNotExist
		default:
			return nil, errors.ErrGetCategoryFailed
		}
	}
	return &category, nil
}

func existCategoryByName(name string) (interface{}, error) {
	var category Category
	err := engine.Where("name = ? and state <= ?", name, CategoryStateToUint["enable"]).Find(&category)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrCategoryNotExist
		default:
			return nil, errors.ErrGetCategoryFailed
		}
	}
	return &category, nil
}

func parseCategoryData(data map[string]interface{}) Category {
	category := Category{
		CategoryId: data["category_id"].(uint32),
		Name:       data["name"].(string),
		State:      data["state"].(uint8),
	}
	return category
}
