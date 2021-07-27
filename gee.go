package gee

import (
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)
type Engine struct {
	router *router
}

func New() *Engine {
	log.Infof("------------new gee Engine------------")
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method, path string, handler HandlerFunc) {
	engine.router.addRoute(method, path, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.handle(w, req)
}

// handler := (http.Handler)(engine)  手动转换为借口类型
// log.Fatal(http.ListenAndServe(":8080", handler))
func (engine *Engine) Run(addr string) error {
	log.Infof("------------Run gee Engine:%s------------", addr)
	return http.ListenAndServe(addr, engine)
}
