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
	"runtime"
)

type Config struct {
	DB struct {
		Type        string `yaml: "type"`
		User        string `yaml: "user"`
		Password    string `yaml: "password"`
		Host        string `yaml: "host"`
		Sid         string `yaml: "sid"`
		Tableprefix string `yaml: "tableprefix"`

		Maxidleconn int    `yaml: "maxidleconn"`
		Maxopenconn int    `yaml: "maxopenconn"`
		DBdebug     string `yaml: "dbdebug"`
	}
	App struct {
		Runtimebasepath    string   `yaml: "runtimebasepath"`
		Imagesavepath      string   `yaml: "imagesavepath"`
		Imagemaxsize       int      `yaml: "imagemaxsize"`
		Imageallowexts     []string `yaml: "imageallowexts"`
		Imageallowextstype []string `yaml: "imageallowextstype"`

		Filepreurl string `yaml: "filepreurl"`
		Imageurl   string `yaml: "imageurl"`
	}
	Logs struct {
		Logsrootpath string `yaml: "logsrootpath"`
		Logslevel    string `yaml: "logslevel"`
		Logsfilename string `yaml: "logsfilename"`
	}
	Server struct {
		Runmode      string `yaml: "runmode"`
		Httpprot     int    `yaml: "httpprot"`
		Readtimeout  int    `yaml: "readtimeout"`
		Wirtetimeout int    `yaml: "wirtetimeout"`
	}
	Security struct {
		Jwt struct {
			Jwtsecret        string `yaml: "jwtsecret"`
			Jwteffectivetime int64  `yaml: "jwteffectivetime"`
		}
	}
}

var (
	//G_cfg_yaml map[string]interface{}
	G_cfg_yaml Config
)

func setup() {

	fmt.Println(" init setting ... ")

	var cfg_file string
	runtimeOS := runtime.GOOS

	if runtimeOS == "windows" {
		cfg_file = "C:\\GOLib\\src\\go-web-scaffold\\conf\\prod.yaml"
	} else {
		cfg_file = "./conf/prod.yaml"
	}

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

func init() {
	setup()
}

func main() {

}
