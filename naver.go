package main

import (
	"net/http"
)

var (
	listAddr    = "https://recruit.navercorp.com/naver/job/list/developer?"
	detailAddr  = "https://recruit.navercorp.com/naver/job/detail/developer?"
	category    = "&entTypeCd="
	categoryMap = map[string]string{
		"전체": "",
		"신입": "001",
		"경력": "002",
		"인턴": "004",
	}
	annoID = "&annoId="
)

func scrapNaver() {
	resp, err := http.Get(listAddr + category + categoryMap["신입"])
	if err != nil {
		panic(err)
	}
	println(resp)
}

func main() {
	scrapNaver()
}
