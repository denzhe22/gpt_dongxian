package main

// url类型，本地还是外网
type UrlType int

// 本地默认端口
var localPort = "8080"

//外网端口
var webPort = "8093"

// 服务器地址类型，本地、外网
const (
	UrlType_Local UrlType = iota
	UrlType_Web
)

// 协议id
type MegId int

const (
	// 翻译
	Msg_Translate MegId = iota
	// 设置语言
	Msg_SetLanguage
	Msg_GetLanguage
	Msg_CheckNet
)
