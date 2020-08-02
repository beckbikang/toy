package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetHashSliceInt64(t *testing.T)  {
	l1 := make([]int64, 3)

	l1[0] = 1
	l1[1] = 2
	l1[2] = 3

	key, err := GetHashSliceInt64(l1)

	assert.Equal(t,true , err == nil)

	t.Logf("err=%v,key=%s",err, key)

}