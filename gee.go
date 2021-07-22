package gee

import (
	"fmt"
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	log.Info("------------new gee Engine------------")
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRouter("GET", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	log.Infof("IP:%s Method:%s Path:%s", req.Host, req.Method, req.URL.Path)
	hadler, ok := engine.router[key]
	if ok {
		hadler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// handler := (http.Handler)(engine)  手动转换为借口类型
// log.Fatal(http.ListenAndServe(":8080", handler))
func (engine *Engine) Run(addr string) error {
	log.Infof("------------Run gee Engine:%s------------", addr)
	return http.ListenAndServe(addr, engine)
}
