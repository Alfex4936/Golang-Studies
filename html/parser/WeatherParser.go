package main

import (
	"fmt"
	"strconv"

	"github.com/anaskhan96/soup"
)

const naver = "https://weather.naver.com/today/02117530?cpName=ACCUWEATHER" // 아큐웨더 제공 날씨

// Weather ...
type Weather struct {
	MaxTemp       int64  // 최고 온도
	MinTemp       int64  // 최저 온도
	CurrentTemp   int64  // 현재 온도
	CurrentStatus string // 흐림, 맑음 ...
	RainDay       int64  // 강수량 낮
	RainNight     int64  // 강수량 밤
	FineDust      string // 미세먼지 [보통, 나쁨]
	UltraDust     string // 초미세먼지 [보통, 나쁨]
	UV            string // 자외선 지수 [낮음, ]
}

// GetWeather is a function that parses suwon's weather today
func GetWeather() Weather {
	var weather Weather

	resp, err := soup.Get(naver)
	if err != nil {
		fmt.Println("Check your HTML connection.")
		return weather // nil
	}
	doc := soup.HTMLParse(resp)

	currentTempInt, _ := strconv.ParseInt(doc.Find("strong", "class", "current").Text(), 10, 64)

	currentStatus := doc.Find("span", "class", "weather")

	// ! 해외 기상은 일출부터 일몰 전이 낮, 일몰부터 일출 전이 밤
	temps := doc.FindAll("span", "class", "data")
	DayTemp, _ := strconv.ParseInt(temps[0].Text(), 10, 64)
	NightTemp, _ := strconv.ParseInt(temps[1].Text(), 10, 64)

	rains := doc.FindAll("strong", "class", "rainfall")
	DayRain, _ := strconv.ParseInt(rains[0].Text(), 10, 64)
	NightRain, _ := strconv.ParseInt(rains[1].Text(), 10, 64)

	// [미세먼지, 초미세먼지, 자외선, 일몰 시간]
	statuses := doc.FindAll("em", "class", "level_text")

	// struct 값 변경
	weather.CurrentTemp = currentTempInt
	weather.CurrentStatus = currentStatus.Text()

	weather.MaxTemp = DayTemp // Assume that day temp must be greater than night temp in general
	weather.MinTemp = NightTemp

	weather.RainDay = DayRain
	weather.RainNight = NightRain

	weather.FineDust = statuses[0].Text()
	weather.UltraDust = statuses[1].Text()
	weather.UV = statuses[2].Text()
	return weather
}

func main() {
	var weather Weather
	weather = GetWeather()
	fmt.Println(weather.UV)
}
