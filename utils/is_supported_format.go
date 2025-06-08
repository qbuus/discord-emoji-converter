package utils

import (
	"path/filepath"
	"slices"
	"strings"
)

var supportedFormats = []string{".png", ".jpg", ".jpeg", ".gif"}

func IsSupportedFormat(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	return slices.Contains(supportedFormats, extension)
}
