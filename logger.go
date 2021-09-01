package gee

import (
	"time"

	"github.com/feiyuanmo/gee/log"
)

func logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.InfofW("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
		//log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
