package wxc1c68623b7bdea7b

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/miniapp/wxa/serviceMarket"
	"github.com/chanxuehong/wechat/mp/core"
)

type RGeoCRequest struct {
	Location string `json:"location"`
	GetPoi   int    `json:"get_poi"`
	Options  string `json:"poi_options,omitempty"`
}

type RGeoCResponse struct {
	Status  int          `json:"status"`           // 状态码，0为正常, 310请求参数信息有误， 311Key格式错误, 306请求有护持信息请检查字符串, 110请求来源未被授权
	Message string       `json:"message"`          // 状态说明
	Result  *RGeoCResult `json:"result,omitempty"` // 逆地址解析结果
}

func (this *RGeoCResponse) IsError() bool {
	return this.Status != 0
}

func (this *RGeoCResponse) Error() string {
	return this.Message
}

type RGeoCResult struct {
	Address            string              `json:"address,omitempty"`             // 地址描述
	FormattedAddresses *FormattedAddresses `json:"formatted_addresses,omitempty"` // 位置描述
	AddressComponent   *AddressComponent   `json:"address_component,omitempty"`   // 地址部件，address不满足需求时可自行拼接
	AdInfo             *AdInfo             `json:"ad_info,omitempty"`             // 行政区划信息
	AddressReference   *AddressReference   `json:"address_reference,omitempty"`   // 坐标相对位置参考
	PoiCount           int                 `json:"poi_count,omitempty"`           // 查询的周边poi的总数
	Pois               []Poi               `json:"pois,omitempty"`                // 周边地点（POI）数组，对象中每个子项为一个地点（POI）对象
}

type FormattedAddresses struct {
	Recommend string `json:"recommend,omitempty"` // 经过腾讯地图优化过的描述方式，更具人性化特点
	Rough     string `json:"rough,omitempty"`     // 大致位置，可用于对位置的粗略描述
}

type AddressComponent struct {
	Nation       string `json:"nation,omitempty"`        // 国家
	Province     string `json:"province,omitempty"`      // 省
	City         string `json:"city,omitempty"`          // 市
	District     string `json:"district,omitempty"`      // 区，可能为空字串
	Street       string `json:"street,omitempty"`        // 街道，可能为空字串
	StreetNumber string `json:"street_number,omitempty"` // 门牌，可能为空字串
}

type AdInfo struct {
	NationCode string    `json:"nation_code,omitempty"` // 国家代码
	Adcode     string    `json:"adcode,omitempty"`      // 行政区划代码
	Name       string    `json:"name,omitempty"`        // 行政区划名称
	Location   *Location `json:"location,omitempty"`    // 行政区划中心点坐标
	Nation     string    `json:"nation,omitempty"`      // 国家
	Province   string    `json:"province,omitempty"`    // 省
	City       string    `json:"city,omitempty"`        // 市
	District   string    `json:"district,omitempty"`    // 区，可能为空字串
}

type Location struct {
	Lat float64 `json:"lat,omitempty"` // 纬度
	Lng float64 `json:"lng,omitempty"` // 经度
}

type AddressReference struct {
	FamousArea   *AreaInfo `json:"famous_area,omitempty"`   // 知名区域，如商圈或人们普遍认为有较高知名度的区域
	Town         *AreaInfo `json:"town,omitempty"`          // 乡镇街道
	LandMarkL1   *AreaInfo `json:"landmark_l1,omitempty"`   // 一级地标，可识别性较强、规模较大的地点、小区等 【注】对象结构同 famous_area
	LandMarkL2   *AreaInfo `json:"landmark_l2,omitempty"`   // 二级地标，较一级地标更为精确，规模更小 【注】：对象结构同 famous_area
	Street       *AreaInfo `json:"street,omitempty"`        // 街道 【注】：对象结构同 famous_area
	StreetNumber *AreaInfo `json:"street_number,omitempty"` // 门牌 【注】：对象结构同 famous_area
	Crossroad    *AreaInfo `json:"crossroad,omitempty"`     // 交叉路口 【注】：对象结构同 famous_area
	Water        *AreaInfo `json:"water,omitempty"`         // 水系 【注】：对象结构同 famous_area
}

type AreaInfo struct {
	Id       string    `json:"id,omitempty"`        // 地点唯一标识
	Title    string    `json:"title,omitempty"`     // 名称/标题
	Location *Location `json:"location,omitempty"`  // 坐标
	Distance float64   `json:"_distance,omitempty"` // 此参考位置到输入坐标的直线距离
	DirDesc  string    `json:"_dir_desc,omitempty"` // 此参考位置到输入坐标的方位关系，如：北、南、内
}

type Poi struct {
	Id       string    `json:"id,omitempty"`        // 地点唯一标识
	Title    string    `json:"title,omitempty"`     // 名称/标题
	Address  string    `json:"address,omitempty"`   // 地址
	Category string    `json:"category,omitempty"`  // 地点（POI）分类
	Location *Location `json:"location,omitempty"`  // 坐标
	Distance float64   `json:"_distance,omitempty"` // 此参考位置到输入坐标的直线距离
}

func RGeoC(clt *core.Client, location string, getPoi bool, options string) (ret *RGeoCResult, err error) {
	data := RGeoCRequest{
		Location: location,
		Options:  options,
	}
	if getPoi {
		data.GetPoi = 1
	}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     RGEOC_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var resp RGeoCResponse
	err = json.Unmarshal([]byte(resData), &resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, &resp
	}
	return resp.Result, nil
}
