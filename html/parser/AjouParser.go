package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

const ajouLink = "https://www.ajou.ac.kr/kr/ajou/notice.do"
const ajouPeople = "https://mportal.ajou.ac.kr/system/phone/selectList.ajax"

// Notice {id: int, title: str, date: str, link: str, writer: str}
type Notice struct {
	id     int64
	title  string
	date   string
	link   string
	writer string
}

// People
type People struct {
	MsgCode     string `json:"msgCode"`
	PhoneNumber []struct {
		BussNm    string `json:"bussNm"`           // "XXX학과(공학인증)"
		DeptCd    string `json:"deptCd"`           // "DS01234657"
		DeptNm    string `json:"deptNm"`           // "정보통신대학교학팀(팔달관 777-1)"
		Email     string `json:"email"`            // "example@ajou.ac.kr"
		KorNm     string `json:"korNm"`            // "이름1(직원)" | "이름2(교원)"
		MdfLineNo int64  `json:"mdfLineNo,string"` // "289"
		TelNo     string `json:"telNo"`            // 031-219-"1234"
		UserNo    int64  `json:"userNo,string"`    // "201900000"
	} `json:"phoneNumber"`
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

func GetPeople() (People, error) {
	jsonValue, _ := json.Marshal(map[string]string{"keyword": "이재현"})

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post((ajouPeople), "application/json;charset=UTF-8", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var people People

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &people)
	if err != nil {
		panic(err)
	}

	return people, nil
}

func main() {
	var notices []Notice
	notices = Parse(5)
	// fmt.Println(notices)

	for _, notice := range notices {
		fmt.Println(notice.id)
		fmt.Println(notice.title)
		fmt.Println(notice.date)
		fmt.Println(notice.writer)
		fmt.Println(notice.link)
	}

	people, _ := GetPeople()
	fmt.Println(people.PhoneNumber[0].KorNm)
	fmt.Println(people.PhoneNumber[0].TelNo)
}
