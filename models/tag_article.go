package models

type TagArticle struct {
	BaseModel
	AuthorId  uint32    `json:"author_id"`
	TagId     uint32    `json:"tag_id"`
	ArticleId uint32    `json:"article_id"`
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
	err = db.Table("tag_article").Select("tag_id").Where("author_id = ?", uid).Find(&tagIDs).Error
	if err != nil {
		return tags
	}
	err = db.Exec("select tag_id, name from tag where tag_id in (?) and state = ?", tagIDs, TagStateToUint["enable"]).Find(&tags).Error
	if err != nil {
		return tags
	}
	return tags
}

func Get() {

}
