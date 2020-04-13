package formatters

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// GetTimestampDiffMS returns diff between given timestamps in ms
func GetTimestampDiffMS(startTime *timestamp.Timestamp, endTime *timestamp.Timestamp) (float64, error) {
	startMs := (startTime.Seconds * 1e3) + int64(startTime.Nanos/1e6)
	endMs := (endTime.Seconds * 1e3) + int64(endTime.Nanos/1e6)
	msDiff := (endMs - startMs)
	if msDiff < 0 {
		return 0, errors.New("Start time is less than end time")
	}
	diffSecond := msDiff / 1e3
	diffMs := msDiff % 1e3
	duration := fmt.Sprintf(
		"%d.%d",
		diffSecond,
		diffMs,
	)
	return strconv.ParseFloat(duration, 8)
}
