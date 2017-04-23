package helpers

import (
	"strconv"
	"time"
)

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func ConvTimestampToMillisecond(startTime int64, endTime int64) string {
	millesecond := endTime - startTime
	millesecondString := strconv.FormatInt(millesecond, 10)
	return millesecondString + "ms"
}
