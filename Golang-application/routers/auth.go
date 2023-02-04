package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Aayush2609/week2-GL1-CipherSchools/Golang-application/handler"
	
)

func AuthRouter(router *gin.Engine, api handler.Handler) {
	router.POST("/signup", api.SignUp)
	router.POST("/signin", api.SignIn)
}