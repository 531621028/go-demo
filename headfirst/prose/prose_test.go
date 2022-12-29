package prose

import (
	"fmt"
	"testing"
)

func TestTwoElement(t *testing.T) {
	list := []string{"apple", "orange"}
	expect := "apple and orange"
	actual := JoinWithCommas(list)
	if expect != actual {
		t.Errorf(errorInfo(list, actual, expect))
	}
}

func TestThreeElement(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	expect := "apple , orange and pear"
	actual := JoinWithCommas(list)
	if expect != actual {
		t.Errorf(errorInfo(list, actual, expect))
	}
}

func errorInfo(data []string, expect string, actual string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", expected is \" %s\"", data, actual, expect)
}
