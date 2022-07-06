package router

import (
	"report-lkl-morning/api"

	"github.com/gin-gonic/gin"
)

func NewRouter(h api.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/getall", h.GetAllReport)
	router.POST("/postinfo", h.Postinfo)
	router.DELETE("/delinfo", h.Delinfo)
	router.GET("/get", h.GetReport)
	return router
}
