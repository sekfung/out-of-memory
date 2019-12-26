package v1

import (
	"github.com/gin-gonic/gin"
	"outofmemory/api"
	"outofmemory/errors"
	"outofmemory/models"
	"outofmemory/service"
	"outofmemory/utils"
	"strconv"
)

type articleForm struct {
	Title      string   `form:"title" validate:"required,min=1,max=50"`
	Content    string   `form:"content" validate:"required,min=15"`
	Tags       []uint32 `form:"tags[]" validate:"omitempty,unique"`
	Categories uint32   `form:"categories" validate:"omitempty,number,min=1"`
	Type       string   `form:"type" validate:"omitempty,oneof=md normal"`
	State      uint8    `form:"state" validate:"omitempty,oneof=publish deleted draft"`
}

// @Summary Add Article
// @Produce  json
// @Tags Article
// @Param Title query string true "Title, limit 50 character"
// @Param Content query string true "Content, not allow to be empty"
// @Param Tags query array false "Tags"
// @Param Category query int false "Category"
// @Param Type query string false "Article type, [md, doc], default value: md"
// @Param State query string false "Article State, [deleted, publish, draft], default value: publish"
// @Success 200 {string} json "{"code":200,"data":[],"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form articleForm
	)

	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}

	// default query publish article
	articleState := c.DefaultPostForm("state", "publish")
	form.State = models.ArticleState[articleState]
	// default query md article
	articleType := c.DefaultPostForm("type", "md")
	form.Type = articleType

	categoryReqId, err := strconv.ParseUint(c.DefaultPostForm("category", "1"), 10, 32)
	if err != nil {
		appG.Response(errors.CodeInvalidParam, nil)
		return
	}
	category := uint32(categoryReqId)
	articleService := service.Article{
		Title:    form.Title,
		Content:  form.Content,
		State:    form.State,
		Type:     articleType,
		AuthorID: utils.MustGetUIDFromHeader(c),
		Tags:     form.Tags,
		Category: category,
	}
	article, err := articleService.AddArticle()
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}
	appG.Response(errors.CodeNoError, article)
}

func UpdateArticle(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form articleForm
	)
	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}

}

func DeleteArticle(c *gin.Context) {
	var (
		appG = api.Gin{C: c}
		form articleForm
	)

	err := api.BindAndValid(c, &form)
	if err != nil {
		appG.Response(err.(*errors.AppError).ErrCode, nil)
		return
	}

}
