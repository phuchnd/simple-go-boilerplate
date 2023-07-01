package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertTimeStr(input *timestamppb.Timestamp, loc *time.Location) string {
	if !input.IsValid() {
		return ""
	}
	inputTime := input.AsTime()
	if inputTime.IsZero() {
		return ""
	}
	return inputTime.In(loc).Format(time.RFC3339)
}
