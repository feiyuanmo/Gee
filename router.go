package gee

import (
	"fmt"
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	log.Infof("Route %4s - %s", method, path)
	key := method + path
	r.handlers[key] = handler
}

func (r *router) handle(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/favicon.ico" {

	} else {
		key := req.Method + "-" + req.URL.Path
		log.Infof("IP:%s Method:%s Path:%s", req.Host, req.Method, req.URL.Path)
		hadler, ok := r.handlers[key]
		if ok {
			hadler(w, req)
		} else {
			fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		}
	}
}
