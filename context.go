package fin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type H map[string]interface{}

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]interface{}
	handlers   []HandlerFunc
	index      int8
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:    req.URL.Path,
		Method:  req.Method,
		Request: req,
		Writer:  w,
		index:   -1,
	}
}

// Next 调用Next方法，相当于在该handler中调用下个handler
//利用Next，可以达到使中间件在请求前后各实现一些行为的功能
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) SetParams(key string, value interface{}) {
	c.Params[key] = value
}

func (c *Context) GetParams() map[string]interface{} {
	raw, _ := c.GetRawData()
	if len(raw) == 0 {
		c.Request.ParseForm()
		for k, v := range c.Request.Form {
			c.Params[k] = v[0]
		}
	} else {
		json.Unmarshal(raw, &c.Params)
	}
	json.Unmarshal(raw, &c.Params)
	return c.Params
}

func (c *Context) Param(key string) interface{} {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) GetRawData() ([]byte, error) {
	return ioutil.ReadAll(c.Request.Body)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Header(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) GetHeader(key string) string {
	return c.requestHeader(key)
}

func (c *Context) requestHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Header("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.Header("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) NewDisplay() *Display {
	c.GetParams()
	d := &Display{
		Context: c,
		funcs:   make(MF),
	}
	return d
}
