package ktime_test

import (
	"fmt"
	"testing"
	"time"

	ktime "github.com/Sotaneum/go-kst-time"
)

func TestRunner(t *testing.T) {
	fmt.Println(ktime.ToDate(2021, 2, 14, 12, 0))
	fmt.Println(ktime.ToDateByTime(time.Now()))
	fmt.Println(ktime.GetNow())
	fmt.Println(ktime.GetNowWithSecond())
	fmt.Println(ktime.ToTimestampByTime(12, 0))
	fmt.Println(ktime.AddMinuteInTimestamp(ktime.ToTimestampByTime(12, 0), 30))
}
