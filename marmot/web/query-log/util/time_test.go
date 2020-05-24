package util


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMillSecond(t *testing.T)  {

	millTime := GetMillSecond()
	t.Logf("millTime:%d", millTime)
	assert.True(t, millTime > 0, "mill time is bigger than 0")
}


func TestGetCurrentYearMonth(t *testing.T){
	t.Logf(GetCurrentYearMonth())
}

func TestGetNextDate(t *testing.T){

	t.Logf("%s", GetNextDate())
}

func TestGetBeforeDate(t *testing.T){
	t.Logf("%s", GetBeforeDate(7))
}