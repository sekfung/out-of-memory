package models

import "time"

type TagArticle struct {
	Id        uint32    `json:"-"`
	AuthorId  uint32    `json:"author_id"`
	TagId     uint32    `json:"tag_id"`
	ArticleId uint32    `json:"article_id"`
	CreatedAt time.Time `xorm:"created" json:"-"`
	UpdatedAt time.Time `xorm:"updated" json:"-"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
}

//
func mustGetTagsForUser(uid uint32) []Tag {
	var (
		tagIDs []TagArticle
		tags   = make([]Tag, 0)
	)
	_, err := exitUserByUID(uid)
	if err != nil {
		return tags
	}
	err = engine.Table("tag_article").Select("tag_id").Where("author_id = ?", uid).Find(&tagIDs)
	if err != nil {
		return tags
	}
	err = engine.SQL("select tag_id, name from tag where tag_id in (?) and state = ?", tagIDs, TagStateToUint["enable"]).Find(&tags)
	if err != nil {
		return tags
	}
	return tags
}

func Get() {

}
