package msghandler

const (
	SUCCESS    = 200
	CREATED    = 201
	ACCEPTED   = 202
	NO_CONTENT = 203
	BAD_REQEST = 400
	ERROR      = 500
)

var Msgs = map[int]string{
	SUCCESS:    "success",
	CREATED:    "created",
	ACCEPTED:   "accepted",
	NO_CONTENT: "no content",
	BAD_REQEST: "bad request",
	ERROR:      "error",
}

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok {
		return msg
	}

	return Msgs[ERROR]
}
