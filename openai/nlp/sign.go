package nlp

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/bububa/wechat/openai/core"
)

func Sign(clt *core.Client, uid string, data map[string]interface{}) (signature string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"data": data,
	})
	signature, err = token.SignedString([]byte(clt.EncodingAESKey))
	return
}
