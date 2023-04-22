package cmd

import (
	"fmt"
	"go-search/common"
	"go-search/pkg"
)

func Zone() []string {
	zoneconfig := pkg.Config()
	key := zoneconfig.Zonekey
	zoneurl := "https://0.zone/api/data/"
	data := fmt.Sprintf(`{
"query":"%s",
		"query_type":"site",
		"page":%s,
		"pagesize":40,
		"zone_key_id":"%s"
	}`, common.Query, common.Page, key)
	var json = []byte(data)
	post := common.Post(zoneurl, json)
	return post

}
