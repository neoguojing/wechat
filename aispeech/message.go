package aispeech

// <xml>
//
//	<userid>用户的唯一ID，（通常为用户在微信的openid）</userid>
//	<appid>应用的appid</appid>
//	<content>
//	  <msg>机器人或用户的的内容</msg>
//	</content>
//	<event>userEnter</event>
//	<from>对话来源ID 0/1/2</from>
//	<kfstate>客服接入状态 0/1/2/3</kfstate>
//	<channel>渠道id 0</channel>
//	<assessment>用户评价0/1/2/3/4/5</assessment>
//	<createtime>unix时间戳</createtime>
//
// </xml>
type CustomerServiceRecvMessage struct {
	UserID  string `xml:"userid"`
	AppID   string `xml:"appid"`
	Content struct {
		Msg string `xml:"msg"`
	} `xml:"content"`
	Event      string `xml:"event"`
	From       int    `xml:"from"`
	KfState    int    `xml:"kfstate"`
	Channel    int    `xml:"channel"`
	Assessment int    `xml:"assessment"`
	CreateTime int64  `xml:"createtime"`
}

// <xml>
//     <appid><![CDATA[wxxxxx]]></appid>
//     <openid><![CDATA[xxxxx]]></openid>
//     <msg><![CDATA[您好，请问需要什么帮助]]></msg>
//     <channel>0</channel>
//     <kefuname><![CDATA[客服昵称]]></kefuname>
//     <kefuavatar><![CDATA[客服头像图片URL地址]]></kefuavatar>
//     <ans_node_name><![CDATA[分类或技能名称]]></ans_node_name>
// </xml>
 
// 客服消息发送结构体
type CustomerServiceSendMessage struct {
	AppID      string `xml:"appid"`
	OpenID     string `xml:"openid"`
	Msg        string `xml:"msg"`
	Channel    int    `xml:"channel"`
	KefuName   string `xml:"kefuname"`
	KefuAvatar string `xml:"kefuavatar"`
	AnsNode    string `xml:"ans_node_name"`
}

// <xml>
//     <appid><![CDATA[wxxxxx]]></appid>
//     <openid><![CDATA[xxxxx]]></openid>
//     <kefustate><![CDATA[needperson]]></kefustate>
// </xml>

// 客服状态结构体
type CustomerServiceState struct {
	AppID    string `xml:"appid"`
	OpenID   string `xml:"openid"`
	KefuState string `xml:"kefustate"`
}


// encrypted	string	是	对话内容加密后字符串
// 客服消息json结构体
type CustomerServiceMessage struct {
	Encrypted string `json:"encrypted"`
}




