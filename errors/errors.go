package errors


var (
	// base
	ErrUnknownError = &AppError{GetMsg(CodeInternalServerError), CodeInternalServerError}
	ErrInvalidParam = &AppError{GetMsg(CodeInvalidParam), CodeInvalidParam}
	// user
	ErrUserNotExist = &AppError{GetMsg(CodeUserNotExist), CodeUserNotExist}
	ErrIncorrectIdentifierOrCredential = &AppError{GetMsg(CodeIncorrectIdentifierOrCredential), CodeIncorrectIdentifierOrCredential}
	// jwt
	ErrGenTokenFailed = &AppError{GetMsg(CodeGenTokenFailed), CodeGenTokenFailed}
	ErrTokenTimeout = &AppError{GetMsg(CodeTokenTimeout), CodeTokenTimeout}
	ErrCheckTokenFailed = &AppError{GetMsg(CodeCheckTokenFailed), CodeCheckTokenFailed}
	// tag
	ErrCreateTagFailed = &AppError{GetMsg(CodeCreateTagFailed), CodeCreateTagFailed}
	ErrUpdateTagFailed = &AppError{GetMsg(CodeUpdateTagFailed), CodeUpdateTagFailed}
	ErrDeleteTagFailed = &AppError{GetMsg(CodeDeleteTagFailed), CodeDeleteTagFailed}
	ErrTagIsExist = &AppError{GetMsg(CodeTagIsExist), CodeTagIsExist}
	ErrTagNotExist = &AppError{GetMsg(CodeTagNotExist), CodeTagNotExist}
	ErrGetTagFailed = &AppError{GetMsg(CodeGetTagFailed), CodeGetTagFailed}
	// article
	ErrCreateArticleFailed = &AppError{GetMsg(CodeCreateArticleFailed), CodeCreateArticleFailed}
	ErrUpdateArticleFailed = &AppError{GetMsg(CodeUpdateArticleFailed), CodeUpdateArticleFailed}
	ErrDeleteArticleFailed = &AppError{GetMsg(CodeDeleteArticleFailed), CodeDeleteArticleFailed}
	ErrGetArticleFailed = &AppError{GetMsg(CodeGetArticleFailed), CodeGetArticleFailed}
	ErrArticleNotExist = &AppError{GetMsg(CodeArticleNotExist), CodeArticleNotExist}
	// category
	ErrCategoryNotExist = &AppError{GetMsg(CodeCategoryNotExist), CodeCategoryNotExist}
	ErrCategoryIsExist = &AppError{GetMsg(CodeCategoryIsExist), CodeCategoryIsExist}
	ErrGetCategoryFailed = &AppError{GetMsg(CodeGetCategoryFailed), CodeGetCategoryFailed}
	ErrUpdateCategoryFailed = &AppError{GetMsg(CodeUpdateCategoryFailed), CodeUpdateCategoryFailed}
	ErrDeleteCategoryFailed = &AppError{GetMsg(CodeDeleteCategoryFailed), CodeDeleteCategoryFailed}
	ErrCreateCategoryFailed = &AppError{GetMsg(CodeCreateCategoryFailed), CodeCreateCategoryFailed}

	// comment
	ErrCommentNotExist = &AppError{GetMsg(CodeCommentNotExist), CodeCommentNotExist}
	ErrCommentIsExist = &AppError{GetMsg(CodeCommentIsExist), CodeCommentIsExist}
	ErrGetCommentFailed = &AppError{GetMsg(CodeGetCommentFailed), CodeGetCommentFailed}
	ErrCreateCommentFailed = &AppError{GetMsg(CodeCreateCommentFailed), CodeCreateCommentFailed}
	ErrUpdateCommentFailed = &AppError{GetMsg(CodeUpdateCommentFailed), CodeUpdateCommentFailed}
	ErrDeleteCommentFailed= &AppError{GetMsg(CodeDeleteCommentFailed), CodeDeleteCommentFailed}
)

type AppError struct {
	ErrMsg string
	ErrCode int
}


func (err *AppError) Error() string {
	return err.ErrMsg
}


func GetMsg(errCode int) string {
	return errMsg[errCode]
}
