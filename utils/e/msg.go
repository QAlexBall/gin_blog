package e

var msgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

// GetMsg get error information based on code
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
