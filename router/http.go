package router

import (
	"ex/configuration"
	"ex/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPRouter struct {
	Configuration configuration.EnvConfiguration
	Handlers      []handler.Handler
}

func (r *HTTPRouter) Route() error {
	router := gin.Default()
	router.GET("/article", func(c *gin.Context) { c.JSON(http.StatusOK, "articles") })
	router.GET("/article/:id", func(c *gin.Context) { c.JSON(http.StatusOK, "article") })
	router.POST("/article", func(c *gin.Context) { c.JSON(http.StatusOK, "add") })
	router.PATCH("/article", func(c *gin.Context) { c.JSON(http.StatusOK, "update") })
	router.DELETE("/article", func(c *gin.Context) { c.JSON(http.StatusOK, "remove") })
	return router.Run(r.Configuration.Listen)
}
