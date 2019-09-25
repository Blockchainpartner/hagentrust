package server

import (
	"github.com/Blockchainpartner/hagentrust/pkg/controllers"
	"github.com/Blockchainpartner/hagentrust/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	_, _ = util.EthClientInit()

	plainGroup := router.Group("plain")
	{
		plain := controllers.PlainController{}
		plainGroup.GET("/", plain.GetPlain)
	}

	return router
}