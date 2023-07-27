package dx_openai

import (
	"bytes"
	"net/http"
)

func CreateReq(model OpenAIReq, bodyByff []byte) (*http.Request, error) {
	return http.NewRequest(model.method, model.url, bytes.NewBuffer(bodyByff))
}

func CreateReqSimple(bodyByff []byte) (*http.Request, error) {
	return CreateReq(Model_Alone, bodyByff)
}
