package models

import "time"

type CategoryArticle struct {
	Id         uint32 `json:"-"`
	AuthorId   uint32 `json:"author_id"`
	CategoryId uint32 `json:"category_id"`
	ArticleId  uint32 `json:"article_id"`
	CreatedAt      time.Time   `xorm:"created" json:"-"`
	UpdatedAt      time.Time   `xorm:"updated" json:"-"`
	DeletedAt      time.Time   `xorm:"deleted" json:"-"`
}

func mustGetCategoryForUser(uid uint32) []Category {
	var (
		categoryIDs = make([]uint32, 0)
		categories  = make([]Category, 0)
	)
	_, err := exitUserByUID(uid)
	if err != nil {
		return categories
	}
	err = engine.SQL("select category_id from category_article where author_id = ?", uid).Find(&categoryIDs)
	if err != nil {
		return categories
	}
	err = engine.SQL("select category_id, name from category where category_id in (?) and state = ?", categoryIDs, TagStateToUint["enable"]).Find(&categories)
	if err != nil {
		return categories
	}
	return categories
}
