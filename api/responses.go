package api

import "github.com/atsuiest/gapigate/model"

var (
	RES500 interface{} = nil
	RES403 interface{} = nil
)

func init() {
	RES500 = &model.Response{
		Code:    "KO500",
		Data:    nil,
		Error:   "Unavailable",
		Message: "Service temporary unavailable",
	}

	RES403 = &model.Response{
		Code:    "E403",
		Error:   "Access Denied",
		Message: "The provided access token has no access to requested path",
	}
}
