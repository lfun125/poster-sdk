package gosdk

import (
	"fmt"
	"testing"
)

func testNewClient() *Client {
	c, err := NewClient("https://api.dev.blingok.com/", "4255512245", `
-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJi2ZrZDJzlCi3fumtiDhORvuVHS5uBphVmyZXZUCHN6wYnqGONW
zl7eWRjMpEsZb8ept80Oj/J3K6ZCnlL1YX8CAwEAAQJAeM9jtjCt6QiR2RE0ArEZ
Ara3/tk/uK0Bx5Hv3opgHSn49vzC7rtC6OUHZsJYIDSN2KuyItBh+Fo9d8JUnWpi
IQIhAMX7WpuTROrXMTHQuxlRQhDIRwyOz0+Ln3GLklVLcXmHAiEAxXbuFhuDFRgG
gS2nMLjRfC7oegMDCZHjdbeUgQEe9kkCIBwH9ZE3bbeOVlHGSudPwPndUWnuwl1x
2FZcO5DGCL/7AiABMjg9AAuqPAwGCk3B+MykEPCtlIkXLMAQ/Xyzz1rtqQIhAJcQ
TsGqMYeVzzZqZcAk8FBjNLZupTBCYdOGgp/zYe6D
-----END RSA PRIVATE KEY-----
		`)
	// 	c, err := NewClient("https://api.dev.blingok.com/", "2997881988", `
	// -----BEGIN RSA PRIVATE KEY-----
	// MIIBOQIBAAJBAL8uDrM5k2PmcF0uoTuSeqjoLQY9+KACateV/cXJMWaAoLlP33e7
	// hiLQyUR8W+UbkSF6II68IlG46uw1mO8DhYsCAwEAAQJAIVt87yKtp/GuS2P2d/l/
	// 83bHXF51wh2J3OHr7JXFS9f6dGeAlaaRO0++uuGRKtc7GihfiMM+HY4U6UOsVtBh
	// sQIhAOTKz6x4iOM5KlUz1oION8EvvuZx1fhBMBnkdwNHq13JAiEA1eo0D1SOvaVe
	// f9OobKWbNhEhV8Xt0tNM63ZHyJL6YrMCIHk+aPNzJMeVQbPJNsHRGwbLcJTaepOG
	// qCDwi4k3b77RAiA/x3iGKZv1h1zJl/3bhvTkBe9/EBB8j2ubuMRmVQw6aQIge0Xv
	// XhrQu1yXaWrSMFJn20pzj5+yg1WUYM1b2x3Cyv0=
	// -----END RSA PRIVATE KEY-----
	// 		`)
	if err != nil {
		panic(err)
	}
	return c
}

func TestClient_Login(t *testing.T) {
	c := testNewClient()
	v, err := c.Login("testM3aLsz4tPB_0218")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_SetUser(t *testing.T) {
	c := testNewClient()
	v, err := c.SetUser("1", "宠物老板测试门店")
	fmt.Println(err)
	fmt.Println(v)
}
func TestClient_ListOrder(t *testing.T) {
	c := testNewClient()
	v, err := c.ListOrder(1)
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_rsaDecrypt(t *testing.T) {
	c := testNewClient()
	v, err := c.rsaDecrypt("6b481592752118f673b7dfaa2620525f6b6a0ea2ddba0c71731a3f674d0e364583725fe00d62f3ee2858af948185781121be31ffbe9586e081dd6d44b8e8e095")
	fmt.Println(err)
	fmt.Println(v)
}

func TestClient_signWithSha256(t *testing.T) {
	c := testNewClient()
	v, err := c.signWithSha256("你好中国")
	fmt.Println(err)
	fmt.Println(v)
}
