package decode

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var data string = "4f89005fd2b1b60021636f6d6d616e64202864617265732d5669727475616c426f782900"

var hexData, _ = hex.DecodeString(data)

func TestHex(t *testing.T) {
	fmt.Println(hex.DecodeString(data))
}

func TestDecodePlain(t *testing.T) {
	// SYN request
	domain := "sub.dares.top"
	DecodePlain(data, domain)
}

func TestParseData(t *testing.T) {
	hex, _ := hex.DecodeString("000000140001000266697265666f780066697265666f7800")
	c, _ := ParseCmd(hex)
	fmt.Printf("%T", c)

	switch v := c.(type) {
	case execPacket:
		fmt.Printf("%v", v.cmd)
	}

}

func TestGetStr(t *testing.T) {
	hex, _ := hex.DecodeString("66697265666f780066697265666f7800")
	s, i := getStr(hex)
	s2, _ := getStr(hex[i:])
	println(s, i, s2)
}

func Test(t *testing.T) {
	hex, _ := hex.DecodeString("b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7800")
	msg := ParseMsg(hex)
	fmt.Println(msg)
}

func TestBody(t *testing.T) {
	case1 := "4f89005fd2b1b60021636f6d6d616e64202864617265732d5669727475616c426f782900"
	case2 := "b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7800"
	hex1, _ := hex.DecodeString(case1)
	hex2, _ := hex.DecodeString(case2)
	fmt.Println(hex1)
	fmt.Println(ParseMsg(hex1))
	fmt.Println(ParseMsg(hex2))
}

func TestTrans(t *testing.T) {
	fmt.Println(TransDomain(`b753015fd2f25cb1b6000000140001000266697265666f780066697265666f7.800.sub.dares.top`, "sub.dares.top"))
}
