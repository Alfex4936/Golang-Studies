package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Base for nested JSON
type Base map[string]interface{}

// KakaoAction from KakaoJSON
type KakaoAction struct {
	ClientExtra  Base
	DetailParams Base
	ID           string
	Name         string
	Params       Base
}

// KakaoBot from KakaoJSON
type KakaoBot struct {
	ID   string
	Name string
}

// KakaoUserRequest from KakaoJSON
type KakaoUserRequest struct {
	Block     Base
	Lang      string
	Params    Base
	TimeZone  string
	User      KakaoUser
	Utterance string
}

// KakaoUser from KakaoUserRequest.User
type KakaoUser struct {
	ID         string
	Properties Base
	Type       string
}

// KakaoIntent from KakaoJSON
type KakaoIntent struct {
	Extra KakaoReason
	ID    string
	Name  string
}

// KakaoReason from KakaoIntent.Extra
type KakaoReason struct {
	Reason Base
}

// KakaoJSON is main JSON body from Kakao
type KakaoJSON struct {
	Action      KakaoAction      `json:"action" binding:"required"`
	Bot         KakaoBot         `json:"bot" binding:"required"`
	Contexts    []interface{}    `json:"contexts" binding:"required"`
	Intent      KakaoIntent      `json:"intent" binding:"required"`
	UserRequest KakaoUserRequest `json:"userRequest" binding:"required"`
}

// JSONMiddleware is to set all types of requests are JSON.
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(JSONMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running well.",
		})
	})

	router.POST("/json", func(c *gin.Context) {
		var json KakaoJSON
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"message": fmt.Sprintf("Reason:%v | Params['sys_text']:%v | Utterance:%v | UserID:%v",
				json.Intent.Extra.Reason["code"],
				json.Action.Params["sys_text"],
				json.UserRequest.Utterance,
				json.UserRequest.User.ID)})

	})

	router.Run(":8000")
}
