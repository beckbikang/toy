package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	globalConfObject *Configer
	Gcfg             *viper.Viper
)

func init() {
	globalConfObject = new(Configer)
	Gcfg = viper.New()
}

func LoadGlobalConfig(configPath, configFileName string) {
	Gcfg.AddConfigPath(configPath)
	Gcfg.SetConfigName(configFileName)
	Gcfg.SetConfigType("toml")

	if err := Gcfg.ReadInConfig(); err != nil {
		fmt.Printf("error :%v\n", err)
		panic(-1)
	}

	err := Gcfg.Unmarshal(globalConfObject)

	if err != nil {
		fmt.Printf("error :%v\n", err)
		panic(-1)
	}

	showSetting()
	Gcfg.WatchConfig()
	Gcfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		showSetting()
	})

}

func showSetting() {
	fmt.Sprintf("globalConfObject:%+v", globalConfObject)
}

type Configer struct {
	Server ServerCfg `toml:"server"`
	App    AppCfg    `toml:"app"`
	Log    LogCfg    `toml:"log"`
	Mysql  MysqlCfg  `toml:"mysql"`
}

type ServerCfg struct {
	Address      string `toml:"addr'`
	RunMod       string `toml:"run_mode'`
	ReadTimeout  int    `toml:"read_timeout'`
	WriteTimeout int    `toml:"write_timeout'`
	ConfigPath   string `toml:"config_path'`
	MaxBodyBytes string `toml:"max_body_bytes"`
}

type AppCfg struct {
}

type LogCfg struct {
}

type MysqlCfg struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Pswd         string `toml:"pswd"`
	Db           string `toml:"db"`
	Charset      string `toml:"charset"`
	Locale       string `toml:"locale"`
	Lifetime     int    `toml:"conn_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
}

type KafkaConfig struct {
	Brokers       string `toml:"brockers"`
	Topic         string `toml:"topic"`
	Group         string `toml:"group"`
	User          string `toml:"user"`
	Pswd          string `toml:"password"`
	FromBeginning bool   `toml:"from_beginning"`
}
