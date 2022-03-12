package fin

import "net/http"

const (
	StatusInit = 0
	StatusOK   = 10000
	StatusWarn = 80006
)

type Render interface {
	Output(interface{})
}

type (
	Body struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Body   interface{} `json:"body"`
	}

	jsonApi struct {
		Context *Context
		Body    Body
	}
)

func (j *jsonApi) Output(mix interface{}) {
	j.Body.Status = StatusOK
	if val, ok := mix.(int); ok {
		j.Body.Status = val
		j.Body.Msg = Message(j.Body.Status)
		j.Body.Body = nil
	} else if val, ok := mix.(string); ok {
		j.Body.Status = 11000
		j.Body.Msg = val
		j.Body.Body = nil
	} else {
		j.Body.Msg = Message(j.Body.Status)
		j.Body.Body = mix
	}
	j.Context.JSON(http.StatusOK, j.Body)
	j.Context.Abort()
}

func Json(c *Context) *jsonApi {
	return &jsonApi{Context: c}
}
