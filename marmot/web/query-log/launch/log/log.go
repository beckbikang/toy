package log

import (
	"github.com/beckbikang/flg"
	"go.uber.org/zap"
	cfg "toy/marmot/web/query-log/launch/config"
)

var (
	Flg  *zap.Logger
	gflg *flg.Logger
)

func InitLog() *zap.Logger {
	var lcfg flg.FConfig
	lcfg.Lfg.Filename = cfg.Gcfg.GetString("jackcfg.filename")
	lcfg.Lfg.MaxSize = cfg.Gcfg.GetInt("jackcfg.maxsize")
	lcfg.Lfg.MaxAge = cfg.Gcfg.GetInt("jackcfg.maxage")
	lcfg.Lfg.MaxBackups = cfg.Gcfg.GetInt("jackcfg.maxbackups")
	lcfg.Lfg.LocalTime = cfg.Gcfg.GetBool("jackcfg.localtime")
	lcfg.Lfg.Compress = cfg.Gcfg.GetBool("jackcfg.compress")
	lcfg.Zfgs = make(map[string]flg.ZConfig)
	lcfg.Zfgs["1"] = flg.ZConfig{Level: Gcfg.GetString("zapcfgs.1.level"),
		IsDev:      cfg.Gcfg.GetBool("zapcfgs.1.isdev"),
		LogMod:     int8(cfg.Gcfg.GetInt("zapcfgs.1.logmod")),
		ServerName: cfg.Gcfg.GetString("zapcfgs.1.servername")}
	var err error
	err = gflg.LoadFromObject(&lcfg)
	if err != nil {
		panic("LoadFromObject faild")
	}
	Flg, err = l.GetLogByKey(fconfig.Zfgs["1"].ServerName)
	if err != nil {
		panic("get log faild")
	}

	return Flg
}
