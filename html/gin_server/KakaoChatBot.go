package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
)

const ajouLink = "https://www.ajou.ac.kr/kr/ajou/notice.do"

// Notice {id: int, title: str, date: str, link: str, writer: str}
type Notice struct {
	id     int64
	title  string
	date   string
	link   string
	writer string
}

// Parse is a function that parses a length of notices
func Parse(length int) []Notice {
	ajouHTML := fmt.Sprintf("%v?mode=list&articleLimit=%v&article.offset=0", ajouLink, length)

	resp, err := soup.Get(ajouHTML)
	if err != nil {
		fmt.Println("Check your HTML connection.")
		os.Exit(2)
	}
	doc := soup.HTMLParse(resp)

	notices := []Notice{}
	ids := doc.FindAll("td", "class", "b-num-box")
	if len(ids) == 0 {
		fmt.Println("Check your parser.")
		os.Exit(2)
	}

	titles := doc.FindAll("div", "class", "b-title-box")
	dates := doc.FindAll("span", "class", "b-date")
	//links := doc.FindAll("div", "class", "b-title-box")
	writers := doc.FindAll("span", "class", "b-writer")
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseInt(strings.TrimSpace(ids[i].Text()), 10, 64)
		title := strings.TrimSpace(titles[i].Find("a").Text())
		link := titles[i].Find("a").Attrs()["href"]
		date := strings.TrimSpace(dates[i].Text())
		writer := writers[i].Text()
		notice := Notice{id: id, title: title, date: date, link: ajouLink + link, writer: writer}
		notices = append(notices, notice)
	}

	return notices
}

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

	router.POST("/last", func(c *gin.Context) {
		var kjson KakaoJSON
		if err := c.BindJSON(&kjson); err != nil {
			errorMsg := SimpleText{Version: "2.0"}
			errorMsg.Template.Outputs.SimpleText.Text = err.Error()
			c.JSON(http.StatusBadRequest, errorMsg)
			return
		}

		notice := Parse(1)[0]
		noticeJSON := gin.H{"title": notice.title, "description": notice.writer, "link": gin.H{"web": notice.link}}

		// Card
		items := []gin.H{}
		buttons := []gin.H{}
		header := gin.H{"title": "header title"}

		// Add one care item
		items = append(items, noticeJSON)
		buttons = append(buttons, gin.H{"label": "공유하기", "action": "share"})

		// Make a template
		template := gin.H{"outputs": []gin.H{{"listCard": gin.H{"header": header, "items": items, "buttons": buttons}}}}

		listCard := gin.H{"version": "2.0", "template": template}
		c.PureJSON(http.StatusOK, listCard)
		// c.JSON(http.StatusOK, listCard)

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
		items = append(items, item, item2)

		// QuickReplies [Optional]
		quickReplies := []gin.H{}

		// Add one quick reply
		quickReply := gin.H{"messageText": "안녕하세요", "action": "message", "label": "안녕"}
		quickReplies = append(quickReplies, quickReply)

		// Make a template
		template := gin.H{"outputs": []gin.H{{"listCard": gin.H{"header": header, "items": items}}}}
		template["quickReplies"] = quickReplies // Optional
		listCard := gin.H{"version": "2.0", "template": template}

		// buf := new(bytes.Buffer)
		// enc := json.NewEncoder(buf)
		// enc.SetEscapeHTML(false)
		// if err := enc.Encode(&listCard); err != nil {
		// 	log.Println(err)
		// }

		c.PureJSON(http.StatusOK, listCard)
	})

	router.Run(":8000")
}
