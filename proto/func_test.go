package proto

import (
	"fmt"
	"testing"
	"time"
)

func TestCon(t *testing.T) {
	caseFalse := "www.baidu.com"
	if ans := IsEval(caseFalse); ans != false {
		fmt.Println(ans)
		t.Errorf("%s expected false, but true", caseFalse)
	}

	caseTrue := "r51646.tunnel.tuns.org"
	if ans := IsEval(caseTrue); ans != true {
		fmt.Println(ans)
		t.Errorf("%s expected true, but false", caseTrue)
	}
}

func TestTmp(t *testing.T) {
	fmt.Println(time.Now())
}
