package v1

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"outofmemory/errors"
	"testing"
)



func TestGetToken(t *testing.T) {
	e := httpexpect.New(t, localServer)
	// success
	form1 := map[string]interface{}{
		"identifier": "sekfung",
		"credential": "sekfung",
		"identity_type": "username",
		"identifier_from": 0,
	}
	e.POST("/v1/auth").
		WithJSON(form1).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeNoError)

	// not exist
	form2 := map[string]interface{}{
		"identifier": "sekfung1",
		"credential": "sekfung",
		"identity_type": "username",
	}
	e.POST("/v1/auth").
		WithJSON(form2).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeUserNotExist)

	// incorrect
	form3 := map[string]interface{}{
		"identifier": "sekfung",
		"credential": "sekfung1",
		"identity_type": "username",
	}
	e.POST("/v1/auth").
		WithJSON(form3).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeIncorrectIdentifierOrCredential)

	// bad request, miss field
	form4 := map[string]interface{}{
		"identifier": "sekfung",
		"identity_type": "username",
	}
	e.POST("/v1/auth").
		WithJSON(form4).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeInvalidParam)

	// bad request, invalid identifier field value
	form5 := map[string]interface{}{
		"identifier": RandStringBytesMaskImprSrcUnsafe(200),
		"identity_type": "username",
		"credential": "sekfung1",
		"identifier_from": 0,
	}
	e.POST("/v1/auth").
		WithJSON(form5).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeInvalidParam)

	// bad request, invalid identifier_from field value
	form6 := map[string]interface{}{
		"identifier": "sekfung",
		"identity_type": "username",
		"credential": "sekfung1",
		"identifier_from": 2,
	}
	e.POST("/v1/auth").
		WithJSON(form6).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeInvalidParam)

	// bad request, invalid identity_type field value
	form7 := map[string]interface{}{
		"identifier": "sekfung",
		"identity_type": "username1",
		"credential": "sekfung1",
		"identifier_from": 2,
	}
	e.POST("/v1/auth").
		WithJSON(form7).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object().
		ContainsKey("code").
		ValueEqual("code", errors.CodeInvalidParam)

}
