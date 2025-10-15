package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// timestampToDatetime converts time to formatted string in specified timezone
func timestampToDatetime(convertedTime time.Time, originTimezone, targetTimezone string) (string, error) {
	// Load origin timezone
	originLoc, err := time.LoadLocation(originTimezone)
	if err != nil {
		return "", err
	}

	// Load target timezone
	targetLoc, err := time.LoadLocation(targetTimezone)
	if err != nil {
		return "", err
	}

	// Convert time to origin timezone
	convertedTime = convertedTime.In(originLoc)
	// Convert to target timezone
	convertedTimeInTarget := convertedTime.In(targetLoc)
	// Format time
	dateStr := convertedTimeInTarget.Format("2006-01-02 15:04:05")
	return dateStr, nil
}

// IosTimestampToDatetime converts iOS timestamp (seconds since 2001-01-01)
func IosTimestampToDatetime(timestamp float64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	baseTime := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Duration(timestamp * float64(time.Second))
	convertedTime := baseTime.Add(duration)

	result, err := timestampToDatetime(convertedTime, originTimezone, targetTimezone)
	if err != nil {
		return ""
	}
	return result
}

// DefaultTimestampToDatetime automatically determines unit by timestamp length and converts
func DefaultTimestampToDatetime(timestamp int64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	tStr := strconv.FormatInt(timestamp, 10)
	tLength := len(tStr)

	var dt time.Time

	switch {
	case tLength == 13: // milliseconds
		dt = time.Unix(timestamp/1000, (timestamp%1000)*int64(time.Millisecond))
	case tLength == 16: // microseconds
		dt = time.Unix(timestamp/1000000, (timestamp%1000000)*int64(time.Microsecond))
	case tLength == 19: // nanoseconds
		dt = time.Unix(timestamp/1000000000, timestamp%1000000000)
	default: // seconds
		dt = time.Unix(timestamp, 0)
	}

	result, err := timestampToDatetime(dt, originTimezone, targetTimezone)
	if err != nil {
		return ""
	}
	return result
}

// ChromeTimestampToDatetime converts Chrome timestamp (microseconds since 1601-01-01)
func ChromeTimestampToDatetime(timestamp int64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	baseTime := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Duration(timestamp) * time.Microsecond
	convertedTime := baseTime.Add(duration)

	result, err := timestampToDatetime(convertedTime, originTimezone, targetTimezone)
	if err != nil {
		return ""
	}
	return result
}

// WindowsFileTimeToDatetime converts Windows file time (100-nanosecond intervals since 1601-01-01)
func WindowsFileTimeToDatetime(timestamp uint64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	baseTime := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Duration(timestamp/10) * time.Microsecond
	convertedTime := baseTime
	if duration < 0 { // Handle overflow when offset is too large
		maxDuration := 1<<63 - 1
		minDuration := -1 << 63
		convertedTime = convertedTime.Add(time.Duration(maxDuration))
		convertedTime = convertedTime.Add(duration - time.Duration(minDuration))
	} else {
		convertedTime = convertedTime.Add(duration)
	}
	result, err := timestampToDatetime(convertedTime, originTimezone, targetTimezone)
	if err != nil {
		return ""
	}
	return result
}

// NineTimestampToDatetime converts special 9-digit timestamp
func NineTimestampToDatetime(timestamp int64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	nine := timestamp - 621355968000000000
	return DefaultTimestampToDatetime(nine/10000, originTimezone, targetTimezone)
}

// AppleTimestampToDatetime converts Apple timestamp
func AppleTimestampToDatetime(timestamp float64, originTimezone, targetTimezone string) string {
	if originTimezone == "" {
		originTimezone = "UTC"
	}
	if targetTimezone == "" {
		targetTimezone = "Asia/Shanghai"
	}
	timestampStr := fmt.Sprintf("%.20f", timestamp)
	parts := strings.Split(timestampStr, ".")
	if len(parts) < 1 {
		return ""
	}

	secondsStr := parts[0]
	if len(secondsStr) < 9 {
		return ""
	}

	seconds, err := strconv.ParseInt(secondsStr[:9], 10, 64)
	if err != nil {
		return ""
	}

	seconds += 978307200 // Seconds from 2001-01-01 to 1970-01-01
	dt := time.Unix(seconds, 0)

	result, err := timestampToDatetime(dt, originTimezone, targetTimezone)
	if err != nil {
		return ""
	}
	return result
}
