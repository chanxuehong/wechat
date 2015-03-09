// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

const (
	// 卡券类型
	CardTypeGeneralCoupon = "GENERAL_COUPON" // 通用券
	CardTypeGroupon       = "GROUPON"        // 团购券
	CardTypeGift          = "GIFT"           // 礼品券
	CardTypeCash          = "CASH"           // 代金券
	CardTypeDiscount      = "DISCOUNT"       // 折扣券
	CardTypeMemberCard    = "MEMBER_CARD"    // 会员卡
	CardTypeScenicTicket  = "SCENIC_TICKET"  // 景点门票
	CardTypeMovieTicket   = "MOVIE_TICKET"   // 电影票
	CardTypeBoardingPass  = "BOARDING_PASS"  // 飞机票
	CardTypeLuckyMoney    = "LUCKY_MONEY"    // 红包
	CardTypeMeetingTicket = "MEERING_TICKET" // 会议门票
)

type Card struct {
	CardType string `json:"card_type,omitempty"`

	*GeneralCoupon `json:"general_coupon,omitempty"`
	*Groupon       `json:"groupon,omitempty"`
	*Gift          `json:"gift,omitempty"`
	*Cash          `json:"cash,omitempty"`
	*Discount      `json:"discount,omitempty"`
	*MemberCard    `json:"member_card,omitempty"`
	*ScenicTicket  `json:"scenic_ticket,omitempty"`
	*MovieTicket   `json:"movie_ticket,omitempty"`
	*BoardingPass  `json:"boarding_pass,omitempty"`
	*LuckyMoney    `json:"lucky_money,omitempty"`
	*MeetingTicket `json:"meeting_ticket,omitempty"`
}

// 通用券
type GeneralCoupon struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	DefaultDetail string `json:"default_detail,omitempty"` // 描述文本
}

// 团购券
type Groupon struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	DealDetail string `json:"deal_detail,omitempty"` // 团购券专用，团购详情
}

// 礼品券
type Gift struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	Gift string `json:"gift,omitempty"` // 礼品券专用，表示礼品名字
}

// 代金券
type Cash struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	LeastCost  *int `json:"least_cost,omitempty"`  // 可选; 代金券专用，表示起用金额（单位为分）
	ReduceCost *int `json:"reduce_cost,omitempty"` // 代金券专用，表示减免金额（单位为分）
}

// 折扣券
type Discount struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	Discount *int `json:"discount,omitempty"` // 折扣券专用，表示打折额度（百分比）。填30 就是七折。
}

// 会员卡
type MemberCard struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	// 是否支持积分，填写true 或false，如填写true，积分相关字段均为必填。
	// 填写false，积分字段无需填写。储值字段处理方式相同。
	SupplyBonus       *bool  `json:"supply_bonus,omitempty"`
	SupplyBalance     *bool  `json:"supply_balance,omitempty"`    // 是否支持储值
	BonusClearedRules string `json:"bonus_cleared,omitempty"`     // 可选; 积分清零规则
	BonusRules        string `json:"bonus_rules,omitempty"`       // 可选; 积分规则
	BalanceRules      string `json:"balance_rules,omitempty"`     // 可选; 储值说明
	Prerogative       string `json:"prerogative,omitempty"`       // 特权说明
	BindOldCardURL    string `json:"bind_old_card_url,omitempty"` // 可选; 绑定旧卡的url，与“activate_url”字段二选一必填。
	ActivateURL       string `json:"activate_url,omitempty"`      // 可选; 激活会员卡的url，与“bind_old_card_url”字段二选一必填。
}

// 景点门票
type ScenicTicket struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	TicketClass string `json:"ticket_class,omitempty"` // 可选; 票类型，例如平日全票，套票等
	GuideURL    string `json:"guide_url,omitempty"`    // 可选; 导览图url
}

// 电影票
type MovieTicket struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	Detail string `json:"detail,omitempty"` // 可选; 电影票详情
}

// 飞机票
type BoardingPass struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	From          string `json:"from,omitempty"`           // 起点，上限为18 个汉字
	To            string `json:"to,omitempty"`             // 终点，上限为18 个汉字
	Flight        string `json:"flight,omitempty"`         // 航班
	DepartureTime int64  `json:"departure_time,omitempty"` // 可选; 起飞时间。Unix 时间戳格式
	LandingTime   int64  `json:"landing_time,omitempty"`   // 可选; 降落时间。Unix 时间戳格式
	CheckinURL    string `json:"check_in_url,omitempty"`   // 可选; 在线值机的链接
	Gate          string `json:"gate,omitempty"`           // 可选; 登机口。如发生登机口变更，建议商家实时调用该接口变更
	BoardingTime  int64  `json:"boarding_time,omitempty"`  // 可选; 登机时间，只显示“时分”不显示日期，按时间戳格式填写。如发生登机时间变更，建议商家实时调用该接口变更。
	AirModel      string `json:"air_model,omitempty"`      // 可选; 机型，上限为8 个汉字
}

// 红包
type LuckyMoney struct {
	*CardBaseInfo `json:"base_info,omitempty"`
}

// 会议门票
type MeetingTicket struct {
	*CardBaseInfo `json:"base_info,omitempty"`

	MeetingDetail string `json:"meeting_detail"`    // 会议详情
	MapURL        string `json:"map_url,omitempty"` // 可选; 会议导览图
}

