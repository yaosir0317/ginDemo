package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CostMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		fmt.Println(time.Since(start))
	}
}
