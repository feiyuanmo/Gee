package gee

import (
	"fmt"
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

type Engine struct{}

func New() *Engine {
	log.Info("------------new gee Engine------------")
	return &Engine{}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func (engine *Engine) Run(addr string) error {
	log.Infof("------------Run gee Engine:%s------------", addr)
	return http.ListenAndServe(addr, engine)
}
