package cmd

import (
	"fmt"
	"go-search/common"
	"go-search/pkg"
)

func Hunter() []string {
	hunterconfig := pkg.Config()
	key := hunterconfig.Hunterkey
	hunterurl := fmt.Sprintf("https://hunter.qianxin.com/openApi/search?api-key=%s&search=%s&page=%s&page_size=%s", key, Base64(), common.Page, common.Size)
	get := common.Get(hunterurl)
	return get
}
