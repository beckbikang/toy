package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

var (
	sf     *sonyflake.Sonyflake
	engine *gin.Engine
)

const (
	Ok     = 200
	Failed = 400
)

type Result struct {
	Code int
	Id   string
}

func getMachinId() (uint16, error) {
	return 123, nil
}
func checkMach(uint16) bool {
	return true
}

func init() {
	var st sonyflake.Settings
	st.MachineID = getMachinId
	st.CheckMachineID = checkMach
	st.StartTime, _ = time.Parse("2006-01-02", "2020-05-05")
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("flake failed")
	}
	engine = gin.New()
	engine.GET("/id", func(c *gin.Context) {
		var re Result
		id, err := sf.NextID()
		if err != nil {
			re.Code = Failed
			c.SecureJSON(http.StatusOK, re)
		}
		re.Code = Ok
		re.Id = strconv.FormatUint(id, 10)
		c.SecureJSON(http.StatusOK, re)
	})
}

func main() {

	engine.Run(":8181")

	endless.ListenAndServe(":4242", engine)
}
