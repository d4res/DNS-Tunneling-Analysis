package decode

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseData(t *testing.T) {
	hex, _ := hex.DecodeString("000000140001000266697265666f780066697265666f7800")
	c, _ := ParseCmd(hex)
	fmt.Printf("%T", c)

	switch v := c.(type) {
	case ExecPacket:
		fmt.Printf("%v", v.Cmd)
	}

}

func TestGetStr(t *testing.T) {
	hex, _ := hex.DecodeString("66697265666f780066697265666f7800")
	s, i := getStr(hex)
	s2, _ := getStr(hex[i:])
	println(s, i, s2)
}

func TestBody(t *testing.T) {
	case1 := "4f89005fd2b1b60021636f6d6d616e64202864617265732d5669727475616c426f782900"
	case2 := "b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7800"
	hex1, _ := hex.DecodeString(case1)
	hex2, _ := hex.DecodeString(case2)
	msg1, _ := ParseMsg(hex1)
	msg2, _ := ParseMsg(hex2)
	res1, _ := json.Marshal(msg1)
	res2, _ := json.Marshal(msg2)
	fmt.Println(string(res1))
	fmt.Println(string(res2))
}

func TestTrans(t *testing.T) {
	fmt.Println(TransDomain(`b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7.800.sub.dares.top`, "sub.dares.top"))
}

func TestHex(t *testing.T) {
	data := `b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7.800.sub.dares.top`

	hex_data, _ := hex.DecodeString(data)
	fmt.Println(hex.EncodeToString(hex_data))
}

func TestErrParse(t *testing.T) {
	data := "www.zhihu.com.ipv6.dsa.dnsv1.com"
	msg, err := ParseMsg(TransDomain(data, "sub.dares.top"))
	fmt.Println(msg, err)
}
