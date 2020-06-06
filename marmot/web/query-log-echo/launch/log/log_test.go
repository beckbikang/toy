package log

import (
	"testing"


	"toy/marmot/web/query-log/launch/config"

	"go.uber.org/zap"
)

func TestLoadGlobalConfig(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	InitLog()

	LOGGER.Info("abc", zap.Int("int", 11))

}
