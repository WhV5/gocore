package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Token jwtToken 结构体
type token struct {
	Header    Header
	Payload   Payload
	Signature string
}

type Header struct {
	Alg string //签名算法  hmac sha256
	Typ string //令牌类型  jwt
}

type Payload struct {
	Iss  string      //签发人
	Exp  int64       //过期时间
	Sub  string      //主题
	Iat  time.Time   //签发时间
	Data interface{} //内容
}

func newToken(secret string, expire int64, data interface{}) (*token, error) {
	h := Header{
		Alg: "hs256",
		Typ: "jwt",
	}

	p := Payload{
		Iss:  "",
		Exp:  expire,
		Sub:  "",
		Iat:  time.Now(),
		Data: data,
	}

	signature, err := sign(h, p, secret)

	if err != nil {
		return nil, err
	}

	return &token{Header: h, Payload: p, Signature: signature}, nil
}

//NewTokenStr
func NewTokenStr(secret string, expire int64, data interface{}) (string, error) {

	t, err := newToken(secret, expire, data)

	if err != nil {
		return "", errors.New("generator token has error occurred")
	}

	sh, err := encode(t.Header)

	if err != nil {
		return "", err
	}

	sp, err := encode(t.Payload)

	jwtToken := sh + "." + sp + "." + t.Signature

	return jwtToken, nil
}

//VerifyToken 严重token
func VerifyToken(ts, secret string) (interface{}, error) {

	//验证 token 是否被篡改

	sp := strings.Split(ts, ".")

	if ok := len(sp) != 3; ok {
		return nil, errors.New("token has been tampered")
	}

	hs := sp[0]
	ps := sp[1]
	ss := sp[3]

	psign := hs + "." + ps

	rs, err := sha245(psign, secret)

	if err != nil {
		return nil, errors.New("token verify has error been occurred")
	}

	if rs != ss {
		return nil, errors.New("signature has been tampered")
	}

	//校验 payload 是否过期

	p, err := getPayload(ps)

	if err != nil {
		return nil, errors.New("token payload decode error has been occurred")
	}

	return p, nil

}

func getPayload(ps string) (*Payload, error) {
	var payload Payload

	err := decode(ps, &payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func sha245(ps string, secret string) (string, error) {
	hmac := sha256.New()

	_, err := hmac.Write([]byte(ps))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hmac.Sum([]byte{})), nil
}

func sign(h Header, p Payload, secret string) (string, error) {
	he, err := encode(h)

	if err != nil {
		return "", err
	}

	pe, err := encode(p)

	if err != nil {
		return "", err
	}

	psign := he + "." + pe

	hmac := sha256.New()

	_, err = hmac.Write([]byte(psign))

	if err != nil {
		return "", errors.New("sign error has occurred")
	}

	return base64.StdEncoding.EncodeToString(hmac.Sum([]byte{})), nil

}

func encode(value interface{}) (string, error) {

	b, err := json.Marshal(value)

	if err != nil {
		return "", errors.New("marshal json has error occurred")
	}
	s := base64.StdEncoding.EncodeToString(b)

	fmt.Println(s)
	return s, nil
}

func decode(s string, v interface{}) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return errors.New("decode string has error occurred")
	}

	err = json.Unmarshal(b, &v)

	if err != nil {
		return errors.New("unmarshal has error occurred")
	}

	return nil
}
