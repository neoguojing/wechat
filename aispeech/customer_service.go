package aispeech

import (
	"github.com/neoguojing/wechat/v2/aispeech/config"
	"net/http"
)

const (
	// 发送客服消息接口:
	//https://chatbot.weixin.qq.com/openapi/sendmsg/{TOKEN}
	sendMsgUrl         = "https://chatbot.weixin.qq.com/openapi/sendmsg/%s"
	kefustateGetUrl    = "https://chatbot.weixin.qq.com/openapi/kefustate/get/%s"
	kefustateChangeUrl = "https://chatbot.weixin.qq.com/openapi/kefustate/change/%s"
)

type CustomerService struct {
	config config.Config
}

func NewCustomerService(cfg config.Config) *CustomerService {
	return &CustomerService{
		config: cfg,
	}
}

func (c *CustomerService) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 解析json请求
	var req CustomerServiceMessage

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
