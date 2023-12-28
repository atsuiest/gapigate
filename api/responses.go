package api

import "github.com/atsuiest/gapigate/model"

var (
	res500 interface{} = nil
	res403 interface{} = nil
)

func init() {
	res500 = &model.Response{
		Code:    "KO500",
		Data:    nil,
		Error:   "Unavailable",
		Message: "Service temporary unavailable",
	}

	res403 = &model.Response{
		Code:    "E403",
		Error:   "Access Denied",
		Message: "The provided access token has no access to requested path",
	}
}
