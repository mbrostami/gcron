package formatters_test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mbrostami/gcron/pkg/formatters"
)

func TestGetTimestampDiffMS(t *testing.T) {
	var variables = []struct {
		startTime *timestamp.Timestamp
		endTime   *timestamp.Timestamp
		result    float64
	}{
		{
			startTime: &timestamp.Timestamp{
				Seconds: 1000,
				Nanos:   100000000,
			},
			endTime: &timestamp.Timestamp{
				Seconds: 1001,
				Nanos:   100000000,
			},
			result: 1.0,
		},
		{
			startTime: &timestamp.Timestamp{
				Seconds: 1000,
				Nanos:   100000000,
			},
			endTime: &timestamp.Timestamp{
				Seconds: 1001,
				Nanos:   500000000,
			},
			result: 1.400,
		},
		{
			startTime: &timestamp.Timestamp{
				Seconds: 1001,
				Nanos:   100000000,
			},
			endTime: &timestamp.Timestamp{
				Seconds: 500,
				Nanos:   500000000,
			},
			result: 0,
		},
	}
	for _, item := range variables {
		testname := fmt.Sprintf("startTime: %v, endTime: %v, result: %v", item.startTime.String(), item.endTime.String(), item.result)
		t.Run(testname, func(t *testing.T) {
			res, _ := formatters.GetTimestampDiffMS(item.startTime, item.endTime)
			if res != item.result {
				t.Errorf("got %v, want %v", res, item.result)
			}
		})
	}
}
