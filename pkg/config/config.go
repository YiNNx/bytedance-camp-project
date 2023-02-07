/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:32:30
 * @LastEditTime: 2023-02-07 22:39:59
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

// Config 配置
type Config struct {
	Debug   bool    `yaml:"debug"`	
	SQLDebug       bool     `yaml:"sql_debug,omitempty"`

	App     app     `yaml:"app"`
	Postgresql  postgresql  `yaml:"postgresql"`
	JWT     jwt     `yaml:"jwt"`
	Service service `yaml:"service"`
}

type app struct {
	Addr       string `yaml:"addr"`
	Prefix     string `yaml:"prefix"`
	TimeFormat string `yaml:"timeformat"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}

type postgresql struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Dbname   string `yaml:"dbname,omitempty"`
}

type service struct {
	UserHostPort     string `yaml:"user_host_port"`
	VideoHostPort    string `yaml:"video_host_port"`
	FavoriteHostPort string `yaml:"favorite_host_port"`
	CommentHostPort  string `yaml:"comment_host_port"`
}

func init() {
	InitConfig()
}

func InitConfig() {
	data, err := os.ReadFile(fmt.Sprintf("./config/dev.yml"))
	if err != nil {
		log.Println("Read config error!\n"+err.Error())
		return
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Println("Unmarshal config error!"+err.Error())
		return
	}
	C = config
	log.Println("Config loaded.")
	// log.Println(C)
}
