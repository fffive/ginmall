package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorUserExist:             "用户已存在",
	ErrorFailEncrypte:          "密码加密失败",
	ErrorDatabase:              "数据库操作失败，请重试",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "账号或密码错误",
	ErrorAuthToken:             "Token签发失败",
	ErrorAuthCheckTokenTimeOut: "Token已过期",
	ErrorUploadFail:            "上传失败",
	ErrSendEmail:               "发送邮件失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[code]
	}
	return msg
}
