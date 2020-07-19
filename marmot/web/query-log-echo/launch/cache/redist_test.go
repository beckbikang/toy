package cache

import (
	"testing"
	"toy/marmot/web/query-log-echo/launch/config"
	"time"
	"github.com/stretchr/testify/assert"
)
func TestRedis(t *testing.T)  {

	ast := assert.New(t)
	config.LoadGlobalConfig("../../conf", "dev")
	InitRedisPool()
	key := "abc"
	v1 := 123
	CommonRedis.Set("abc", v1, time.Duration(130))
	v2,err := CommonRedis.Get(key).Int()
	if err != nil {
		ast.Fail("GET FAILED")
	}
	ast.Equal(v1,v2, "set failed")

	t.Logf("v2=%d",v2)

}