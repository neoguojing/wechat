package aispeech

// 错误码	描述
// 1001	token无效
// 1002	机器人审核没有通过
// 1003	签名缺少userid字段
// 1004	签名字段为空
// 1005	签名过期或无效
// 1006	签名校验失败，缺少userid字段
// 1007	appid, category,label, desc 字段不能未空
// 1008	appid, openid,msg, 字段不能未空

 
const (
    ErrCodeInvalidToken = 1001
    ErrCodeRobotNotApproved = 1002
    ErrCodeMissingUserIdInSignature = 1003
    ErrCodeEmptySignature = 1004
    ErrCodeExpiredOrInvalidSignature = 1005
    ErrCodeSignatureVerificationFailedMissingUserId = 1006
    ErrCodeEmptyAppIdCategoryLabelDesc = 1007
    ErrCodeEmptyAppIdOpenIdMsg = 1008
)

var ErrCodeMap = map[int]string{
    ErrCodeInvalidToken: "token无效",
    ErrCodeRobotNotApproved: "机器人审核没有通过",
    ErrCodeMissingUserIdInSignature: "签名缺少userid字段",
    ErrCodeEmptySignature: "签名字段为空",
    ErrCodeExpiredOrInvalidSignature: "签名过期或无效",
    ErrCodeSignatureVerificationFailedMissingUserId: "签名校验失败，缺少userid字段",
    ErrCodeEmptyAppIdCategoryLabelDesc: "appid, category,label, desc 字段不能未空",
    ErrCodeEmptyAppIdOpenIdMsg: "appid, openid,msg, 字段不能未空",
}


