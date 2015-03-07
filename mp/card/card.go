// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com)

package card

type Card struct {
	// 卡券类型
	CardType string `json:"card_type"`
	// 基本卡券数据
	GeneralCoupon struct {
		BaseInfo      BaseInfo `json:"base_info"`
		DefaultDetail string   `json:"default_detail"`
	} `json:"general_coupon,omitempty"`
	// 团购券
	Groupon struct {
		BaseInfo   BaseInfo `json:"base_info"`
		DealDetail string   `json:"deal_detail"`
	} `json:"groupon,omitempty"`
	// 礼品券
	Gift struct {
		BaseInfo BaseInfo `json:"base_info"`
		Gift     string   `json:"gift"`
	} `json:"gift,omitempty"`
	// 代金券
	Cash struct {
		BaseInfo   BaseInfo `json:"base_info"`
		LeastCost  int      `json:"least_cost"`
		ReduceCost int      `json:"reduce_cost"`
	} `json:"cash,omitempty"`
	// 折扣券
	Discount struct {
		BaseInfo BaseInfo `json:"base_info"`
		Discount int      `json:"discount"`
	} `json:"discount,omitempty"`
	// 会员卡
	MemberCard struct {
		BaseInfo       BaseInfo `json:"base_info"`
		SupplyBonus    bool     `json:"supply_bonus"`
		SupplyBalance  bool     `json:"supply_balance"`
		BonusCleared   string   `json:"bonus_cleared,omitempty"`
		BonusRules     string   `json:"bonus_rules"`
		BalanceRules   string   `json:"balance_rules,omitempty"`
		Prerogative    string   `json:"prerogative"`
		BindOldCardUrl string   `json:"bind_old_card_url,omitempty"`
		ActivateUrl    string   `json:"activate_url,omitempty"`
	} `json:"member_card,omitempty"`
	// 门票
	ScenicTicket struct {
		BaseInfo    BaseInfo `json:"base_info"`
		TicketClass string   `json:"ticket_class,omitempty"`
		GuideURl    string   `json:"guide_url,omitempty"`
	} `json:"scenic_ticket,omitempty"`
	// 电影票
	MovieTicket struct {
		BaseInfo BaseInfo `json:"base_info"`
		Detail   string   `json:"detail,omitempty"`
	} `json:"movie_ticket,omitempty"`
	// 飞机票
	BoardingPass struct {
		BaseInfo      BaseInfo `json:"base_info"`
		From          string   `json:"from"`
		To            string   `json:"to"`
		Flight        string   `json:"flight"`
		DepartureTime string   `json:"departure_time,omitempty"`
		LandingTime   string   `json:"landing_time,omitempty"`
		CheckInUrl    string   `json:"check_in_url,omitempty"`
		AirModel      string   `json:"air_model,omitempty"`
	} `json:"boarding_pass,omitempty"`
	// 红包
	LuckMoney struct {
		BaseInfo BaseInfo `json:"base_info"`
	} `json:"lucky_money,omitempty"`
}

// 基本的卡券数据，所有卡券通用。作为 Card_BaseInfo和 的基类
type BaseInfo struct {
	// 卡券的商户logo，尺寸为300*300。
	LogoUrl string `json:"logo_url"`
	// code码表示类型
	CodeType string `json:"code_type"`
	// 商户名字,字数上限为12 个汉字。（填写直接提供服务的商户名， 第三方商户名填写在source 字段）
	BrandName string `json:"brand_name"`
	// 券名，字数上限为9 个汉字。(建议涵盖卡券属性、服务及金额)
	Title string `json:"title"`
	// 券名的副标题，字数上限为18个汉字。
	SubTitle string `json:"sub_title,omitempty"`
	// 券颜色。按色彩规范标注填写Color010-Color100
	Color string `json:"color"`
	// 使用提醒，字数上限为9 个汉字。（一句话描述，展示在首页，示例：请出示二维码核销卡券）
	Notice string `json:"notice"`
	// 使用说明。长文本描述，可以分行，上限为1000 个汉字。
	Description string `json:"description"`
	// 有效日期
	DateInfo DateInfo `json:"date_info"`
	// 商品信息
	Sku struct {
		// 上架的数量。（不支持填写0或无限大）
		Quantity int64 `json:"quantity"`
	} `json:"sku"`
	// 门店地址ID
	LocationIdList []int64 `json:"location_id_list,omitempty"`
	// 是否自定义code码
	UseCustomCode bool `json:"use_custom_code,omitempty"`
	// 是否指定用户领取，填写true或false。不填代表默认为否。
	BindOpenid bool `json:"bind_openid,omitempty"`
	// 领取卡券原生页面是否可分享，填写true 或false，true 代表可分享。默认可分享。
	CanShare bool `json:"can_share,omitempty"`
	// 卡券是否可转赠，填写true 或false,true 代表可转赠。默认可转赠。
	CanGiveFriend bool `json:"can_give_friend,omitempty"`
	// 每人最大领取次数
	GetLimit int `json:"get_limit,omitempty"`
	//每人使用次数限制
	UseLimit int `json:"use_limit,omitempty"`
	// 客服电话
	ServicePhone string `json:"service_phone,omitempty"`
	// 第三方来源名，如携程
	Source string `json:"source,omitempty"`
	// 商户自定义cell名字
	UrlNameType string `json:"url_name_type,omitempty"`
	// 商户自定义cell跳转外链的地址
	CustomUrl string `json:"custom_url,omitempty"`
}

// 使用日期，有效期的信息
type DateInfo struct {
	// 使用时间的类型 1：固定日期区间，2：固定时长（自领取后按天算）
	Type int `json:"type"`
	// 固定日期区间专用，表示起用时间。从1970 年1 月1 日00:00:00 至起用时间的秒数，最终需转换为字符串形态传入，下同。（单位为秒）
	BeginTimestamp int64 `json:"begin_timestamp"`
	// 固定日期区间专用，表示结束时间。（单位为秒）
	EndTimestamp int64 `json:"end_timestamp"`
	// 固定时长专用，表示自领取后多少天内有效。（单位为天））
	FixedTerm int `json:"fixed_term"`
	// 固定时长专用，表示自领取后多少天开始生效。（单位为天）
	FixedBeginTerm int `json:"fixed_begin_term"`
}

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// json 返回卡券信息
type ResultCard struct {
	CardId    string `json:"card_id"` // 卡券ID
	BeginTime int64  `json:"begin_time"`
	EndTime   int64  `json:"end_time"`
}
