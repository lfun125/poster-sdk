package gosdk

import (
	"fmt"
	"testing"
)

func testNewClient() *Client {
	c, err := NewClient("https://api.daishuhaibao.com", "2276907093", `-----BEGIN RSA PRIVATE KEY-----
MIIBPQIBAAJBALBQtjwSjxB8IPBsyzKU9sAHs+Z2ybOFvbv9+suMtvSc9HtyK3Cv
Jl8XbcUjKwO3AT8jjE+z5Aa3MDadZ9tuSlsCAwEAAQJBAK2xNXWEiu+geQqhAqR3
34VZkT5us1FpZXq3P8QagXoDKmQMlJG0l+Evoe+l+gjsKwBn/a12srd5k8a8W4tH
12ECIQDMSs362HdnAN3PbHDirlEr3l0DqZZXwnWlgoH1qCGxUwIhANzxIZUOJ0su
ZBzm63xKLO73jzoj9fmFxH+z8K0yHrnZAiEAm/6G2IOQcUO1G90nOcB31PKfvdsy
JOMlKHPGPsnaqGcCIQDYenbyCeUcN78lxmhS5oayOeOfDt8Sdiu7CD64JFSQuQIh
AMwIiwqeF1bG6yS8/dwtDkjblV64Z1KhHnWuLXLDq8cd
-----END RSA PRIVATE KEY-----
	`)
	if err != nil {
		panic(err)
	}
	return c
}

func TestClient_Login(t *testing.T) {
	c := testNewClient()
	v, err := c.Login("M3aLsz4tPB_0218")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_SetUser(t *testing.T) {
	c := testNewClient()
	v, err := c.SetUser("M3aLsz4tPB_0218", "迅德研发测试")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_rsaDecrypt(t *testing.T) {
	c := testNewClient()
	v, err := c.rsaDecrypt("b863044e1244f2e0ce5e6a7fa4948fe689d3827cbf4d9e4d0db9dab3db153ca40bcafa9c4ed1e9db717c3942709e46660a39602829fa53b25ad57faf380bb324")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_signWithSha256(t *testing.T) {
	c := testNewClient()
	v, err := c.signWithSha256("你好中国")
	fmt.Println(err)
	fmt.Println(v)
}
