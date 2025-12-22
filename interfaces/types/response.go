package types

// Response resp
type Response struct {
	StatusCode Code   `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// CodeOf code
func (r *Response) CodeOf(code Code) Response {
	if nil == r {
		r = new(Response)
	}
	r.StatusCode = code
	r.StatusMsg = code.Msg()
	return *r
}

// Success 成功
func (r *Response) Success() {
	r.CodeOf(CodeSuccess)
}
