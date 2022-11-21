package main

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

type APIResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	var r APIResponse

	// Get Json
	_, _, errs := gorequest.New().
		SetDebug(true).
		Get("https://mocki.io/v1/be4bb003-c368-465b-92c5-1a211d2d36d7").
		EndStruct(&r)
	if len(errs) != 0 {
		log.Panic(errs)
	}

	// Send Json
	// _, _, errs := gorequest.New().
	// 	SetDebug(true).
	// 	Get("https://mocki.io/v1/be4bb003-c368-465b-92c5-1a211d2d36d7").
	// 	Send(map[string]string{"data": "msg"}).
	// 	EndStruct(&r)
	// if len(errs) != 0 {
	// 	log.Panic(errs)
	// }
	log.Printf("%v\n", r)
}
