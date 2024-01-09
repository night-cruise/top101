package str

import "strings"

func SolveBM85(IP string) string {
	ipV4List := strings.Split(IP, ".")
	if len(ipV4List) == 4 {
		for _, ip := range ipV4List {
			if len(ip) == 0 || len(ip) > 1 && ip[0] == '0' {
				return "Neither"
			}
			ipNumber := 0
			for i := 0; i < len(ip); i++ {
				ipNumber = ipNumber*10 + int(ip[i]-'0')
			}
			if ipNumber > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	ipV6List := strings.Split(IP, ":")
	if len(ipV6List) == 8 {
		for _, ip := range ipV6List {
			if len(ip) > 4 || len(ip) == 0 {
				return "Neither"
			}
			for i := 0; i < len(ip); i++ {
				if !(ip[i] >= '0' && ip[i] <= '9' || ip[i] >= 'a' && ip[i] <= 'f' || ip[i] >= 'A' && ip[i] <= 'F') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}
