package server

import (
	"app/lesson4/config"
	_ "app/lesson4/docs" // 千万不要忘了导入把你上一步生成的docs
	"app/lesson4/pkg/logger"
	"app/lesson4/pkg/shutdown"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"strings"
	"time"
)

func Init(r *gin.Engine, callback func(engine *gin.Engine)) {

	r = gin.Default()
	r.Use(otelgin.Middleware(config.GetConfig().AppName))
	r.Use(logmiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	callback(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.TimeoutHandler(r, 10*time.Second, "服务超时"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	shutdown.Add(func() {
		logger.Logger.Infof("Shutdown gin Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			logger.Logger.Fatal("Server Shutdown:", err)
		}
		logger.Logger.Println("Server Shutdown success")
	})

}

func logmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if os.Getenv("DEPLOY_ENV") == "prd" {
		//	return
		//}
		path := c.Request.URL.Path
		//read request body content
		var bodyBytes []byte = []byte("binary")

		if strings.Contains(c.ContentType(), "json") {
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
			//因为HTTP的请求 Body，在读取过后会被置空，所以这里读取完后会重新赋值
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		method := c.Request.Method

		traceId := getOtlTraceId(c.Request.Context())
		spanId := getOtlSpanId(c.Request.Context())
		c.Set("trace", traceId)
		c.Set("span", spanId)

		logger.Logger.WithContext(c).Infof("req path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		//logger.Logger.Infof("[trace:%s, span:%s] req path: %s, Method: %s, body `%s`", traceId, spanId, path, method, string(bodyBytes))
		// Continue.
		c.Next()
	}
}

func getOtlTraceId(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}

func getOtlSpanId(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().SpanID().String()
}
