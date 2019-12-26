package errors

var errMsg = map[int]string {
	// base
	CodeNoError: "success",
	CodeInternalServerError: "unknown error",
	CodeInvalidParam: "invalid param",
	CodeUnauthorized: "token is required",
	CodeForbidden: "not allow to access",
	// user
	CodeUserIsExist: "user is exist",
	CodeUserNotExist: "user is not exist",
	CodeIncorrectIdentifierOrCredential: "incorrect identifier or credential",
	// jwt
	CodeTokenTimeout: "token timeout",
	CodeCheckTokenFailed: "invalid token",
	CodeGenTokenFailed: "generate token failed",
	// tag
	CodeCreateTagFailed: "create tag failed",
	CodeUpdateTagFailed: "update tag failed",
	CodeDeleteTagFailed: "delete tag failed",
	CodeTagIsExist: "tag is exist",
	CodeTagNotExist: "tag is not exist",
	// article
	CodeCreateArticleFailed: "create article failed",
	CodeUpdateArticleFailed: "update article failed",
	CodeDeleteArticleFailed: "delete article failed",
	CodeGetArticleFailed: "get article failed",
	CodeArticleNotExist: "article is not exist",
	// category
	CodeCategoryIsExist: "category is exist",
	CodeCategoryNotExist: "category not exist",
	CodeCreateCategoryFailed: "create category failed",
	CodeUpdateCategoryFailed: "update category failed",
	CodeDeleteCategoryFailed: "delete category failed",
	CodeGetCategoryFailed: "get category failed",



}
