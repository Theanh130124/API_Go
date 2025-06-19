package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social/common"
)

// minh se tu bat dc loi
func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok { //neu ok = true
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
				}
				//Van giu lai loi golang bat trong terminal
				panic(r)
			}
		}()
		c.Next()
	}
}
