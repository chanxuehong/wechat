package component

type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`         // 授权方昵称
	HeadImg         string `json:"head_img"`          // 授权方头像
	ServiceTypeInfo uint   `json:"service_type_info"` // 授权方公众号类型，0代表订阅号，1代表由历史老帐号升级后的订阅号，2代表服务号
	VerifyTypeInfo  int    `json:"verify_type_info"`  // 授权方认证类型，-1代表未认证，0代表微信认证，1代表新浪微博认证，2代表腾讯微博认证，3代表已资质认证通过但还未通过名称认证，4代表已资质认证通过、还未通过名称认证，但通过了新浪微博认证，5代表已资质认证通过、还未通过名称认证，但通过了腾讯微博认证
	UserName        string `json:"user_name"`         // 授权方公众号的原始ID
	Signature       string `json:"signature"`         // 帐号介绍
	PrincipalName   string `json:"principal_name"`    // 公众号的主体名称
	Alias           string `json:"alias,omitempty"`   // 授权方公众号所设置的微信号，可能为空
	BusinessInfo    *struct {
		OpenStore uint `json:"open_store"` // 是否开通微信门店功能
		OpenScan  uint `json:"open_scan"`  // 是否开通微信扫商品功能
		OpenPay   uint `json:"open_pay"`   // 是否开通微信支付功能
		OpenCard  uint `json:"open_card"`  // 是否开通微信卡券功能
		OpenShake uint `json:"open_shake"` // open_shake
	} `json:"business_info"` // 用以了解以下功能的开通状况（0代表未开通，1代表已开通）
	QrcodeUrl         string           `json:"qrcode_url"`      // 二维码图片的URL，开发者最好自行也进行保存
	MiniProgramInfo   *MiniProgramInfo `json:"MiniProgramInfo"` // 可根据这个字段判断是否为小程序类型授权
	AuthorizationInfo *struct {
		AuthorizationAppId string     `json:"authorization_appid"` // 授权方appid
		FuncInfo           []FuncInfo `json:"func_info"`           // 公众号授权给开发者的权限集列表，ID为1到15时分别代表： 1.消息管理权限 2.用户管理权限 3.帐号服务权限 4.网页服务权限 5.微信小店权限 6.微信多客服权限 7.群发与通知权限 8.微信卡券权限 9.微信扫一扫权限 10.微信连WIFI权限 11.素材管理权限 12.微信摇周边权限 13.微信门店权限 14.微信支付权限 15.自定义菜单权限 请注意： 1）该字段的返回不会考虑公众号是否具备该权限集的权限（因为可能部分具备），请根据公众号的帐号类型和认证情况，来判断公众号的接口权限。
	} `json:"authorization_info"` // 授权信息
}

type MiniProgramInfo struct {
	Network  *MiniProgramNetwork `json:"network"`   // 小程序已设置的各个服务器域名
	FuncInfo []FuncInfo          `json:"func_info"` // 小程序授权给开发者的权限集列表，ID为17到19时分别代表： 17.帐号管理权限 18.开发管理权限 19.客服消息管理权限 请注意： 1）该字段的返回不会考虑小程序是否具备该权限集的权限（因为可能部分具备）。
}

type MiniProgramNetwork struct {
	RequestDomain   []string `json:"RequestDomain"`
	WsRequestDomain []string `json:"WsRequestDomain"`
	UploadDomain    []string `json:"UploadDomain"`
	DownloadDomain  []string `json:"DownloadDomain"`
}
