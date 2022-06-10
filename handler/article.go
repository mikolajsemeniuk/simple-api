package handler

import (
	"encoding/json"
	"ex/article"
	"ex/storage"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ArticleHandler struct {
	Redis storage.RedisStorage
}

func (h *ArticleHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, "list")
}

func (h *ArticleHandler) Read(c *gin.Context) {
	c.JSON(http.StatusOK, "read")
}

func (h *ArticleHandler) Add(c *gin.Context) {
	// we want to validate data
	article := article.Article{}

	if err := json.NewDecoder(c.Request.Body).Decode(&article); err != nil {
		c.JSON(http.StatusBadRequest, "error decoding json")
		return
	}

	if !article.Valid() {
		c.JSON(http.StatusBadRequest, "article is not valid")
		return
	}

	// we want to get cache if cache exists return it save it
	result, err := h.Redis.Read("one")
	if err != redis.Nil && json.Unmarshal(result, &article) == nil {
		c.JSON(http.StatusOK, article)
		return
	}

	fmt.Println(string(result))

	// if doesnt return get value if value is earlier than today send it to queue
	// save data to elastic search
	// save it to postgres as well
	c.JSON(http.StatusOK, "add")
}
func (h *ArticleHandler) Update(c *gin.Context) {
	// we want to validate data
	// we want to get cache if cache exists return it save it
	// if doesnt return get value if value is earlier than today send it to queue
	// save data to postgres
	c.JSON(http.StatusOK, "update")
}

// Remove implements Handler
func (*ArticleHandler) Remove(c *gin.Context) {
	c.JSON(http.StatusOK, "Remove")
}
