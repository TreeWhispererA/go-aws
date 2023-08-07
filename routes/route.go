package routes

import (
	"blockparty.co/test/controllers"
	"github.com/gin-gonic/gin"
)

// routes => define endpoint
func Route(router *gin.Engine) {
	router.GET("/test", controllers.Test)
	router.GET("/token", controllers.GetTokens)
	router.GET("/token/:cid", controllers.GetTokenByID)
	router.GET("/scrap", controllers.ScrapFunc)
}
