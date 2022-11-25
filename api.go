package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type IPResponse struct {
	Code    int    `json:"code"`
	Address string `json:"data"`
}

func GetIP(host string) string {
	res, err := http.Get(host + "/api/v1/ip")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var ret IPResponse
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	if ret.Code != 200 {
		return ""
	}
	return ret.Address
}

type LoginRequest struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin int    `json:"ifautologin,string"`
	ChannelId string `json:"channel"`
	Action    string `json:"pagesign"`
	IPAddress string `json:"usripadd"`
}

type ChannelInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

const (
	ChannelIdList       = "_GET"
	ChannelIdOnlineInfo = "_ONELINEINFO"
	ChannelIdLogout     = "0"
)

const (
	ActionFirstAuth  = "firstauth"
	ActionSecondAuth = "secondauth"
	ActionThirdAuth  = "thirdauth"
)

type LoginResponse struct {
	Code int `json:"code"`
	Data struct {
		Channels       []ChannelInfo `json:"channels"`
		Reauth         bool          `json:"reauth"`
		UserName       string        `json:"username"`
		Balance        float32       `json:"balance,string"`
		OnlineDuration int           `json:"duration,string"`
		CurrentPort    string        `json:"outport"`
		TotalDuration  int           `json:"totaltimespan,string"`
		IPAddress      string        `json:"usripadd"`
		PromptText     string        `json:"text"`
		PromptURL      string        `json:"url"`
	} `json:"data"`
}

func Login(host string, req LoginRequest) LoginResponse {
	postBody, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := http.Post(host+"/api/v1/login", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
	}
	trans := transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(trans)
	if err != nil {
		log.Fatalln(err)
	}
	var ret LoginResponse
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	return ret
}
