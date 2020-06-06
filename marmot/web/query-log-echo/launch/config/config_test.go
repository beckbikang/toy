package config

import (
	"testing"

	"github.com/beckbikang/flg"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoadGlobalConfig(t *testing.T) {
	asrt := assert.New(t)

	LoadGlobalConfig("../../conf", "dev")

	t.Logf("%d", Gcfg.GetInt("server.read_timeout"))
	t.Logf("%s", Gcfg.GetString("jackcfg.filename"))
	t.Logf("%s", Gcfg.GetString("zapcfgs.1.level"))
	asrt.Equal(60, Gcfg.GetInt("server.read_timeout"), "timeout not ok")

	t.Logf("GetGlobalConfObject:%v", GetGlobalConfObject().Lfg)

	l := &flg.Logger{}

	var fconfig flg.FConfig
	fconfig.Lfg.Filename = Gcfg.GetString("jackcfg.filename")
	fconfig.Lfg.MaxSize = Gcfg.GetInt("jackcfg.maxsize")
	fconfig.Lfg.MaxAge = Gcfg.GetInt("jackcfg.maxage")
	fconfig.Lfg.MaxBackups = Gcfg.GetInt("jackcfg.maxbackups")
	fconfig.Lfg.LocalTime = Gcfg.GetBool("jackcfg.localtime")
	fconfig.Lfg.Compress = Gcfg.GetBool("jackcfg.compress")
	fconfig.Zfgs = make(map[string]flg.ZConfig)
	fconfig.Zfgs["1"] = flg.ZConfig{Level: Gcfg.GetString("zapcfgs.1.level"),
		IsDev:      Gcfg.GetBool("zapcfgs.1.isdev"),
		LogMod:     int8(Gcfg.GetInt("zapcfgs.1.logmod")),
		ServerName: Gcfg.GetString("zapcfgs.1.servername")}

	t.Logf("%v", fconfig)

	err := l.LoadFromObject(&fconfig)
	if err != nil {
		panic("LoadFromObject faild")
	}
	lg, err := l.GetLogByKey(fconfig.Zfgs["1"].ServerName)
	defer lg.Sync()
	lg.Info("a test")
	lg.Info("abc", zap.Int("int", 11))

}
