package e

var MsgFlag = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	// 判断是否存在
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ERROR]
}
