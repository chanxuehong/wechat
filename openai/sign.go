package openai

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"

	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
	"github.com/chanxuehong/wechat/openai/model"
)

// Sign 获取signature
func Sign(clt *core.Client, req *model.User) (signature string, expiresIn int64, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/sign/"
	var result struct {
		mpCore.Error
		Signature string `json:"signature"`
		ExpiresIn int64  `json:"expiresIn"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	signature = result.Signature
	expiresIn = result.ExpiresIn
	return
}

func LocalSign(clt *core.Client, req *model.User) (signature string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   req.ID,
		"username": req.Name,
		"avatar":   req.Avatar,
	})
	signature, err = token.SignedString([]byte(clt.EncodingAESKey))
	return
}

func ParseSign(clt *core.Client, signature string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(signature, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			return nil, errors.New("invalid signing algorithm")
		}
		return []byte(clt.EncodingAESKey), nil
	})
	return
}
