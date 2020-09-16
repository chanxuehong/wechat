package nlp

import (
	"gopkg.in/dgrijalva/jwt-go.v3"

	"github.com/chanxuehong/wechat/openai/core"
)

func Sign(clt *core.Client, uid string, data map[string]interface{}) (signature string, err error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["data"] = data
	signature, err = token.SignedString([]byte(clt.EncodingAESKey))
	return
}
