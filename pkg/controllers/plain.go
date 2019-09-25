package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlainController struct {}

func (a PlainController) GetPlain(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{"error": nil, "plain": "Plain"})
}