package util




import (
"time"
"fmt"
"strconv"
)


const(
	TimeYearMonth = "200601"
	TimeDate = "2006-01-02"
	Time2DatetimeFormat = "2006-01-02 15:04:05"
	Time2DatetimeFormatShort="20060102150405"
	SecondOneDay = 86400
	TimeDateFull = "2006-01-02 00:00:00"

)

var(
	asiaLocal *time.Location
)

func init() {
	asiaLocal,_ = time.LoadLocation("Asia/Shanghai")
}

//返回当前的年月
func GetCurrentYearMonth()string{
	return time.Now().In(asiaLocal).Format(TimeYearMonth)
}






func GetCurrentDate()string{
	return time.Now().In(asiaLocal).Format(TimeDate)
}

func GetCurrentTime() string{
	return time.Now().In(asiaLocal).Format(Time2DatetimeFormat)
}


func GetMillSecond() int64 {
	now := time.Now()
	return now.UnixNano()/1e6 - now.Unix() * 1000
}



func UnixTimeToDateTime(timestamp int64)int64{

	timestr := time.Unix(timestamp,0).Format(Time2DatetimeFormatShort)

	timeInt,err := strconv.ParseInt(timestr,10,64)
	if err != nil{
		return 0
	}
	return timeInt
}


/**
("yyyyMMdd", "HHmmss", error)
 */
func TimeSplite(timestamp string)(string,string, error){
	if len(timestamp) != 14{
		return "","", fmt.Errorf("len error")
	}
	return timestamp[0:8],timestamp[8:14],nil
}

func StrToUnix(timeStr string)(int64,error){
	t, err := time.Parse("2006-01-02 15:04:05 MST", timeStr+" CST")
	if err != nil{
		return 0, fmt.Errorf("format error")
	}
	return t.Unix(),nil
}

func UnixToStr (timestamp int64) string{
	t := time.Unix(timestamp,0)
	return t.Format(Time2DatetimeFormat)
}

//get time
func GetDateByTimestamp(timestamp int64) string{
	t := time.Unix(timestamp,0)
	return t.Format(TimeDate)
}

//get one date
func GetNextDate() string{
	return time.Now().In(asiaLocal).Add(SecondOneDay*time.Second).Format(TimeDateFull)
}

func GetBeforeDate(b int) string{
	ts := time.Now().In(asiaLocal).Unix()-int64(SecondOneDay*b)

	return time.Unix(ts,0).Format(TimeDateFull)
}