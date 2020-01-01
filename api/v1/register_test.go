package v1

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"outofmemory/errors"
	"testing"
)

func TestRegister(t *testing.T) {
	e := httpexpect.New(t, localServer)
	form := map[string]interface{}{
		"identifier":      "sekfung",
		"credential":      "sekfung",
		"identity_type":   "username",
		"identifier_from": 0,
	}
	e.POST("/v1/register").
		WithJSON(form).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeUserIsExist)

	form1 := map[string]interface{}{
		"identifier":      RandStringBytesMaskImprSrcUnsafe(10),
		"credential":      "sekfung",
		"identity_type":   "username",
		"identifier_from": 0,
	}
	e.POST("/v1/register").
		WithJSON(form1).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeNoError)

}