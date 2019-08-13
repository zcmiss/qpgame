package middlewares

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

//跨域中间件
func CorsMiddleware() (handler iris.Handler) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Origin", "Accept", "X-Requested-With", "PLATFORM", "Content-Type"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		Debug:            false,
	})
	return crs
}
