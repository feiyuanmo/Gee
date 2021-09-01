package gee

import (
	"net/http"
	"strings"

	"github.com/feiyuanmo/gee/log"
)

type HandlerFunc func(c *Context)

//父类
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

// type RouterGroup struct {
// 	prefix      string // 支持叠加
// 	router      *router
// }

//子类
type Engine struct {
	*RouterGroup //继承于父类
	router       *router
	groups       []*RouterGroup
}

// type Engine struct {
// 	*RouterGroup
// }

func New() *Engine {
	log.InfofW("------------new gee Engine------------")
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// func New() *Engine {
// 	engine := &Engine{}
// 	engine.RouterGroup = &RouterGroup{router: newRouter()}
// 	return engine
// }

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// func (group *RouterGroup) Group(prefix string) *RouterGroup {
// 	newGroup := &RouterGroup{
// 		prefix: group.prefix + prefix,
// 		router: group.router,
// 	}
// 	return newGroup
// }

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
// 	pattern := group.prefix + comp
// 	group.router.addRoute(method, pattern, handler)
// }

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// func (engine *Engine) addRouter(method, path string, handler HandlerFunc) {
// 	engine.router.addRoute(method, path, handler)
// }

// func (engine *Engine) GET(path string, handler HandlerFunc) {
// 	engine.addRouter("GET", path, handler)
// }

// func (engine *Engine) POST(path string, handler HandlerFunc) {
// 	engine.addRouter("POST", path, handler)
// }

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

// handler := (http.Handler)(engine)  手动转换为借口类型
// log.Fatal(http.ListenAndServe(":8080", handler))
func (engine *Engine) Run(addr string) error {
	log.InfofW("------------Run gee Engine:%s------------", addr)
	return http.ListenAndServe(addr, engine)
}
