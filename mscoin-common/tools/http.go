package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Post 人机验证
func Post(url string, params any) ([]byte, error) {
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqBody := strings.NewReader(s)
	httpReq, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()
	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
