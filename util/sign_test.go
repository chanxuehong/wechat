package util

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"testing"
)

func TestSign(t *testing.T) {
	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"

	wantSignature := "f9c725922f6844701ba71e98031978e40023c09f"

	haveSignature := Sign(token, timestamp, nonce)
	if haveSignature != wantSignature {
		t.Errorf("test Sign failed,\nhave signature: %s\nwant signature: %s\n", haveSignature, wantSignature)
		return
	}
}

func TestMsgSign(t *testing.T) {
	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	msg := "u0KvD5kCUzGq9QmWTsSRcolKAH92oiMZDBJ840OqXXFhzUFsBZtdDv6Fv2W9zlrP3Rx3jQGiNXY8sV0kWgzdcefN3WznfM0TGArGow" +
		"C9bgHE+4QKGrkly0wQ6ouQe7UrKDvkfMy3t8t8njawTl2z0MvgIsaVAvllB3vDWzY/Oo1P1Q9PLdIsgbQjyrJVObZMgFK4KN175ygxBWk" +
		"b4fdiIAIXSpxN/48LYixSkgohbAcUGlDP8FD3mPTDppz1yI4fGY8jHCYy3s2CsylkE/1+/6zAtv/8160FvB/C6CwRd7Q4OypAW2JKDOek" +
		"WULMhahY+43C8GQaKWWbsmOHdiEoL6N8jCyrquvCBnvekZTrRY/CTuyqmpo/5LrDXXORwKDdK0uSVU/9N1Lpz9waEWMFqIPSnYC+MgelL" +
		"zE9hBlnJ+5qC6O0GXrm6QmrncijD2SvdOhfT6OKIDiYJUM2cHHMcg=="

	wantSignature := "247dea54e9ca03ab1e75d03430d205a5b967c44d"

	haveSignature := MsgSign(token, timestamp, nonce, msg)
	if haveSignature != wantSignature {
		t.Errorf("test MsgSign failed,\nhave signature: %s\nwant signature: %s\n", haveSignature, wantSignature)
		return
	}

	haveSignature2 := MsgSign2(token, timestamp, nonce, msg)
	if haveSignature2 != wantSignature {
		t.Errorf("test MsgSign2 failed,\nhave signature: %s\nwant signature: %s\n", haveSignature2, wantSignature)
		return
	}
}

func BenchmarkMsgSign(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	msg := "u0KvD5kCUzGq9QmWTsSRcolKAH92oiMZDBJ840OqXXFhzUFsBZtdDv6Fv2W9zlrP3Rx3jQGiNXY8sV0kWgzdcefN3WznfM0TGArGow" +
		"C9bgHE+4QKGrkly0wQ6ouQe7UrKDvkfMy3t8t8njawTl2z0MvgIsaVAvllB3vDWzY/Oo1P1Q9PLdIsgbQjyrJVObZMgFK4KN175ygxBWk" +
		"b4fdiIAIXSpxN/48LYixSkgohbAcUGlDP8FD3mPTDppz1yI4fGY8jHCYy3s2CsylkE/1+/6zAtv/8160FvB/C6CwRd7Q4OypAW2JKDOek" +
		"WULMhahY+43C8GQaKWWbsmOHdiEoL6N8jCyrquvCBnvekZTrRY/CTuyqmpo/5LrDXXORwKDdK0uSVU/9N1Lpz9waEWMFqIPSnYC+MgelL" +
		"zE9hBlnJ+5qC6O0GXrm6QmrncijD2SvdOhfT6OKIDiYJUM2cHHMcg=="

	for i := 0; i < b.N; i++ {
		MsgSign(token, timestamp, nonce, msg)
	}
}

func BenchmarkMsgSign2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	msg := "u0KvD5kCUzGq9QmWTsSRcolKAH92oiMZDBJ840OqXXFhzUFsBZtdDv6Fv2W9zlrP3Rx3jQGiNXY8sV0kWgzdcefN3WznfM0TGArGow" +
		"C9bgHE+4QKGrkly0wQ6ouQe7UrKDvkfMy3t8t8njawTl2z0MvgIsaVAvllB3vDWzY/Oo1P1Q9PLdIsgbQjyrJVObZMgFK4KN175ygxBWk" +
		"b4fdiIAIXSpxN/48LYixSkgohbAcUGlDP8FD3mPTDppz1yI4fGY8jHCYy3s2CsylkE/1+/6zAtv/8160FvB/C6CwRd7Q4OypAW2JKDOek" +
		"WULMhahY+43C8GQaKWWbsmOHdiEoL6N8jCyrquvCBnvekZTrRY/CTuyqmpo/5LrDXXORwKDdK0uSVU/9N1Lpz9waEWMFqIPSnYC+MgelL" +
		"zE9hBlnJ+5qC6O0GXrm6QmrncijD2SvdOhfT6OKIDiYJUM2cHHMcg=="

	for i := 0; i < b.N; i++ {
		MsgSign2(token, timestamp, nonce, msg)
	}
}

func BenchmarkMsgSign_2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	msg := string(make([]byte, 1024))

	for i := 0; i < b.N; i++ {
		MsgSign(token, timestamp, nonce, msg)
	}
}

func BenchmarkMsgSign2_2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	token := "token1234567890"
	timestamp := "1458864389"
	nonce := "2066297436"
	msg := string(make([]byte, 1024))

	for i := 0; i < b.N; i++ {
		MsgSign2(token, timestamp, nonce, msg)
	}
}

func MsgSign2(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(encryptedMsg))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
