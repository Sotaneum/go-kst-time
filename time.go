package ktime

import "time"

func getKST() *time.Location {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	return loc
}

// ToDate : 년, 월, 일, 시, 분을 time.Time 형태로 반환합니다.
func ToDate(year int, month time.Month, day, hour, min int) time.Time {
	t := time.Date(year, month, day, hour, min, 0, 0, getKST())
	return t
}

// ToDateByTime : time.Time을 KST로 변경해서 반환합니다.
func ToDateByTime(t time.Time) time.Time {
	return ToDate(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}

// GetNow : 현재 날짜를 KST로 반환합니다.
func GetNow() time.Time {
	return ToDateByTime(time.Now().In(getKST()))
}

// GetNowWithSecond : 초단위까지 포함한 값을 반환합니다.
func GetNowWithSecond() time.Time {
	return time.Now().In(getKST())
}

// ToTimestampByTime : hour, min 값을 현재 날짜로 변경해 timestamp로 반환합니다.
func ToTimestampByTime(hour, min int) int64 {
	now := GetNow()

	return ToDate(now.Year(), now.Month(), now.Day(), hour, min).Unix()
}

// AddMinuteInTimestamp : timestamp 값에 분 값만큼 더해 반환합니다.
func AddMinuteInTimestamp(timestamp int64, minute int64) int64 {
	return timestamp + (minute * 60)
}
