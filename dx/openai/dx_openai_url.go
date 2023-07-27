package dx_openai

var HeaderType_Content = "Content-Type"
var HeaderValue_Json = "application/json"

// 授权
var HeaderType_Authorization = "Authorization"

// 授权关键字 空格必须要
var HeaderValue_Authorization = "Bearer "

var Req_POST = "POST"

// 达芬奇,单一回复，使用的是prompt模式
var DavinciCodexCompletion = "https://api.openai.com/v1/engines/davinci-codex/completions"

// 聊天专用的，需要涉及到上下文的，使用的是messages模式，
var ChatCompletion = "https://api.openai.com/v1/chat/completions"

// chat 聊天对话返回
var CompletionUrl = "https://api.openai.com/v1/engines/text-davinci-003/completions"

var temperature = 0.2

var Model_Chat = OpenAIReq{
	method: Req_POST,
	url:    ChatCompletion,
}

var Model_Alone = OpenAIReq{
	method: Req_POST,
	url:    CompletionUrl,
}
