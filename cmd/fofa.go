package cmd

import (
	"fmt"
	"go-search/common"
	"go-search/pkg"
)

func Fofa() []string {
	fofaconfig := pkg.Config()
	email := fofaconfig.Fofamail
	key := fofaconfig.Fofakey
	//fofaurl := `https://fofa.info/api/v1/info/my?email=om2bg0dsbiafivpxdq6l1dzj_87k@open_wechat&key=d2ae426b0f635b89d8d817b4424b77d3`
	fofaurl := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%s", email, key, Base64(), common.Size)
	get := common.Get(fofaurl)
	return get
}
