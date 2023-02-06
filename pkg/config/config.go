/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:32:30
 * @LastEditTime: 2023-02-07 00:53:49
 * @Description:
 */
package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	// C 全局配置文件，在Init调用前为nil
	C    *Config
)

const (
	UserPort          = ":8801"
	VideoPort         = ":8802"
	FavoritePort      = ":8803"
	CommentPort       = ":8804"
	RelationPort      = ":8805"
	UserHostPorts     = "0.0.0.0" + UserPort
	VideoHostPorts    = "0.0.0.0" + VideoPort
	FavoriteHostPorts = "0.0.0.0" + FavoritePort
	CommentHostPorts  = "0.0.0.0" + CommentPort
	RelationHostPorts = "0.0.0.0" + RelationPort
)

// Config 配置
type Config struct {
	App     app     `yaml:"app"`
	DB      db      `yaml:"db"`
	Redis   redis   `yaml:"redis"`
	JWT     jwt     `yaml:"jwt"`
	Wechat  wechat  `yaml:"wechat"`
	Postgresql  postgresql  `yaml:"postgresql"`
	LogConf logConf `yaml:"logConf"`
	Debug   bool    `yaml:"debug"`	
	SQLDebug       bool     `yaml:"sql_debug,omitempty"`
}

type app struct {
	Addr       string `yaml:"addr"`
	Prefix     string `yaml:"prefix"`
	TimeFormat string `yaml:"timeformat"`
}

type db struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Timeout  int    `yaml:"timeout_second"`
}

type redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}

type wechat struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

type logConf struct {
	LogPath     string `yaml:"log_path"`
	LogFileName string `yaml:"log_file_name"`
}

type postgresql struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Dbname   string `yaml:"dbname,omitempty"`
}

func init() {
	InitConfig()
}

func InitConfig() {
	env := "dev"
	// 如果有设置 ENV ，则使用ENV中的环境
	if v, ok := os.LookupEnv("ENV"); ok {
		env = v
	}
	configFile := fmt.Sprintf("%s.yml", env)
	data, err := os.ReadFile(fmt.Sprintf("../config/%s", configFile))
	if err != nil {
		log.Println("Read config error!")
		return
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Println("Unmarshal config error!")
		return
	}
	C = config
	log.Println("\033[36m" + "[INFO]" + "\033[0m"+" Config "+configFile+" loaded.")
}
