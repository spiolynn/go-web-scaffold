package setting

/*
配置文件读取
 */

import (
	"fmt"
	"go-web-scaffold/pkgs/file"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	DB struct{
		Type string `yaml: "type"`
		User string `yaml: "user"`
		Password string  `yaml: "password"`
		Host string  `yaml: "host"`
		Sid string  `yaml: "sid"`
		Tableprefix string  `yaml: "tableprefix"`
	}
	App struct{
		Temp string `yaml: "temp"`
	}
	Logs struct{
		Logsrootpath string `yaml: "logsrootpath"`
		Logslevel string `yaml: "logslevel"`
		Logsfilename string `yaml: "logsfilename"`
	}
	Server struct{
		Runmode string `yaml: "runmode"`
		Httpprot int `yaml: "httpprot"`
		Readtimeout int `yaml: "readtimeout"`
		Wirtetimeout int `yaml: "wirtetimeout"`
	}
}


var (
	cfg_file   string
	//G_cfg_yaml map[string]interface{}
	G_cfg_yaml Config
)


func init() {

	fmt.Println(" init setting ... ")

	cfg_file = "C:\\GOLib\\src\\go-web-scaffold\\conf\\dev.yaml"

	if !file.Exists(cfg_file) {
		log.Fatal(cfg_file + " is not exists")
	}

	buffer, err := ioutil.ReadFile(cfg_file)
	err = yaml.Unmarshal(buffer, &G_cfg_yaml)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// ffmt.Println("%v",G_cfg_yaml)
}




func main(){

}
