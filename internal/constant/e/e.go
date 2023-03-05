package e

const (
	SUCCESS            = 200
	ERROR              = 500
	INVALID_PARAMS     = 400
	AUTH_TOKEN_TIMEOUT = 1001
	TOKEN_ERROR        = 1002
)

var codeMsg = map[int]string{
	SUCCESS: "success",
	ERROR:   "failed",

	INVALID_PARAMS:     "invalid params",
	TOKEN_ERROR:        "invalid token",
	AUTH_TOKEN_TIMEOUT: "invalid, token",
}

// CodeMsg : get the code msg.
func CodeMsg(code int) (msg string) {
	msg, ok := codeMsg[code]
	if !ok {
		return codeMsg[ERROR]
	}
	return
}
