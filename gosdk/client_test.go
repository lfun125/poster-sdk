package gosdk

import (
	"fmt"
	"testing"
)

func testNewClient() *Client {
	c, err := NewClient("https://api.daishuhaibao.com", "1713995332", `-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBANMjb7zmwG9nTqm3gnJj8YBJPtBPLihfXeQ6wWk+5LeXkCoOhoOu
m6Vl/TeSQxbYIR3KXxVR7HRV8l6dFP5KVDcCAwEAAQJBAK6GMCqZIp/zaNZo1e7w
JgPAh6dwi9rbWaT53D2+uxf1MkjkvYItc4YUnSYq36AbfjaWoQWtBoSkxYukqppu
zQECIQD+moTfM5sodc4mwpJJm+h83QrKwuMTebHS5+KX/oNwIQIhANRL460QeR6s
hVEQ6OLDkIy/OQM0QSignFMaqV7kLBlXAiEA2V/v5elXpZm/ItSotXQJcOIAXtE5
st0J9/nHOuIwMkECIEnxwPSn1zgq6SiqViOd8HxFoqsOCWAISUrc73+AxCdzAiEA
vqwGNUdJ26kd4Uho9e4sct2iDHdwFbuLs5TN3mMJdb4=
-----END RSA PRIVATE KEY-----
	`)
	if err != nil {
		panic(err)
	}
	return c
}

func TestClient_Login(t *testing.T) {
	c := testNewClient()
	v, err := c.Login("aaa123")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_SetUser(t *testing.T) {
	c := testNewClient()
	v, err := c.SetUser("aaa123", "dly")
	fmt.Println(err)
	fmt.Println(v)
}
