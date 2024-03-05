package v1

import (
	"encoding/json"
	"fmt"
	"ginmall/serializer"
)

func ErrorRespond(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg: "不匹配JSON类型",
			Error: fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg: "参数不匹配",
		Error: fmt.Sprint(err),
	}
}