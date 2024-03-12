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
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
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
	DB       string `mapstructure:"dbname"`
}

func Init() error {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		//load config info failed
		fmt.Printf("viper.ReadInConfig Failed, err:%v\n", err)
		return err
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal Failed, err:%v\n", err)
	}

	// 打印加载的配置文件内容
	fmt.Printf("Loaded configuration: %+v\n", viper.AllSettings())

	// 检查是否发生了错误
	data := viper.AllSettings()
	if len(data) == 0 {
		fmt.Println("配置文件中没有找到任何数据")
	} else {
		fmt.Println("成功加载配置文件")
	}
	fmt.Printf("配置文件中的数据: %+v\n", Conf)
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
