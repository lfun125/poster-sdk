package gosdk

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/lfun125/poster-sdk/gosdk/model"
)

const contentType = "application/json"

type Client struct {
	ctx        context.Context
	baseurl    string
	mid        string
	privateKey *rsa.PrivateKey
}

type Option func(c *Client) error

func WithTimeout(t time.Duration) Option {
	return func(c *Client) error {
		var fn func()
		c.ctx, fn = context.WithTimeout(c.ctx, t)
		_ = fn
		return nil
	}
}

func NewClient(url, mid, priKey string, opts ...Option) (*Client, error) {
	//获取私钥
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	c := &Client{}
	c.baseurl = strings.TrimRight(url, "/")
	c.ctx = ctx
	c.mid = mid
	c.privateKey = priv
	for _, v := range opts {
		if err := v(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Login(openId string) (resp *model.LoginResp, err error) {
	url := fmt.Sprintf("%s/%s", c.baseurl, "merchant/login")
	httpClient := &http.Client{}
	var bts []byte
	if bts, err = c.marshal(&model.LoginArgs{
		Mid:    c.mid,
		OpenId: openId,
	}); err != nil {
		return
	}
	var response *http.Response
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewReader(bts)); err != nil {
		return
	}
	req = req.WithContext(c.ctx)
	if response, err = httpClient.Do(req); err != nil {
		return
	}
	resp = &model.LoginResp{}
	if err = c.unmarshal(response, resp); err != nil {
		return
	}
	// 解密
	if resp.Token, err = c.rsaDecrypt(resp.Token); err != nil {
		return
	}
	return
}

func (c *Client) SetUser(openId, nickname string) (resp *model.SetUserResp, err error) {
	url := fmt.Sprintf("%s/%s", c.baseurl, "merchant/set_user")
	httpClient := &http.Client{}
	var bts []byte
	if bts, err = c.marshal(&model.SetUserArgs{
		Mid:      c.mid,
		OpenId:   openId,
		Nickname: nickname,
		Sign:     "",
	}); err != nil {
		return
	}
	var response *http.Response
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewReader(bts)); err != nil {
		return
	}
	req = req.WithContext(c.ctx)
	if response, err = httpClient.Do(req); err != nil {
		return
	}
	resp = &model.SetUserResp{}
	if err = c.unmarshal(response, resp); err != nil {
		return
	}
	return
}

func (c *Client) ListOrder(page int) (resp *model.ListOrderResp, err error) {
	url := fmt.Sprintf("%s/%s", c.baseurl, "merchant/order/list")
	httpClient := &http.Client{}
	var bts []byte
	if bts, err = c.marshal(&model.ListOrderArgs{
		Mid:       c.mid,
		Timestamp: time.Now().Unix(),
		Page:      int64(page),
		Sign:      "",
	}); err != nil {
		return
	}
	var response *http.Response
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewReader(bts)); err != nil {
		return
	}
	req = req.WithContext(c.ctx)
	if response, err = httpClient.Do(req); err != nil {
		return
	}
	resp = &model.ListOrderResp{}
	if err = c.unmarshal(response, resp); err != nil {
		return
	}
	return
}

func (c *Client) unmarshal(response *http.Response, out interface{}) error {
	if response.StatusCode != 200 {
		var ret struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(response.Body).Decode(&ret); err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("%d: %s", ret.Code, ret.Message))
	}
	all, _ := io.ReadAll(response.Body)
	fmt.Println(string(all))
	if err := json.NewDecoder(response.Body).Decode(out); err != nil {
		return err
	}
	return nil
}

func (c *Client) marshal(data interface{}) ([]byte, error) {
	refVal := reflect.ValueOf(data).Elem()
	signVal := refVal.FieldByName("Sign")
	if !signVal.CanSet() {
		return nil, errors.New("no sign field")
	}
	var souce string
	for i := 0; i < refVal.NumField(); i++ {
		field := refVal.Type().Field(i).Name
		if field == "Mid" || field == "Sign" {
			continue
		}
		val := refVal.Field(i).Interface()
		souce += fmt.Sprintf("%v", val)
	}
	sign, err := c.signWithSha256(souce)
	fmt.Println("souce: ", souce)
	fmt.Println("sign: ", sign)
	if err != nil {
		return nil, err
	}
	signVal.Set(reflect.ValueOf(sign))
	return json.Marshal(data)
}

// 签名
func (c *Client) signWithSha256(source string) (string, error) {
	h := sha256.New()
	h.Write([]byte(source))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}

// 私钥解密
func (c *Client) rsaDecrypt(ciphertext string) (string, error) {
	bts, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, c.privateKey, bts)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
