package config

import "github.com/go-chinese-site/cfg"

var YamlConfig *cfg.YamlConfig

func Parse(configFile string){
	var err error
	YamlConfig,err=cfg.ParseYaml(configFile)
	if err!=nil{
		panic(err)
	}
}