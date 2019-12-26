package constants

type Gender int

const (
	UNKNOWN Gender = iota
	MALE
	FEMALE
)



var ArticleState = map[string]uint8{
	"deleted": 0,
	"publish": 1,
	"draft": 2,
}

