package wechat

import (
	"github.com/neoguojing/wechat/v2/aispeech"
	speechConfig "github.com/neoguojing/wechat/v2/aispeech/config"
	"github.com/neoguojing/wechat/v2/cache"
	"github.com/neoguojing/wechat/v2/miniprogram"
	miniConfig "github.com/neoguojing/wechat/v2/miniprogram/config"
	"github.com/neoguojing/wechat/v2/officialaccount"
	offConfig "github.com/neoguojing/wechat/v2/officialaccount/config"
	"github.com/neoguojing/wechat/v2/openplatform"
	openConfig "github.com/neoguojing/wechat/v2/openplatform/config"
	"github.com/neoguojing/wechat/v2/pay"
	payConfig "github.com/neoguojing/wechat/v2/pay/config"
	"github.com/neoguojing/wechat/v2/work"
	workConfig "github.com/neoguojing/wechat/v2/work/config"
)

func init() {

}

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

// SetCache 设置cache
func (wc *Wechat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return openplatform.NewOpenPlatform(cfg)
}

// GetWork 获取企业微信的实例
func (wc *Wechat) GetWork(cfg *workConfig.Config) *work.Work {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return work.NewWork(cfg)
}

// GetAiSpeech 获取企业微信的实例
func (wc *Wechat) GetAiSpeech(cfg *speechConfig.Config) *aispeech.CustomerService {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return aispeech.NewCustomerService(cfg)
}
