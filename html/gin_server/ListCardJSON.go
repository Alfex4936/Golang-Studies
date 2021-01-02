package main

import (
	"fmt"

	"github.com/clarketm/json"
)

type Template struct {
	Outputs []struct {
		ListCard ListCardJSON `json:"listCard"`
	} `json:"outputs"`

	QuickReplies []QuickReply `json:"quickReplies,omitempty"`
}

type QuickReply struct {
	Action      string `json:"action"`
	Label       string `json:"label"`
	MessageText string `json:"messageText"`
}

type ListCardJSON struct {
	Buttons []Button `json:"buttons"`
	Header  Header   `json:"header,omitempty"`
	Items   []Item   `json:"items"`
}

type Button struct {
	Label  string `json:"label"`
	Action string `json:"action"`
}

type Item struct {
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	Link        Link   `json:"link,omitempty"`
	Title       string `json:"title"`
}

type Link struct {
	Web string `json:"web"`
}

type Header struct {
	Title string `json:"title"`
}

// KakaoListCard is main
type KakaoListCard struct {
	Template Template `json:"template"`
	Version  string   `json:"version"` // 2.0
}

func main() {
	data := KakaoListCard{Version: "2.0"}

	// Card Buttons
	buttons := []Button{{Label: "hey1", Action: "share"}, {Label: "hey2", Action: "share"}}

	// Card Contents
	items := []Item{
		{Description: "desc1", ImageUrl: "img", Link: Link{Web: "web"}, Title: "title1"},
		{Description: "desc2", Title: "title2"},
	}
	listCard := ListCardJSON{
		Buttons: buttons,
		Header:  Header{Title: "hi"},
		Items:   items,
	}

	// Add outputs to template
	var temp []struct {
		ListCard ListCardJSON `json:"listCard"`
	}

	lc := struct {
		ListCard ListCardJSON `json:"listCard"`
	}{ListCard: listCard}

	temp = append(temp, lc)
	data.Template.Outputs = temp // To make a tag [{"listCard"}]

	// Make QuickReplies
	quickreplies := []QuickReply{{
		MessageText: "어제 보여줘",
		Action:      "message",
		Label:       "어제",
	}}

	data.Template.QuickReplies = quickreplies

	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
