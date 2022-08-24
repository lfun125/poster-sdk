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
	// 昵称
	Nickname string `json:"nickname,omitempty"`
	// 服务商系统的用户唯一标识
	OpenId string `json:"open_id,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty"`
}

type SetUserResp struct {
	// 用户在猫盒的唯一id
	Id int `json:"id,string,omitempty"`
}

type ListOrderArgs struct {
	// 商户id
	Mid       string `json:"mid,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty,string"`
	Page      int64  `json:"page,omitempty,string"`
	Sign      string `json:"sign,omitempty"`
}

type ListOrderResp struct {
	Page  int64    `json:"page,omitempty,string"`
	Count int64    `json:"count,omitempty,string"`
	List  []*Order `json:"list,omitempty"`
}

type Order struct {
	Amount        int64  `json:"amount,omitempty,string"`
	OpenId        string `json:"open_id,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
	PayOrderId    string `json:"pay_order_id,omitempty"`
	PaymentMethod int64  `json:"payment_method"`
	PaymentTime   int64  `json:"payment_time,omitempty,string"`
}
