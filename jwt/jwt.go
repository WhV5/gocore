package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Token jwtToken 结构体
type Token struct {
	Header    map[string]string
	Payload   map[string]interface{}
	Signature string
}

//NewToken 返回新的token,这里为了方便 算法统一采用 HS256 算法
//用户需要传入过期时间戳,token 验证时候验证是否过期
func NewToken(secret string, expire int64, data interface{}) (*Token, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	payload := map[string]interface{}{
		"expire": expire,
		"data":   data,
	}

	t := &Token{
		Header:  header,
		Payload: payload,
	}

	err := t.signToken(secret)

	return t, err

}

//ValidateToken  check token exists and expire
func ValidateToken(secret, token string) (interface{}, error) {

	return verify(token, secret)

}

//signToken token 验签
func (t *Token) signToken(secret string) error {

	hb, err := json.Marshal(t.Header)

	if err != nil {

		return errors.New("json marshal header errs")
	}

	pb, err := json.Marshal(t.Payload)

	if err != nil {

		return errors.New("json marshal payload err")
	}

	s := encode(hb) + "." + encode(pb)

	sign := encrypt(secret, s)

	t.Signature = sign

	return nil
}

//base64 将字符串base64 成字符串
func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//decode 将token 解码
func decode(pstr string) (map[string]interface{}, error) {
	b, err := base64.StdEncoding.DecodeString(pstr)

	if err != nil {

		return nil, errors.New("token part decode err")
	}
	m := make(map[string]interface{})

	err = json.Unmarshal(b, &m)

	return m, nil
}

//encrypt hmac encrypt and encode to string
func encrypt(secret, pstr string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(pstr))
	return encode(h.Sum(nil))

}

//verify Check if the token has been tampered,
func verify(token, secret string) (interface{}, error) {
	t := strings.Split(token, ".")

	if len(t) != 3 {

		return nil, fmt.Errorf("token split errs, required 3 parts,but find %d parts", len(t))
	}

	pstr := t[0] + "." + t[1]

	sign := encrypt(secret, pstr)

	if sign != t[2] {
		return nil, fmt.Errorf("signature errs;token has been tampered")
	}

	pb, err := decode(t[1])

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	if f := isExpire(pb["expire"].(int64), now); !f {
		return nil, errors.New("token has expire")
	}

	return pb["data"], nil

}

//isExpire 是否过期
func isExpire(old, now int64) bool {
	return old >= now
}
