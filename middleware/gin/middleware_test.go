package gin

import (
	limit "github.com/NotFound1911/rratelimit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"testing"
	"time"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	builder := MiddlewareBuilder{
		limiter: limit.NewFixWindowLimiter(rdb, "test", time.Second, 1),
	}
	r := gin.Default()
	r.Use(builder.Build())
	r.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "test"})
	})
	r.Run(":8000")
}
