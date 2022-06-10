package router

import (
	"ex/configuration"
	"ex/handler"

	"github.com/gin-gonic/gin"
)

type HTTPRouter struct {
	Configuration  configuration.EnvConfiguration
	ArticleHandler handler.Handler
}

func (r *HTTPRouter) Route() error {
	router := gin.Default()
	router.GET("/article", r.ArticleHandler.List)
	router.GET("/article/:id", r.ArticleHandler.Read)
	router.POST("/article", r.ArticleHandler.Add)
	router.PATCH("/article", r.ArticleHandler.Update)
	router.DELETE("/article", r.ArticleHandler.Remove)
	return router.Run(r.Configuration.Listen)
}
