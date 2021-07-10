package configs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
	"strings"
	"sync"
)

var (
	configFolder = "configs"
)

type config struct {
	Mode string `yaml:"mode"`
	Log  struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
	Mongo struct {
		Db               string `yaml:"db"`
		Host             string `yaml:"host"`
		Port             int    `yaml:"port"`
		Username         string `yaml:"username"`
		Password         string `yaml:"password"`
		ConnectionString string
	} `yaml:"mongo"`

	Strategy struct {
		Bet struct {
			ConfidenceBase struct {
				Threshold float64 `yaml:"threshold"`
			} `yaml:"confidenceBase"`
		} `yaml:"bet"`
	} `yaml:"strategy"`
}

var (
	once     sync.Once
	instance *config
)

func load(filename string, config interface{}) {
	viper.SetConfigName(filename)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("fail to load config file: ", err.Error())
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("fail to decode config: ", err.Error())
	}
	return
}

func New() *config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configFolder)
		load("config", &instance)

		instance.initLog()

		if instance.Mongo.Username != "" && instance.Mongo.Password != "" {
			instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%d", instance.Mongo.Username, instance.Mongo.Password, instance.Mongo.Host, instance.Mongo.Port)
		} else {
			instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%d", instance.Mongo.Host, instance.Mongo.Port)
		}

		log.Debug("config initialized")
	})
	return instance
}

func (c *config) initLog() {
	logLevel := map[string]log.Level{
		"DEBUG": log.DebugLevel,
		"INFO":  log.InfoLevel,
		"WARN":  log.WarnLevel,
		"ERROR": log.ErrorLevel,
		"FATAL": log.FatalLevel,
	}

	callerFormatter := func(path string) string {
		arr := strings.Split(path, "/")
		return arr[len(arr)-1]
	}
	customFormatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", callerFormatter(f.File), f.Line)
		},
	}

	log.SetLevel(logLevel[c.Log.Level])
	log.SetFormatter(customFormatter)
	log.SetReportCaller(true)
	log.Debug("logger initialized")
}
