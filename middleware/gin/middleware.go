package gin

import (
	limit "github.com/NotFound1911/rratelimit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MiddlewareBuilder struct {
	limiter *limit.RateLimit
}

func (m *MiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		allow, err := m.limiter.Limiter.Allow(context)
		if err != nil || !allow {
			context.AbortWithStatusJSON(http.StatusInternalServerError, "限流")
			return
		}
		context.Next()
	}
}
