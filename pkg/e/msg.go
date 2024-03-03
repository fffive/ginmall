package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorUserExist:  "用户已存在",
	ErrFailEncrypte: "密码加密失败",
	ErrDatabase:     "数据库操作失败，请重试",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[code]
	}
	return msg
}
