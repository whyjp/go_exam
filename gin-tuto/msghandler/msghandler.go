package msghandler

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

var Msgs = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "invalid parameter",
}

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok {
		return msg
	}

	return Msgs[ERROR]
}
