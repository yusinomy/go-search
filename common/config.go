package common

type Config struct {
	Fofamail   string `yaml:"Fofamail"`
	Fofakey    string `yaml:"Fofakey"`
	Hunterkey  string `yaml:"Hunterkey"`
	Zonekey    string `yaml:"Zonekey"`
	Zommeyekey string `yaml:"Zommeyekey"`
}

var (
	Query     string
	Size      string
	HttpProxy string
	Timeout   = 10
	Page      string
)
