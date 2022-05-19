package model

type LoginArgs struct {
	// 商户id
	Mid string `json:"mid,omitempty"`
	// 商户用户的唯一标识
	OpenId string `json:"open_id,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty"`
}

type LoginResp struct {
	// 认证token
	Token string `json:"token,omitempty"`
	// 用户在猫盒的唯一id
	Id int `json:"id,string,omitempty"`
}

type SetUserArgs struct {
	// 商户id
	Mid string `json:"mid,omitempty"`
	// 服务商系统的用户唯一标识
	OpenId string `json:"open_id,omitempty"`
	// 昵称
	Nickname string `json:"nickname,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty"`
}

type SetUserResp struct {
	// 用户在猫盒的唯一id
	Id int `json:"id,string,omitempty"`
}
