package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SimpleText for Kakao Response
type SimpleText struct {
	Template struct {
		Outputs struct {
			SimpleText struct {
				Text string `json:"text"`
			} `json:"simpleText"`
		} `json:"outputs"`
	} `json:"template"`
	Version string `json:"version"`
}

// KakaoJSON request main
type KakaoJSON struct {
	Action struct {
		ID          string `json:"id"`
		ClientExtra struct {
		} `json:"clientExtra"`
		DetailParams map[string]interface{} `json:"detailParams"`
		Name         string                 `json:"name"`
		Params       map[string]interface{} `json:"params"`
	} `json:"action"`
	Bot struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"bot"`
	Contexts []interface{} `json:"contexts"`
	Intent   struct {
		ID    string `json:"id"`
		Extra struct {
			Reason struct {
				Code    int64  `json:"code"`
				Message string `json:"message"`
			} `json:"reason"`
		} `json:"extra"`
		Name string `json:"name"`
	} `json:"intent"`
	UserRequest struct {
		Block struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"block"`
		Lang   string `json:"lang"`
		Params struct {
			IgnoreMe bool   `json:"ignoreMe,string"`
			Surface  string `json:"surface"`
		} `json:"params"`
		Timezone string `json:"timezone"`
		User     struct {
			ID         string `json:"id"`
			Properties struct {
				BotUserKey  string `json:"botUserKey"`
				BotUserKey2 string `json:"bot_user_key"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"user"`
		Utterance string `json:"utterance"`
	} `json:"userRequest"`
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
		var kjson KakaoJSON
		if err := c.BindJSON(&kjson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK,
			gin.H{"json": kjson, "user entered": kjson.UserRequest.Utterance, "params": kjson.Action.Params})

	})

	// SimpleText 만들고 보내는 과정 예제
	router.POST("/simple", func(c *gin.Context) {
		var kjson KakaoJSON
		u := SimpleText{Version: "2.0"}

		if err := c.BindJSON(&kjson); err != nil {
			u.Template.Outputs.SimpleText.Text = err.Error()
			c.JSON(http.StatusBadRequest, u)
			return
		}

		u.Template.Outputs.SimpleText.Text = fmt.Sprintf("Entered: %v", kjson.UserRequest.Utterance)

		c.JSON(http.StatusOK,
			gin.H{"message": u})
	})

	// ListCard 만들고 보내는 과정 예제
	router.POST("/card", func(c *gin.Context) {
		var kjson KakaoJSON
		if err := c.BindJSON(&kjson); err != nil {
			errorMsg := SimpleText{Version: "2.0"}
			errorMsg.Template.Outputs.SimpleText.Text = err.Error()
			c.JSON(http.StatusBadRequest, errorMsg)
			return
		}

		// Card
		items := []gin.H{}
		header := gin.H{"title": "header title"}

		// Card items
		item := gin.H{"title": "card", "description": "desc", "imageUrl": "img", "link": gin.H{"web": "webhref"}}
		item2 := gin.H{"title": "card2", "description": "desc2"}

		// Add two cards
		items = append(items, item)
		items = append(items, item2)

		// QuickReplies [Optional]
		quickReplies := []gin.H{}

		// Add one quick reply
		quickReply := gin.H{"messageText": "안녕하세요", "action": "message", "label": "안녕"}
		quickReplies = append(quickReplies, quickReply)

		// Make a template
		template := gin.H{"outputs": []gin.H{gin.H{"listCard": gin.H{"header": header, "items": items}}}}
		template["quickReplies"] = quickReplies // Optional
		listCard := gin.H{"version": "2.0", "template": template}

		c.JSON(http.StatusOK, listCard)
	})

	router.Run(":8000")
}