// base_info ===================================================================

const (
	// 卡券code码展示类型
	CodeTypeText        = "CODE_TYPE_TEXT"         // 文本
	CodeTypeBarCode     = "CODE_TYPE_BARCODE"      // 一维码
	CodeTypeQRCode      = "CODE_TYPE_QRCODE"       // 二维码
	CodeTypeOnlyBarCode = "CODE_TYPE_ONLY_BARCODE" // 一维码无code 显示
	CodeTypeOnlyQRCode  = "CODE_TYPE_ONLY_QRCODE"  // 二维码无code 显示
)

const (
	// 自定义cell字段type类型
	URLNameTypeTakeAway       = "URL_NAME_TYPE_TAKE_AWAY"       // 外卖
	URLNameTypeReservation    = "URL_NAME_TYPE_RESERVATION"     // 在线预订
	URLNameTypeUseImmediately = "URL_NAME_TYPE_USE_IMMEDIATELY" // 立即使用
	URLNameTypeAppointment    = "URL_NAME_TYPE_APPOINTMENT"     // 在线预约
	URLNameTypeExchange       = "URL_NAME_TYPE_EXCHANGE"        // 在线兑换
	URLNameTypeVIPService     = "URL_NAME_TYPE_VIP_SERVICE"     // 会员服务(仅会员卡类型可用)
)

// 基本的卡券数据，所有卡券通用
type CardBaseInfo struct {
	CardId string `json:"id,omitempty"`     // 查询的时候有返回
	Status string `json:"status,omitempty"` // 查询的时候有返回

	LogoURL     string `json:"logo_url,omitempty"`    // 卡券的商户logo，尺寸为300*300。
	CodeType    string `json:"code_type,omitempty"`   // code 码展示类型
	BrandName   string `json:"brand_name,omitempty"`  // 商户名字,字数上限为12 个汉字。（填写直接提供服务的商户名， 第三方商户名填写在source 字段）
	Title       string `json:"title,omitempty"`       // 券名，字数上限为9 个汉字。(建议涵盖卡券属性、服务及金额)
	SubTitle    string `json:"sub_title,omitempty"`   // 可选; 券名的副标题，字数上限为18个汉字。
	Color       string `json:"color,omitempty"`       // 券颜色。按色彩规范标注填写Color010-Color100
	Notice      string `json:"notice,omitempty"`      // 使用提醒，字数上限为9 个汉字。（一句话描述，展示在首页，示例：请出示二维码核销卡券）
	Description string `json:"description,omitempty"` // 使用说明。长文本描述，可以分行，上限为1000 个汉字。

	DateInfo struct {
		// 使用时间的类型1：固定日期区间，2：固定时长（自领取后按天算）
		Type int `json:"type,omitempty"`
		// 固定日期区间专用，表示起用时间。从1970 年1 月1 日00:00:00至起用时间的秒数，最终需转换为字符串形态传入，下同。（单位为秒）
		BeginTimestamp int64 `json:"begin_timestamp,omitempty"`
		// 固定日期区间专用，表示结束时间。（单位为秒）
		EndTimestamp int64 `json:"end_timestamp,omitempty"`
		// 固定时长专用，表示自领取后多少天内有效。（单位为天）领取后当天有效填写0。
		FixedTerm *int `json:"fixed_term,omitempty"`
		// 固定时长专用，表示自领取后多少天开始生效。（单位为天）
		FixedBeginTerm *int `json:"fixed_begin_term,omitempty"`
	} `json:"date_info"` // 有效日期

	SKU struct {
		Quantity int `json:"quantity,omitempty"` // 上架的数量。（不支持填写0或无限大）
	} `json:"sku"` // 商品信息

	LocationIdList []int64 `json:"location_id_list,omitempty"` // 可选; 门店地址ID
	UseCustomCode  *bool   `json:"use_custom_code,omitempty"`  // 可选; 是否自定义code 码。
	BindOpenId     *bool   `json:"bind_openid,omitempty"`      // 可选; 是否指定用户领取，填写true或false。不填代表默认为否。
	CanShare       *bool   `json:"can_share,omitempty"`        // 可选; 领取卡券原生页面是否可分享，填写true 或false，true 代表可分享。默认可分享。
	CanGiveFriend  *bool   `json:"can_give_friend,omitempty"`  // 可选; 卡券是否可转赠，填写true 或false,true 代表可转赠。默认可转赠。
	UseLimit       *int    `json:"use_limit,omitempty"`        // 可选; 每人使用次数限制。
	GetLimit       *int    `json:"get_limit,omitempty"`        // 可选; 每人最大领取次数，不填写默认等于quantity。
	ServicePhone   string  `json:"service_phone,omitempty"`    // 可选; 客服电话
	Source         string  `json:"source,omitempty"`           // 可选; 第三方来源名，如携程
	URLNameType    string  `json:"url_name_type,omitempty"`    // 可选; 商户自定义cell 名称， 与custom_url 字段共同使用
	CustomURL      string  `json:"custom_url,omitempty"`       // 可选; 商户自定义cell 跳转外链的地址链接,跳转页面内容需与自定义cell 名称保持一致。
}
