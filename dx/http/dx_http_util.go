package dx_http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ReadReq(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	// 我们只处理POST请求
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("只支持POST请求"))
		//return nil, dx_error.Error("变量类型错误")
		return nil, nil
	}

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("读取请求体失败"))
		log.Printf("读取请求体失败: %v", err)
		return nil, nil
	}
	defer r.Body.Close()
	return body, err
}
