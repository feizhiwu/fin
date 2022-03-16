package fin

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
)

type (
	MI map[string]interface{}
	MS map[string]string
	MU map[string]uint
	MF map[string]func()
	MG map[string]*gorm.DB
)

type HandlerFunc func(*Context)

type Engine struct {
	*routerGroup
	*router
	groups []*routerGroup
}

type routerGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *routerGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.routerGroup = &routerGroup{engine: engine}
	engine.groups = []*routerGroup{engine.routerGroup}
	engine.Use(recovery, logger)
	return engine
}

func Default() *Engine {
	engine := New()
	return engine
}

func (group *routerGroup) Group(prefix string) *routerGroup {
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}
	engine := group.engine
	newGroup := &routerGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *routerGroup) addRoute(method, path string, handler HandlerFunc) {
	if path == "" {
		path = "/"
	}
	Assert(path[0] == '/', "path must begin with '/'")
	pattern := group.prefix + path
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *routerGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
}

func (group *routerGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handler)
}

func (group *routerGroup) Put(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPut, pattern, handler)
}

func (group *routerGroup) Delete(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodDelete, pattern, handler)
}

func (group *routerGroup) Options(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodOptions, pattern, handler)
}

func (group *routerGroup) Any(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
	group.addRoute(http.MethodPost, pattern, handler)
	group.addRoute(http.MethodPut, pattern, handler)
	group.addRoute(http.MethodDelete, pattern, handler)
}

func (group *routerGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) SetConfig(path string) {
	setConfig(path)
}

func (engine *Engine) SetMessage(path string) {
	setMessage(path)
}

func (engine *Engine) SetLog(path string) {
	setLog(path)
}

func (engine *Engine) Run(address string) (err error) {
	defer func() { Assert(err == nil, err.Error()) }()
	Assert(configPath != "", "config path is not set")
	Assert(messagePath != "", "message path is not set")
	Assert(logPath != "", "log path is not set")
	connectDB()
	log.Printf("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
