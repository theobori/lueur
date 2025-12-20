package common

type PortFromProtocolFunc func(string) int

func CreatePortFromProtocolFunc(defaultPort int) PortFromProtocolFunc {
	return func(protocol string) int {
		switch protocol {
		case "http":
			return 80
		case "https":
			return 443
		case "telnet", "tn3270":
			return 23
		default:
			return defaultPort
		}
	}
}
