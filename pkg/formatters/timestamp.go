package formatters

import (
	"fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// GetTimestampDiffMS returns diff between given timestamps in ms
func GetTimestampDiffMS(startTime *timestamp.Timestamp, endTime *timestamp.Timestamp) (float64, error) {
	durationSecond := endTime.Seconds - startTime.Seconds
	nanosDiff := (endTime.Nanos - startTime.Nanos) / 1e6
	if nanosDiff < 0 {
		nanosDiff = 0
	}
	duration := fmt.Sprintf(
		"%d.%d",
		durationSecond,
		nanosDiff,
	)
	return strconv.ParseFloat(duration, 8)
}
