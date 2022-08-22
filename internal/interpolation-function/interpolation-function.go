package interpolation_function

import "strings"

func interpolate(content string, values map[string]string) string {
	for key, val := range values {
		content = strings.ReplaceAll(content, "$"+key, val)
	}
	return content
}
