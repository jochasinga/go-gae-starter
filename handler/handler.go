package handler

import (
	"strconv"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jochasinga/boo/model"
)

func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"hello": name,
	})
}

func HelloPostHandler(c *gin.Context) {
        data := make(map[string]interface{})
        if err := c.ShouldBindJSON(&data); err != nil {
                c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
                        "error": err.Error(),
                })
                return
        }

	c.JSON(http.StatusOK, gin.H{
		"hello": data["name"],
	})
}

func GetUsersHandler(c *gin.Context) {

	sorting := strings.ToLower(c.Query("sort"))
	ordering := strings.ToLower(c.Query("ord"))
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		limit = 20
	}

	if sorting == "" {
		sorting = "DESC"
	}

	if ordering == "" {
		ordering = "created"
	}
	
	users, err := model.GetUsers(sorting, ordering, limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"users": users,
		},
	})
}
