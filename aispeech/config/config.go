package config

import (
	"github.com/neoguojing/wechat/v2/cache"
)

// Config .config for 智能对话
type Config struct {
	AppID          string `json:"app_id"`           // appid
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // EncodingAESKey
	Cache          cache.Cache
}
