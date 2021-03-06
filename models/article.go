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
)

var ArticleState = map[string]uint8{
	"deleted": 0,
	"publish": 1,
	"draft":   2,
}

type Article struct {
	BaseModel
	ArticleId    uint32      `json:"aid"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	Tags         interface{} `gorm:"-" json:"tags"`
	Category     interface{} `gorm:"-" json:"category"`
	Comments     interface{} `gorm:"-" json:"comments"`
	Type         string      `json:"type"` // article type: md, normal
	State        uint8       `json:"-"`       // 0: deleted 1: publish 2: draft
	AuthorId     uint32      `json:"aid"`
	AuthorName   string      `json:"author_name"`
	AuthorAvatar string      `json:"author_avatar"`
}

func GetArticlesByTagID(tags []uint32, page, perPage int) ([]*Article, error) {
	// todo: only support one tag filter articles currently
	if len(tags) > 1 {
		return nil, errors.ErrInvalidParam
	}
	tagID := tags[0]
	//// check if tag is exist
	//err := ExistTagByID(tagID)
	//if err != nil {
	//	return nil, err
	//}
	// check if article by tag id
	_, err := ExistArticleByTagID(tagID)
	if err != nil {
		return nil, err
	}

	var articles []*Article
	err = db.Limit(perPage).Offset((page-1)*perPage).Find(articles).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return articles, nil
		default:
			return nil, errors.ErrGetArticleFailed

		}
	}
	return articles, nil
}

func GetArticlesByCategoryID(categoryID uint32, page, perPage int) ([]*Article, error) {
	var articles []*Article
	err := db.Where("category_id = ?", categoryID).Limit(perPage).Offset((page-1)*perPage).Find(articles).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return articles, nil
		default:
			return nil, errors.ErrGetArticleFailed

		}
	}
	return articles, nil
}

func GetArticlesByAuthorID(authorID uint32, page, perPage int) ([]*Article, error) {
	var articles []*Article
	err := db.Where("author_id = ?", authorID).Limit(perPage).Offset((page-1)*perPage).Find(articles).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.ErrGetArticleFailed

		}
	}
	return articles, nil
}

func GetArticleByID(articleID uint32) (*Article, error) {
	var article Article
	err := db.Where("article_id = ?", articleID).Find(&article).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return &article, nil
		default:
			return nil, errors.ErrGetArticleFailed

		}
	}
	return &article, nil
}

func AddArticle(data map[string]interface{}) (interface{}, error) {
	var (
		articleData = parseArticleData(data)
		article     Article
	)
	// generate a unique short article id
	articleID := uuid.New().ID()
	articleData.ArticleId = articleID

	// check if user exist and get user info
	authorID := articleData.AuthorId
	userInfo, err := getUserInfo(authorID)
	if err != nil {
		return nil, err
	}
	// add author info
	articleData.AuthorName = userInfo.Username
	articleData.AuthorAvatar = userInfo.AvatarUrl

	// check if tag is exist and get tag info
	tagIDs := make([]uint32, 0)
	tags := make([]interface{}, 0)
	for _, item := range articleData.Tags.([]Tag) {
		tagID := item.TagId
		tagIDs = append(tagIDs, tagID)
		tag, err := GetTagById(tagID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	// check if category is exist and get category info
	categoryID := articleData.Category.(Category).CategoryId
	category, err := GetCategoryById(categoryID)
	if err != nil {
		return nil, err
	}
	// assign
	article.Tags = tags
	article.Category = category
	session := db.Begin()
	defer session.Close()
	// insert article
	if err = session.Where("article_id = ?", articleID).Create(&articleData).Error; err != nil {
		_ = session.Rollback()
		return nil, errors.ErrCreateArticleFailed
	}

	// insert category article relation table
	var categoryArticle = CategoryArticle{
		CategoryId: categoryID,
		ArticleId:  articleID,
		AuthorId:   authorID,
	}
	err = session.Create(&categoryArticle).Error
	if err != nil {
		_ = session.Rollback()
		return nil, errors.ErrCreateArticleFailed
	}

	// batch insert tag article relation table
	if len(tagIDs) != 0 {
		var tagArticles []interface{}
		for _, tagID := range tagIDs {
			tagArticles = append(tagArticles, TagArticle{
				AuthorId:  userInfo.Uid,
				ArticleId: articleID,
				TagId:     uint32(tagID),
			})
		}
		err := session.Create(tagArticles).Error
		if err != nil {
			_ = session.Rollback()
			return nil, errors.ErrUnknownError
		}
	}
	_ = session.Commit()
	return &article, nil
}

func UpdateArticle(data map[string]interface{}) (*Article, error) {
	var (
		articleData = parseArticleData(data)
		article     Article
	)
	_, err := ExistArticleByID(articleData.ArticleId)
	if err != nil {
		return nil, err
	}
	err = db.Update(articleData).Error
	if err != nil {
		return nil, errors.ErrUpdateArticleFailed
	}
	return &article, nil
}

func DeleteArticle(data map[string]interface{}) error {
	var (
		articleData = parseArticleData(data)
	)
	_, err := ExistArticleByID(articleData.ArticleId)
	if err != nil {
		return nil
	}
	err = db.Delete(&articleData).Error
	return err
}

func ExistArticleByID(articleID uint32) (*Article, error) {
	var article Article
	err := db.Select("article_id").Where("article_id = ?", articleID).Find(&article).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrArticleNotExist
		default:
			return nil, errors.ErrGetArticleFailed
		}
	}
	return &article, nil
}

func ExistArticleByTagID(tagID uint32) (*TagArticle, error) {
	var tagArticle TagArticle
	err := db.Select("tag_id").Where("tag_id = ?", tagID).Find(&tagArticle).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrArticleNotExist
		default:
			return nil, errors.ErrGetArticleFailed
		}
	}
	return &tagArticle, nil
}

func ExistArticleByCategoryID(categoryID uint32) (*Article, error) {
	var article Article
	err := db.Select("category_id").Where("category_id = ?", categoryID).Find(&article).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrArticleNotExist
		default:
			return nil, errors.ErrUnknownError
		}
	}
	return &article, nil
}

func parseArticleData(data map[string]interface{}) Article {
	category := Category{CategoryId: data["category"].(uint32)}
	tags := make([]Tag, 0)
	tagsData := data["tags"].([]uint32)
	for _, tagID := range tagsData {
		tags = append(tags, Tag{TagId: tagID})
	}
	article := Article{
		ArticleId:    data["article_id"].(uint32),
		Title:        data["title"].(string),
		Content:      data["content"].(string),
		Type:         data["type"].(string),
		State:        data["state"].(uint8),
		AuthorId:     data["author_id"].(uint32),
		AuthorName:   data["author_name"].(string),
		AuthorAvatar: data["author_avatar"].(string),
		Tags:         tags,
		Category:     category,
	}
	return article
}
