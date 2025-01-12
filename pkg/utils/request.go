package utils

import (
	"net/url"
	"strings"
)

func ExtractPath(url *url.URL, position uint8) string {
	path := url.Path
	dirs := strings.FieldsFunc(path, func(char rune) bool {
		return char == '/'
	})

	if len(dirs) == 0 {
		return ""
	}

	return dirs[position]
}
