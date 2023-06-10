package aispeech

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/neoguojing/wechat/v2/aispeech/config"
	"github.com/neoguojing/wechat/v2/util"
	log "github.com/sirupsen/logrus"
)

const (
	ApiName = "aispeech"
	// 发送客服消息接口:
	//https://chatbot.weixin.qq.com/openapi/sendmsg/{TOKEN}
	sendMsgUrl         = "https://chatbot.weixin.qq.com/openapi/sendmsg/%s"
	kefustateGetUrl    = "https://chatbot.weixin.qq.com/openapi/kefustate/get/%s"
	kefustateChangeUrl = "https://chatbot.weixin.qq.com/openapi/kefustate/change/%s"
)

type CustomerService struct {
	config *config.Config
	random []byte
	msgCh  chan CustomerServiceRecvMessage
}

func NewCustomerService(cfg *config.Config) *CustomerService {
	cs := &CustomerService{
		config: cfg,
		msgCh:  make(chan CustomerServiceRecvMessage, 10),
	}
	go cs.handler()
	return cs
}

func (c *CustomerService) handler() {
	var reply string
	for {
		select {
		case msg, ok := <-c.msgCh:
			if !ok {
				break
			}
			log.Debug(msg)

			if msg.From == FromUser {
				if msg.Event == EventUserEnter && msg.Content.Msg == "" {
					reply = "请问，有什么需要帮助？"
				} else {
					if c.config.AiBot != nil {
						reply = c.config.AiBot(msg.Content.Msg)
					}
				}

			} else {
				continue
			}

			sendMsg := CustomerServiceSendMessage{
				AppID:    msg.AppID,
				OpenID:   msg.UserID,
				Msg:      reply,
				Channel:  msg.Channel,
				KefuName: "aibot",
			}
			c.SendMessage(sendMsg)
		}
	}
}

func (c *CustomerService) Close() {
	close(c.msgCh)
}

func (c *CustomerService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 解析json请求
	var req CustomerServiceMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		resp := util.NewCommonError(ApiName, 1111, "")
		resp.ErrMsg = err.Error()
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	log.Info(req)
	var err error
	var decodeMsg []byte
	c.random, decodeMsg, err = util.DecryptMsg(c.config.AppID, req.Encrypted, c.config.EncodingAESKey)
	if err != nil {
		log.Error(err)
		resp := util.NewCommonError(ApiName, 1111, "")
		resp.ErrMsg = err.Error()
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	log.Info(decodeMsg)

	// 解析xml请求
	var recvMsg CustomerServiceRecvMessage
	if err := xml.Unmarshal(decodeMsg, &recvMsg); err != nil {
		log.Error(err)
		resp := util.NewCommonError(ApiName, 1111, "")
		resp.ErrMsg = err.Error()
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}
	log.Info(recvMsg)
	c.msgCh <- recvMsg

	resp := util.NewCommonError(ApiName, 0, "")
	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (c *CustomerService) SendMessage(msg CustomerServiceSendMessage) error {
	accessToken := c.config.Token

	url := fmt.Sprintf(sendMsgUrl, accessToken)

	msgXml, err := xml.Marshal(msg)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// 加密消息
	encryptedMsg, err := util.EncryptMsg(c.random, msgXml, c.config.AppID, c.config.EncodingAESKey)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	req := CustomerServiceMessage{
		Encrypted: string(encryptedMsg),
	}
	resp, err := util.PostJSON(url, &req)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	result := util.DecodeWithCommonError(resp, ApiName)
	log.Info(result.Error())

	return nil
}
