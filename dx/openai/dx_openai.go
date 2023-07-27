package dx_openai

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	dx_user "dongxian.com/http_link/dx/user"
)

// "dongxian.com/http_link/dx_error"

var api_key = "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"

// 每个人最大请求次数（单日）
var maxRequest = 200

// 超时时间
var timeoutNum = 600

// 初始化openai，
func Init(_api_key string, _maxRequest int, _timeoutNum int) {
	api_key = _api_key
	maxRequest = _maxRequest
	timeoutNum = _timeoutNum
}

func SendToOpenAI(msg ReqTrasnlateMessage, doType EOderType) string {
	relay := ""
	targetLanguages := dx_user.GetLanguage(msg.User)
	switch doType {
	case TRANSLATION:
		relay = TranslateText(msg.Message, targetLanguages)
	}
	return relay
}

// 请求一个内容，返回结构
func Req(prompt string) string {
	if api_key == "" {
		return ""
	}
	return send(prompt)
}

func send(message string) string {
	data := Gpt3Request{
		Prompt:    message,
		MaxTokens: 60,
	}

	// Temperature: 0.2,
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)

	req, _ := http.NewRequest("POST", Model_Alone.url, payloadBuf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", HeaderValue_Authorization+api_key) //"Bearer "+apiKey

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var gpt3Response Gpt3Response
	err = json.NewDecoder(resp.Body).Decode(&gpt3Response)

	if err != nil {
		return "解析GPT-3响应失败:" + err.Error()
	}

	if gpt3Response.Choices == nil || len(gpt3Response.Choices) == 0 {
		return "null"
	}
	txt := gpt3Response.Choices[0].Text
	gpt3Response.Choices[0].Text = strings.TrimLeft(txt, "\n\n")

	// 写入响应
	return gpt3Response.Choices[0].Text
}
