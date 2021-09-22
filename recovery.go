package gee

import (
	"fmt"
	"net/http"

	"github.com/feiyuanmo/gee/log"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.InfoW("%s\n\n", message)
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
