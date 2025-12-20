package common

import "strings"

func IsURL(s string) bool {
	parts := strings.Split(s, "://")

	return len(parts) == 2
}
