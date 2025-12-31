package walker

import (
	"golang.org/x/net/html"
)

func MapFromAttributes(attributes []html.Attribute) map[string]*html.Attribute {
	ans := map[string]*html.Attribute{}

	for _, attribute := range attributes {
		ans[attribute.Key] = &attribute
	}

	return ans
}
