package tweets

import "strings"

// globalToLocalXPath converts a global XPath format to a local XPath format, the format that can be used to obtain
// elements from another element
func globalToLocalXPath(globalXPath string) string {
	return strings.ReplaceAll(globalXPath, "[", "[position()=")
}
