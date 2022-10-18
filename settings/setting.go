package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"runtime"
	"strings"
)

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Dbnum    int    `mapstructure:"dbnum"`
	Password string `mapstructure:"password"`
}

type Api struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Email struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	From        string `mapstructure:"from"`
	Subject     string `mapstructure:"subject"`
	AccessToken string `mapstructure:"accessToken"`
}

type Config struct {
	MysqlConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
	Api         `mapstructure:"api"`
	Email       `mapstructure:"email"`
}

var (
	storePath string
)

var (
	ApiConfig Config
	basePath  string
)

func Init() (err error) {
	_, basePath, _, _ = runtime.Caller(0)
	basePaths := strings.Split(basePath, "/")
	basePath = strings.Join(basePaths[:len(basePaths)-1], "/")
	storePath = strings.Join(basePaths[:len(basePaths)-2], "/") + "/data/"
	fmt.Println(storePath)
	viper.AddConfigPath(basePath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	err = viper.Unmarshal(&ApiConfig)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变，注意更改!!!\a")
	})
	return
}

func Mysql() MysqlConfig {
	return ApiConfig.MysqlConfig
}

func Redis() RedisConfig {
	return ApiConfig.RedisConfig
}

func API() Api {
	return ApiConfig.Api
}

func EmailSetting() Email {
	return ApiConfig.Email
}

func BasePath() string {
	return basePath
}

func StorePath() string {
	return storePath
}
