package confer

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var globalConfig *ServerConfig
var mutex sync.RWMutex

type ServerConfig struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	sync.RWMutex
}

type App map[string]interface{}

type Mysql struct {
	DBName string   `mapstructure:"dbname" json:"dbName" yaml:"dbname"`
	Prefix string   `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Pool   DBPool   `mapstructure:"pool" json:"pool" yaml:"pool"`
	Write  DBBase   `mapstructure:"write" json:"write" yaml:"write"`
	Reads  []DBBase `mapstructure:"reads" json:"reads" yaml:"reads"`
}

type DBPool struct {
	PoolMinCap      int   `mapstructure:"pool-min-cap" json:"poolMinCap" yaml:"pool-min-cap"`
	PoolExCap       int   `mapstructure:"pool-ex-cap" json:"poolExCap" yaml:"pool-ex-cap"`
	PoolMaxCap      int   `mapstructure:"pool-max-cap" json:"pool-max-cap" yaml:"pool-max-cap"`
	PoolIdleTimeout int   `mapstructure:"pool-idle-timeout" json:"poolIdleTimeout" yaml:"pool-idle-timeout"`
	PoolWaitCount   int64 `mapstructure:"pool-wait-count" json:"poolWaitCount" yaml:"pool-wait-count"`
	PoolWaitTimeout int   `mapstructure:"pool-wai-timeout" json:"poolWaitTimeout" yaml:"pool-wai-timeout"`
}

type DBBase struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DBName   string `json:"-"`
	Prefix   string `json:"-"`
}

type Log struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	SendToFile  bool   `mapstructure:"send-to-file" json:"send_to_file" yaml:"send-to-file"`
	Filename    string `mapstructure:"filename" json:"filename" yaml:"filename"`
	NoCaller    bool   `mapstructure:"no-calle" json:"no_caller" yaml:"no-caller"`
	NoLogLevel  bool   `mapstructure:"no-log-level" json:"no_log_level" yaml:"no-log-level"`
	Development bool   `mapstructure:"development" json:"development" yaml:"development"`
	MaxSize     int    `mapstructure:"max-size" json:"max_size" yaml:"max-size"` // megabytes
	MaxAge      int    `mapstructure:"max-age" json:"max_age" yaml:"max-age"`    // days
	MaxBackups  int    `mapstructure:"max-backups" json:"max_backups" yaml:"max-backups"`
}

func Init(configURL string) (err error) {
	f, err := os.Open(configURL)
	if err != nil {
		return
	}
	if err = yaml.NewDecoder(f).Decode(&globalConfig); err != nil {
		return
	}
	handleConfig(globalConfig)
	return
}

func handleConfig(config *ServerConfig) {
	config.replaceByEnv(&config.Mysql.Write.Host)
	config.replaceByEnv(&config.Mysql.Write.User)
	config.replaceByEnv(&config.Mysql.Write.Password)
	config.Mysql.Write.DBName = globalConfig.Mysql.DBName
	config.Mysql.Write.Prefix = globalConfig.Mysql.Prefix
}

func (*ServerConfig) replaceByEnv(conf *string) {
	if s := os.Getenv(*conf); len(s) > 0 {
		*conf = s
	}
}

func GlobalConfig() *ServerConfig {
	mutex.RLock()
	defer mutex.RUnlock()
	return globalConfig
}
