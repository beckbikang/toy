package util

import (
	"github.com/speps/go-hashids"
)

var HD *hashids.HashIDData
var HashIdUtil *hashids.HashID

func init() {
	HD = hashids.NewData()
	var err error
	HashIdUtil,err = hashids.NewWithData(HD)
	if err != nil {
		panic("hash init error")
	}


}

func GetHashSliceInt64(numbers []int64)(string,error) {
	return HashIdUtil.EncodeInt64(numbers)
}

