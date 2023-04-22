package pkg

import (
	"fmt"
	"go-search/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Config() common.Config {
	var con common.Config
	config, err := ioutil.ReadFile("Xonfig.yaml")
	if err != nil {
		file, _ := os.OpenFile("Xonfig.yaml", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		defer func() { _ = file.Close() }()
		file.WriteString(`Fofamail: ***************
Fofakey: ***************
Hunterkey: ***************
Zonekey: ***************`)
	}
	err = yaml.Unmarshal(config, &con)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return con
}
