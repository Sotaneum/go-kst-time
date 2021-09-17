# go-kst-time

docker의 리눅스 계열 사용시, KST 기준으로 time 값을 사용할 수 있도록 합니다.

## Testing

```bash
go test
```

## 예시

```go
import (
    ktime "github.com/Sotaneum/go-kst-time"
)
ktime.ToDate(2021, 2, 14, 12, 0)  // 2021-02-14 12:00:00 +0900 KST
ktime.ToDateByTime(time.Now())  // 2021-09-18 03:03:00 +0900 KST
ktime.GetNow()  // 2021-09-18 03:03:00 +0900 KST
ktime.GetNowWithSecond()  // 2021-09-18 03:03:06.010944 +0900 KST
ktime.ToTimestampByTime(12, 0)  // 1631934000
ktime.AddMinuteInTimestamp(ktime.ToTimestampByTime(12, 0), 30)  // 1631935800
```

## MIT License

Copyright (c) 2021 Sotaneum
