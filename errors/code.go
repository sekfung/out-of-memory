// Copyright 2019 sekfung
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package errors

// 错误码定义：
//

const (
	CodeNoError             = 200
	CodeInternalServerError = 500
	CodeInvalidParam        = 400
	CodeUnauthorized        = 401
	CodeForbidden           = 403
	// tag module error code
	CodeTagNotExist     = 10001
	CodeTagIsExist      = 10002
	CodeGetTagFailed    = 10003
	CodeCreateTagFailed = 10004
	CodeUpdateTagFailed = 10005
	CodeDeleteTagFailed = 10006
	// jwt
	CodeCheckTokenFailed = 20001
	CodeTokenTimeout     = 20002
	CodeGenTokenFailed   = 20003
	// user
	CodeIncorrectIdentifierOrCredential = 20004
	CodeUserIsExist                     = 20005
	CodeUserNotExist                    = 20006
	// article
	CodeCreateArticleFailed = 30001
	CodeUpdateArticleFailed = 30002
	CodeDeleteArticleFailed = 30003
	CodeGetArticleFailed    = 30004
	CodeArticleNotExist     = 30005
	// category
	CodeCategoryNotExist = 40001
	CodeCategoryIsExist = 40002
	CodeGetCategoryFailed = 40003
	CodeCreateCategoryFailed = 40004
	CodeUpdateCategoryFailed = 40005
	CodeDeleteCategoryFailed = 40006
	// comment
	CodeCommentNotExist = 50001
	CodeCommentIsExist = 50002
	CodeGetCommentFailed = 50003
	CodeCreateCommentFailed = 50004
	CodeUpdateCommentFailed = 50005
	CodeDeleteCommentFailed = 50006

)
