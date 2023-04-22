package cmd

import (
	"encoding/base64"
	"go-search/common"
)

func Base64() string {
	base := base64.URLEncoding.EncodeToString([]byte(common.Query))
	return base
}
