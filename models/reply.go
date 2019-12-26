package models

type Reply struct {
	BaseModel
	CommentId  uint32
	ReplyId    uint32
	ReplyType  uint8
	Content    string
	FromUid    uint32
	FromAvatar string
	ToUid      uint32
}
