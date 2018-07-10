package time

import (
	"time"
)

func Format(t time.Time) string {
	return t.Format(YYYYMMDDhhmmss)
}
func UnixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format(YYYYMMDDhhmmss)
}

func UnixTsFormatMinute(ts int64) string {
	return time.Unix(ts, 0).Format(YYYYMMDDhhmm)
}

func UnixMillisecond() int64 {
	nano := time.Now().UnixNano()
	return int64(nano / 1000000)
}

func NowDate() string {
	return UnixTsFormat(time.Now().Unix())
}

//"2016-11-02 05:14:25"
func Parse(timestr string) (int64, error) {
	//t, err := time.ParseInLocation("2006-01-02 15:04:05", "2016-11-02 05:14:25", time.Local)
	t, err := time.ParseInLocation(YYYYMMDDhhmmss, timestr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func ParseNoSecond(timestr string) (int64, error) {
	//t, err := time.ParseInLocation("2006-01-02 15:04:05", "2016-11-02 05:14:25", time.Local)
	t, err := time.ParseInLocation(YYYYMMDDhhmm, timestr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func DayStartTime(ti time.Time) (int64, error) {
	return Parse(ti.Format(YYYYMMDD000000))
}

func DayStart(ti time.Time) string {
	return ti.Format(YYYYMMDD000000)
}

func DayEndTime(ti time.Time) (int64, error) {
	i, err := DayStartTime(ti.Add(time.Hour * 24))
	if err != nil {
		return 0, err
	}
	return i - 1, nil
}

func IsWeekend(t time.Time) bool {
	switch t.Weekday().String() {
	case "Monday":
		fallthrough
	case "Tuesday":
		fallthrough
	case "Wednesday":
		fallthrough
	case "Thursday":
		fallthrough
	case "Friday":
		return false
	case "Saturday":
		fallthrough
	case "Sunday":
		return true
	}
	return true
}
