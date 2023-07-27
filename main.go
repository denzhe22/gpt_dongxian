package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	dx_openai "dongxian.com/http_link/dx/openai"
	dx_user "dongxian.com/http_link/dx/user"
)

// "dongxian.com/http_link/dx_openai"
var Config *Configuration
var port UrlType = UrlType_Web //UrlType_Local
func main() {
	transMsg := "/" + EnumToString(int(Msg_Translate))
	transSetMsg := "/" + EnumToString(int(Msg_SetLanguage))
	portStr := GetUrlPort()
	http.HandleFunc(transMsg, OnTranslateHandler)
	http.HandleFunc(transSetMsg, OnSetLanguageHandler)
	http.HandleFunc("/"+EnumToString(int(Msg_CheckNet)), OnCheckNetHandler)
	http.HandleFunc("/"+EnumToString(int(Msg_GetLanguage)), OnGetLanguageHandler)
	log.Fatal(http.ListenAndServe(portStr, nil))
	//logger.Info(fmt.Sprintf("提示: %s，您今日请求次数已达上限，请明天再来，交互发问资源有限，请务必斟酌您的问题，给您带来不便，敬请谅解!", rmsg.SenderNick))
	log.Printf("trans:" + transMsg)
	log.Printf("transSet:" + transSetMsg)
	log.Printf("Port:" + portStr)
	Config = LoadConfig()
}

func EnumToString(v int) string {
	return strconv.Itoa(v)
}

// 获取端口
func GetUrlPort() string {
	if port == UrlType_Local {
		return ":" + localPort
	}
	return ":" + webPort
}

// 设置语言
func OnSetLanguageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("设置语言")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("只支持POST请求"))
		return
	}

	// 读取请求体，从数据流中把body读取出来，准备解析
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("读取请求体失败"))
		log.Printf("读取请求体失败: %v", err)
		return
	}
	defer r.Body.Close()

	var msg dx_openai.ReqSetTrasnlateLanguagesMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("解析JSON失败"))
		log.Printf("解析JSON失败: %v", err)
		return
	}

	if msg.User == "" {
		w.Write([]byte("请先设置用户名"))
		return
	}

	// 请求语言列表
	if len(msg.Languages) == 0 {
		languages := dx_user.GetLanguage(msg.User)
		languagesStr := ""
		for i := 0; i < len(languages); i++ {
			languagesStr += languages[i]
			if i != len(languages)-1 {
				languagesStr += ","
			}
		}
		w.Write([]byte(languagesStr))
		return
	}
	dx_openai.SetLanguages(msg.User, msg.Languages)

	//返回
	w.Write([]byte("0"))
}

// 获取语言
func OnGetLanguageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("设置语言")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("1")) //只支持POST请求
		return
	}

	// 读取请求体，从数据流中把body读取出来，准备解析
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("2")) //读取请求体失败
		log.Printf("读取请求体失败: %v", err)
		return
	}
	defer r.Body.Close()

	userName := string(body)

	if userName == "" {
		w.Write([]byte("3")) //请先设置用户名
		return
	}

	languages := dx_user.GetLanguage(userName)
	languagesStr := ""
	if languages != nil {
		for i := 0; i < len(languages); i++ {
			languagesStr += languages[i]
			if i != len(languages)-1 {
				languagesStr += ","
			}
		}
	}
	w.Write([]byte(languagesStr))
}

func OnCheckNetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("0"))
}

func OnTranslateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("翻译")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("只支持POST请求"))
		return
	}

	// 读取请求体，从数据流中把body读取出来，准备解析
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("读取请求体失败"))
		log.Printf("读取请求体失败: %v", err)
		return
	}
	defer r.Body.Close()

	// 获取当前时间
	//currentTime := time.Now().Format(time.RFC3339)

	// 把客户端的数据解析为JSON
	var msg dx_openai.ReqTrasnlateMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("解析JSON失败"))
		log.Printf("解析JSON失败: %v", err)
		return
	}

	if msg.User == "" {
		w.Write([]byte("请先设置用户名"))
		return
	}

	if !dx_user.HasSeted(msg.User) {
		w.Write([]byte("请先设置需要翻译的语言列表，例如: en,zh,fr"))
		return
	}

	// 创建我们的消息
	//message := string(body) + " @ " + currentTime
	reldy := dx_openai.SendToOpenAI(msg, dx_openai.TRANSLATION)

	// 写入响应
	_, err = w.Write([]byte(reldy))
	if err != nil {
		log.Printf("写入响应失败: %v", err)
	}
}
