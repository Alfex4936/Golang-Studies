package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

const (
	length   = 10
	ajouLink = "https://www.ajou.ac.kr/kr/ajou/notice.do"
)

// Notice {id: int, title: str, date: str, link: str, writer: str}
type Notice struct {
	id     int64
	title  string
	date   string
	link   string
	writer string
}

// Parse is a function that parses a length of notices
func Parse() []Notice {
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

func main() {
	var notices []Notice
	notices = Parse()
	// fmt.Println(notices)

	for _, notice := range notices {
		fmt.Println(notice.id)
		fmt.Println(notice.title)
		fmt.Println(notice.date)
		fmt.Println(notice.writer)
		fmt.Println(notice.link)
	}
}
