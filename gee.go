package gee

import (
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type HandlerFunc func(c *Context)
type Engine struct {
	router *router
}

func New() *Engine {
	log.InfofW("------------new gee Engine------------")
	return &Engine{router: newRouter()}
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
