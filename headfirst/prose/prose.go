package prose

import "strings"

func JoinWithCommas(phase []string) string {
	result := strings.Join(phase[:len(phase)-1], " , ")
	result += " and "
	result += phase[len(phase)-1]
	return result
}
