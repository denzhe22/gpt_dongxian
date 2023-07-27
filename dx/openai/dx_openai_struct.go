package dx_openai

// Http请求创建参数
type OpenAIReq struct {
	method string
	url    string
}

// Header的设置kv
type OpenAIHeaderKV struct {
	K string
	V string
}

// 请求
type Gpt3Request struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
	// Temperature float32
}

// 返回
type Gpt3Response struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	//Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

// 客户端给过来的json结构
type ReqTrasnlateMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

// 服务器生成好翻译内容后的结果，第一个始终为原文
type ResTranslateMessage struct {
	Languages []string `json:"languages"`
}

// 请求设置翻译语言内容
type ReqSetTrasnlateLanguagesMessage struct {
	User      string   `json:"user"`
	Languages []string `json:"languages"`
}

// 返回成功与否
type ResSetTrasnlateLanguagesMessage struct {
	CodeId int `json:"code_id"`
}
