package main

import (
	"promg/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin,Content-Type, Content-Length,Authorization, accept, email"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))

	routes.UserRoutes(router)
	routes.QuotesRoutes(router)

	router.Run("localhost:9876")
}
