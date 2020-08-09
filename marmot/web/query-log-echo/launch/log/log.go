package log

import (
	sl "log"
	"github.com/beckbikang/flg"
	"go.uber.org/zap"
	cfg "toy/marmot/web/query-log-echo/launch/config"
)

var (
	LOGGER  *zap.Logger
	Gflg *flg.Logger
)

func InitLog() *zap.Logger {
	sl.Println("init log ...")

	Gflg = &flg.Logger{}

	var lcfg flg.FConfig
	lcfg.Lfg.Filename = cfg.Gcfg.GetString("jackcfg.filename")
	lcfg.Lfg.MaxSize = cfg.Gcfg.GetInt("jackcfg.maxsize")
	lcfg.Lfg.MaxAge = cfg.Gcfg.GetInt("jackcfg.maxage")
	lcfg.Lfg.MaxBackups = cfg.Gcfg.GetInt("jackcfg.maxbackups")
	lcfg.Lfg.LocalTime = cfg.Gcfg.GetBool("jackcfg.localtime")
	lcfg.Lfg.Compress = cfg.Gcfg.GetBool("jackcfg.compress")

	lcfg.Zfgs = make(map[string]flg.ZConfig)

	sl.Printf("level=%s", cfg.Gcfg.GetString("zapcfgs.1.level"))

	lcfg.Zfgs["1"] = flg.ZConfig{
		Level:cfg.Gcfg.GetString("zapcfgs.1.level"),
		IsDev:      cfg.Gcfg.GetBool("zapcfgs.1.isdev"),
		LogMod:     int8(cfg.Gcfg.GetInt("zapcfgs.1.logmod")),
		ServerName: cfg.Gcfg.GetString("zapcfgs.1.servername"),
	}

	var err error
	err = Gflg.LoadFromObject(&lcfg)
	if err != nil {
		panic("LoadFromObject failed")
	}
	LOGGER, err = Gflg.GetLogByKey(lcfg.Zfgs["1"].ServerName)
	if err != nil {
		panic("get log failed")
	}
	sl.Println("")
	sl.Printf("init log end %+v\n", lcfg)
	LOGGER.Info("log is ok , it's running!")
	return LOGGER
}
