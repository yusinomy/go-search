package pkg

import (
	"flag"
	"fmt"
	"go-search/common"
)

func title() {
	title := `
	__  ____  ____  ___  ___  __    _    
\ \/ /\ \/ /\ \/ / |/ / |/ /   / \   
 \  /  \  /  \  /| ' /| ' /   / _ \  
 /  \  /  \  /  \| . \| . \  / ___ \ 
/_/\_\/_/\_\/_/\_\_|\_\_|\_\/_/   \_\

	`
	fmt.Println(title)
}

func Flag() {
	flag.StringVar(&common.Query, "q", "", "语法")
	flag.StringVar(&common.Size, "s", "100", "大小")
	flag.StringVar(&common.HttpProxy, "p", "", "代理")
	flag.StringVar(&common.Page, "a", "1", "页面多少")
	flag.Parse()
}

func Pase() {
	Flag()
	if common.Query == "" {
		flag.Usage()
	}
}
