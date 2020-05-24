package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGlobalConfig(t *testing.T) {
	asrt := assert.New(t)

	LoadGlobalConfig("../../conf", "dev")

	t.Logf("%d", Gcfg.GetInt("server.read_timeout"))
	asrt.Equal(60, Gcfg.GetInt("server.read_timeout"), "timeout not ok")
}
