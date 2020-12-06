package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
	*LogConfig   `mapstructure:"log"`
	*Neo4jConfig `mapstructure:"neo4j"`
}

type Neo4jConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		panic(fmt.Errorf("Fatal error while reading config file: %v\n", err))
		return
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file modified...")
		viper.Unmarshal(&Conf)
	})

	return
}
