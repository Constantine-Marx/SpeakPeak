package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf global variable
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*MySQLConfig `mapstructure:"mysql"`
	//Redis *RedisConfig `mapstructure:"redis"`
	*LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host" `
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password" `
	DB       string `mapstructure:"db"`
}

func Init() error {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	fmt.Printf("1..%d", viper.GetInt("app.port"))
	err := viper.ReadInConfig()
	if err != nil {
		//load config info failed
		fmt.Printf("viper.ReadInConfig Failed, err:%v\n", err)
		return err
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal Failed, err:%v\n", err)
		return err
	}
	fmt.Printf("2..%d", viper.GetInt("app.port"))

	viper.WatchConfig()
	//hook
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Hot reload success!!")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal Failed, err:%v\n", err)
		}
	})
	return err
}
