package component

type AuthorizerInfo struct {
	NickName        string `json:"nick_name"` // 授权方昵称
	HeadImg         string `json:"head_img"`  // 授权方头像
	ServiceTypeInfo *struct {
		Id uint `json:"id"`
	} `json:"service_type_info"` // 授权方公众号类型，0代表订阅号，1代表由历史老帐号升级后的订阅号，2代表服务号
	VerifyTypeInfo *struct {
		Id int `json:"id"`
	} `json:"verify_type_info"` // 授权方认证类型，-1代表未认证，0代表微信认证，1代表新浪微博认证，2代表腾讯微博认证，3代表已资质认证通过但还未通过名称认证，4代表已资质认证通过、还未通过名称认证，但通过了新浪微博认证，5代表已资质认证通过、还未通过名称认证，但通过了腾讯微博认证
	UserName          string             `json:"user_name"`          // 授权方公众号的原始ID
	Signature         string             `json:"signature"`          // 帐号介绍
	PrincipalName     string             `json:"principal_name"`     // 公众号的主体名称
	Alias             string             `json:"alias,omitempty"`    // 授权方公众号所设置的微信号，可能为空
	BusinessInfo      *BusinessInfo      `json:"business_info"`      // 用以了解以下功能的开通状况（0代表未开通，1代表已开通）
	QrcodeUrl         string             `json:"qrcode_url"`         // 二维码图片的URL，开发者最好自行也进行保存
	MiniProgramInfo   *MiniProgramInfo   `json:"MiniProgramInfo"`    // 可根据这个字段判断是否为小程序类型授权
	AuthorizationInfo *AuthorizationInfo `json:"authorization_info"` // 授权信息
}

type BusinessInfo struct {
	OpenStore uint `json:"open_store"` // 是否开通微信门店功能
	OpenScan  uint `json:"open_scan"`  // 是否开通微信扫商品功能
	OpenPay   uint `json:"open_pay"`   // 是否开通微信支付功能
	OpenCard  uint `json:"open_card"`  // 是否开通微信卡券功能
	OpenShake uint `json:"open_shake"` // 是否开通微信摇一摇功能
}

type MiniProgramInfo struct {
	Network     *MiniProgramNetwork `json:"network"`              // 小程序已设置的各个服务器域名
	Categories  []map[string]string `json:"categories,omitempty"` //
	VisitStatus uint                `json:"visit_status"`
}

type MiniProgramNetwork struct {
	RequestDomain   []string `json:"RequestDomain,omitempty"`
	WsRequestDomain []string `json:"WsRequestDomain,omitempty"`
	UploadDomain    []string `json:"UploadDomain,omitempty"`
	DownloadDomain  []string `json:"DownloadDomain,omitempty"`
	BizDomain       []string `json:"BizDomain,omitempty"`
	UDPDomain       []string `json:"UDPDomain,omitempty"`
}
