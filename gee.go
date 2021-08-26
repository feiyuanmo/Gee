package gee

import (
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix string
	engine *Engine
}
type Engine struct {
	*RouterGroup
	router *router
}

func New() *Engine {
	log.InfofW("------------new gee Engine------------")
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) addRouter(method, path string, handler HandlerFunc) {
	engine.router.addRoute(method, path, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRouter("POST", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// handler := (http.Handler)(engine)  手动转换为借口类型
// log.Fatal(http.ListenAndServe(":8080", handler))
func (engine *Engine) Run(addr string) error {
	log.InfofW("------------Run gee Engine:%s------------", addr)
	return http.ListenAndServe(addr, engine)
}
