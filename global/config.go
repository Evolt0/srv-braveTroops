package global

import (
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/Evolt0/fabric-sdk-yml/base"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	InitOnce sync.Once
	Conf     *Config
)

// 初始化
func init() {
	InitConfig()
	Conf.Client.Init()
}

// 配置文件读取
type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"app"`
	Client *base.Client `yaml:"client"`
}

// 初始化读取配置文件
func InitConfig() {
	InitOnce.Do(
		func() {
			pwd, err := os.Getwd()
			if err != nil {
				logrus.Fatalf("failed to get current directory %v", err)
			}
			file, err := ioutil.ReadFile(path.Join(pwd, "config", "local.yaml"))
			if err != nil {
				logrus.Fatalf("failed to load file! %v", err)
			}
			Conf = &Config{}
			err = yaml.Unmarshal(file, Conf)
			if err != nil {
				logrus.Fatalf("failed to unmarshal file! %v", err)
			}
			logrus.Infoln("config file init successfully!")
		})
}
