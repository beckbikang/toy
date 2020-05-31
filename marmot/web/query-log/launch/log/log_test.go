package log

import (
	"testing"

	"toy/marmot/web/query-log/launch/config"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoadGlobalConfig(t *testing.T) {
	asrt := assert.New(t)

	config.LoadGlobalConfig("../../conf", "dev")
	InitLog()

	lg.Info("a test")
	lg.Info("abc", zap.Int("int", 11))

}
