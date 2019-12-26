package api

import (
	"github.com/gin-gonic/gin"
	"outofmemory/errors"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return errors.ErrInvalidParam
	}
	validate = validator.New()
	err = validate.Struct(form)
	if err != nil {
		return errors.ErrInvalidParam
	}
	return nil
}
