package tools

import (
	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// AuthCodeToOpenId 授权码查询openid.
func AuthCodeToOpenId(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/tools/authcodetoopenid", req)
}

type AuthCodeToOpenIdRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数
	AuthCode string `xml:"auth_code"` // 扫码支付授权码，设备读取用户微信中的条码或者二维码信息

	// 可选参数
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
}

type AuthCodeToOpenIdResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	OpenId string `xml:"openid"` // 用户在商户appid下的唯一标识

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	SubOpenId string `xml:"sub_openid"` // 用户在子商户appid下的唯一标识
}

// AuthCodeToOpenId2 授权码查询openid.
func AuthCodeToOpenId2(clt *core.Client, req *AuthCodeToOpenIdRequest) (resp *AuthCodeToOpenIdResponse, err error) {
	m1 := make(map[string]string, 8)
	m1["auth_code"] = req.AuthCode
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	m2, err := AuthCodeToOpenId(clt, m1)
	if err != nil {
		return nil, err
	}

	resp = &AuthCodeToOpenIdResponse{
		OpenId:    m2["openid"],
		SubOpenId: m2["sub_openid"],
	}
	return resp, nil
}
