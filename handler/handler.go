package handler

import (
	"strconv"
	"strings"
	"time"
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

func GetUserHandler(c *gin.Context) {
	u := new(model.User)
	email := c.Query("email")
	if email == "" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		u.ID = id
	} else {
		u.Email = email
	}

	u, err := model.GetUser(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": u,
		},
	})
}

func GetUsersHandler(c *gin.Context) {
	if c.Query("email") != "" {
		GetUserHandler(c)
		return
	}
	
	
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

func CreateUserHandler(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	nameMissing := u.Name == ""
	emailMissing := u.Email == ""
	if nameMissing || emailMissing {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing data: name or email missing",
		})
		return
	}

	u.Created = time.Now()
	if err := model.CreateUser(u); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok": true,
	})
}
