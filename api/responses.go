package api

import "github.com/atsuiest/gapigate/model"

var (
	RES500 interface{} = nil
	RES503 interface{} = nil
	RES403 interface{} = nil
	RES404 interface{} = nil
)

func init() {
	RES500 = &model.Response{
		Code:    "KO500",
		Data:    nil,
		Error:   "Unavailable",
		Message: "Service temporary unavailable",
	}

	RES503 = &model.Response{
		Code:    "KO503",
		Data:    nil,
		Error:   "Bad Gateway",
		Message: "Bad Gateway",
	}

	RES403 = &model.Response{
		Code:    "KO403",
		Error:   "Access Denied",
		Message: "The provided access token has no access to requested path",
	}

	RES404 = &model.Response{
		Code:    "KO404",
		Error:   "Not Found",
		Message: "The provided route was not found",
	}
}
