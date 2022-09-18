package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRouter(withTemplates bool) *gin.Engine {

	router := gin.Default()
	if withTemplates {
		router.LoadHTMLGlob("forms/*")
	}
	return router
}

func RunRouter(router *gin.Engine, host string, port int) {

	router.Run(fmt.Sprintf("%s:%d", host, port))
}
