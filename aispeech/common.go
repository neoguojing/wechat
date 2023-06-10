package aispeech

// on/off on代表已接入，off代表未接入
type KefuState string

const (
	KefuStateOn  KefuState = "on"
	KefuStateOff KefuState = "off"
)

// 客服接入状态(kfstate):
// 状态 ID	描述
// 0	智能客服 -- 待接入 (asking)
// 1	人工客服 -- 已接入 (personserving)
// 2	结束人工客服(对话) -- 对话关闭 (complete)
// 3	待人工接入 -- 待转人工 (needperson)
type KfState int

const (
	KfStateAsking KfState = iota
	KfStatePersonServing
	KfStateComplete
	KfStateNeedPerson
)

// 对话来源(from):
// 来源 ID	描述
// 0	用户发出的内容
// 1	机器人回复的内容
// 2	客服在微信对话开放平台上进行回复时的内容

type From int

const (
	FromUser            From = 0
	FromRobot           From = 1
	FromCustomerService From = 2
)

// 渠道说明(channel):
// 渠道ID	渠道描述
// 0	接入方式为，扫码绑定的公众号、小程序接入智能对话的渠道
// 2	接入方式为，通过开放接口方式接入智能对话的渠道 (暂未开放)
// 3	接入方式为，企业微信群聊机器人（暂未开放）
// 4	接入方式为，企业微信（面向企业内部的员工）客服机器人（暂未开放）
// 5	接入方式为，公众号菜单H5机器人
// 6	接入方式为，小程序插件接入智能对话的渠道
// 7	接入方式为，网页H5机器人智能对话的渠道
type Channel int

const (
	ChannelScanBind Channel = iota
	ChannelOpenApi
	ChannelEnterpriseWeChatGroup
	ChannelEnterpriseWeChatCustomerService
	ChannelPublicMenuH5Robot
	ChannelMiniProgramPlugin
	ChannelWebH5Robot
)

// 客服触发事件(event):
type Event string
const (
    EventUserEnter Event = "userEnter"
    EventUserQuit Event = "userQuit"
    EventCustomerStuffEnter Event = "customerStuffEnter"
    EventCustomerStuffQuit Event = "customerStuffQuit"
    EventWaiterQualityEvaluate Event = "waiterQualityEvaluate"
    EventWaiterEnter Event = "waiterEnter"
    EventWaiterSwitch Event = "waiterSwitch"
    EventWaiterQuit Event = "waiterQuit"
)

// 事件 customerStuffEnter/customerStuffQuit 出现时 会额外补充一个字段 customerInfo
// 字段名称	描述
// customerInfo.name	客服人员的微信昵称
// customerInfo.avatar	客服人员的头像
// customerInfo.openid	客服人员的微信 openid
type CustomerInfo struct {
    Name   string `json:"name"`
    Avatar string `json:"avatar"`
    Openid string `json:"openid"`
}

// 客户评价(assessment):
// 评价 ID	描述
// 0	非客户评价消息
// 1	很不满
// 2	不满
// 3	一般
// 4	满意
// 5	很满意
type Assessment int

const (
	AssessmentNotCustomer Assessment = iota
	AssessmentVeryUnsatisfied
	AssessmentUnsatisfied
	AssessmentNormal
	AssessmentSatisfied
	AssessmentVerySatisfied
)
