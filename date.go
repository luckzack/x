package utils

import (
	"time"
)

func UnixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func UnixTsFormatMinute(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04")
}



func UnixMillisecond() int64 {
	nano := time.Now().UnixNano()
	return int64(nano/1000000)
}

func NowDate() string{
	return UnixTsFormat(time.Now().Unix())
}

//"2016-11-02 05:14:25"
func Parse(timestr string) (int64, error){
	//t, err := time.ParseInLocation("2006-01-02 15:04:05", "2016-11-02 05:14:25", time.Local)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	if err != nil{
		return 0, err
	}
	return t.Unix(), nil
}

func DayStartTime(ti time.Time) (int64, error){
	return Parse(ti.Format("2006-01-02 00:00:00"))
}

func DayStart(ti time.Time) (string){
	return ti.Format("2006-01-02 00:00:00")
}

func DayEndTime(ti time.Time) (int64, error) {
	i, err := DayStartTime(ti.Add(time.Hour * 24))
	if err != nil {
		return 0, err
	}
	return i - 1, nil
}

func IsWeekend(t time.Time) bool{
	switch t.Weekday().String() {
	case "Monday":
		fallthrough
	case "Tuesday":
		fallthrough
	case "Wednesday":
		fallthrough
	case "Thursday":
		return false
	case "Friday":
		fallthrough
	case "Saturday":
		fallthrough
	case "Sunday":
		return true
	}
	return true
}


