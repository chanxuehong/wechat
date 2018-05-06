package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAESEncryptMsg(t *testing.T) {
	appId := "wx45f133bf6fce646e"
	aesKey, err := base64.StdEncoding.DecodeString("AdiqDDDvUNCeE1ZW5XJmjf9fqNBJpGBs4vL4cHKmHBS=")
	if err != nil {
		t.Error(err)
		return
	}

	random := []byte("cc9632a98304f81c")
	plaintext := []byte(`<xml><ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
<FromUserName><![CDATA[okPEat9FRX96xG8JQvTxHLpzDV64]]></FromUserName>
<CreateTime>1458889120</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[tests text message]]></Content>
<MsgId>6265881059295447969</MsgId>
</xml>`)

	wantBase64Ciphertext := "gsKjKVChgrDDifoOdfrL/MWujrKxnSR1jBopv0zEdHFQXcx5I0bf4UxIRjektLHEziRxpU5zTw0s+gYF3WOFjk" +
		"In30gzQl9XNKr2A+DCHVG05I2EcbOnnAR5EYgjGLAipra5nOfCPPIRFQTZ7SdanniXX73YOiCAKJlNH21+PApcYu4rPXxJ4eTXbBLwmfnS" +
		"l7iojDX1LQIcC3FYmaapMQq/u+sJGsxshp4dLXJ6A5Ji3cSYAzXRVIxbNlHN1MWfdcZ0O5+ZtOU6dZiD8hJ4kkxh05EfOedWFjUy7ZhmXg" +
		"rpZ4WpYPsrythXHE2Bg1Ohz8uf5h5X31yuU/3FPoa8rD21pfZnAjBT1QCkn6MxtL5lR+yoQReLwElVbtB6yJFPIZ+n9Qh/yKfIasxkzgIE" +
		"o0pwPEbS17WCTyvRItJtU6tlo3rawJX+fsV63tKIfmLZjyZPDlV9/ka/nA/DT0KODw=="

	ciphertext := AESEncryptMsg(random, plaintext, appId, aesKey)
	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
	if base64Ciphertext != wantBase64Ciphertext {
		t.Errorf("tests AESEncryptMsg failed,\nhave: %s\nwant: %s\n", base64Ciphertext, wantBase64Ciphertext)
		return
	}
}

func TestAESDecryptMsg(t *testing.T) {
	aesKey, err := base64.StdEncoding.DecodeString("AdiqDDDvUNCeE1ZW5XJmjf9fqNBJpGBs4vL4cHKmHBS=")
	if err != nil {
		t.Error(err)
		return
	}

	wantRandom := []byte("cc9632a98304f81c")
	wantPlaintext := []byte(`<xml><ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
<FromUserName><![CDATA[okPEat9FRX96xG8JQvTxHLpzDV64]]></FromUserName>
<CreateTime>1458889120</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[tests text message]]></Content>
<MsgId>6265881059295447969</MsgId>
</xml>`)
	wantAppId := []byte("wx45f133bf6fce646e")

	base64Ciphertext := "gsKjKVChgrDDifoOdfrL/MWujrKxnSR1jBopv0zEdHFQXcx5I0bf4UxIRjektLHEziRxpU5zTw0s+gYF3WOFjk" +
		"In30gzQl9XNKr2A+DCHVG05I2EcbOnnAR5EYgjGLAipra5nOfCPPIRFQTZ7SdanniXX73YOiCAKJlNH21+PApcYu4rPXxJ4eTXbBLwmfnS" +
		"l7iojDX1LQIcC3FYmaapMQq/u+sJGsxshp4dLXJ6A5Ji3cSYAzXRVIxbNlHN1MWfdcZ0O5+ZtOU6dZiD8hJ4kkxh05EfOedWFjUy7ZhmXg" +
		"rpZ4WpYPsrythXHE2Bg1Ohz8uf5h5X31yuU/3FPoa8rD21pfZnAjBT1QCkn6MxtL5lR+yoQReLwElVbtB6yJFPIZ+n9Qh/yKfIasxkzgIE" +
		"o0pwPEbS17WCTyvRItJtU6tlo3rawJX+fsV63tKIfmLZjyZPDlV9/ka/nA/DT0KODw=="
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		t.Error(err)
		return
	}

	random, plaintext, appId, err := AESDecryptMsg(ciphertext, aesKey)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(random, wantRandom) {
		t.Errorf("tests AESDecryptMsg failed,\nhave random: %s\nwant random: %s\n", random, wantRandom)
		return
	}
	if !bytes.Equal(plaintext, wantPlaintext) {
		t.Errorf("tests AESDecryptMsg failed,\nhave plaintext: %s\nwant plaintext: %s\n", plaintext, wantPlaintext)
		return
	}
	if !bytes.Equal(appId, wantAppId) {
		t.Errorf("tests AESDecryptMsg failed,\nhave appid: %s\nwant appid: %s\n", appId, wantAppId)
		return
	}
}

func TestAEESDecryptData(t *testing.T) {
	aesKey, err := base64.StdEncoding.DecodeString("HogNecGqZeDxFIDGjBwWKw==")
	aesIv, err := base64.StdEncoding.DecodeString("aoZqkfGDWwqj6rFgWdafyw==")

	base64Ciphertext := "4B9B1aknFM6yQjAh9mFxH3iN4PZXUGfpBLB98CRAVZzNfUI4J1WUur70+NSH/5MXcCidrt44hi6dkByTtRPxrTIV1BOOTvaa2G5NWunTXZJ37/Oq0ezydbDal5v+X3bvVVeFR6MkhYI+hT9xVl2XnE/Bzon2gIq9F9Fy8Yny0VPqsQ95xUyXnN3/IuhiquR1pAgKjDK3kgCoqhUVNa0dRRQQgTNIpy1djbLfyErPXGTXe1qhAj7RvDdJtRloEfg63JgaB/QTR2BLEGT7/GfHnwROfngxa3esGDeBr9Mtav67R4PjYESVFtH2Yf2npaDWAvAMvr+8hNVd2tjy/3HgImctZWo7bh1OHa/ktH4wKYVbTgxZPg/lgTWC+zsl1Z9g07vWQyGpaA11HODDGn1Kdvh6esiY7T3JQ15nKs1rdyXigNQFb9+kicPiInKVfOiuMcGEazli5/UQHF5rnuWhkg/wy+PRbZybR18lLh5d+EW1CnK8gNcQblDJPvx6QFcFhPxr28WkEd7ys0BZQGCIkQ=="
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		t.Error(err)
		return
	}

	raw, err := AESDecryptData(ciphertext, aesKey, aesIv)

	fmt.Println(string(raw), err)
	if err != nil {
		t.Error(err)
		return
	}

}
