package timestamp

import (
	"MobileForensicsToolkit/utils"
	"context"
	"strconv"
)

// TimeStampParser handles various timestamp format conversions for forensic analysis
type TimeStampParser struct {
	ctx context.Context
}

// NewTimeStampParser creates a new TimeStampParser instance
func NewTimeStampParser() *TimeStampParser {
	return &TimeStampParser{}
}

// InitCtx initializes the context for the parser
func (t *TimeStampParser) InitCtx(ctx context.Context) {
	t.ctx = ctx
}

// ConvertUnixTimestamp converts Unix timestamp to readable format
func (t *TimeStampParser) ConvertUnixTimestamp(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "Invalid timestamp format"
	}
	return utils.DefaultTimestampToDatetime(ts, sourceTimezone, targetTimezone)
}

// ConvertIosTimestamp converts iOS timestamp (seconds since 2001-01-01) to readable format
func (t *TimeStampParser) ConvertIosTimestamp(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseFloat(timestamp, 64)
	if err != nil {
		return "Invalid iOS timestamp format"
	}
	return utils.IosTimestampToDatetime(ts, sourceTimezone, targetTimezone)
}

// ConvertChromeTimestamp converts Chrome timestamp (microseconds since 1601-01-01) to readable format
func (t *TimeStampParser) ConvertChromeTimestamp(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "Invalid Chrome timestamp format"
	}
	return utils.ChromeTimestampToDatetime(ts, sourceTimezone, targetTimezone)
}

// ConvertWindowsFileTime converts Windows file time (100-nanosecond intervals since 1601-01-01) to readable format
func (t *TimeStampParser) ConvertWindowsFileTime(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseUint(timestamp, 10, 64)
	if err != nil {
		return "Invalid Windows file time format"
	}
	return utils.WindowsFileTimeToDatetime(ts, sourceTimezone, targetTimezone)
}

// ConvertAppleTimestamp converts Apple timestamp to readable format
func (t *TimeStampParser) ConvertAppleTimestamp(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseFloat(timestamp, 64)
	if err != nil {
		return "Invalid Apple timestamp format"
	}
	return utils.AppleTimestampToDatetime(ts, sourceTimezone, targetTimezone)
}

// ConvertNineDigitTimestamp converts special 9-digit timestamp format
func (t *TimeStampParser) ConvertNineDigitTimestamp(timestamp, sourceTimezone, targetTimezone string) string {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "Invalid 9-digit timestamp format"
	}
	return utils.NineTimestampToDatetime(ts, sourceTimezone, targetTimezone)
}

// BatchConvertTimestamps converts multiple timestamps in different formats
func (t *TimeStampParser) BatchConvertTimestamps(timestamps []string, format, sourceTimezone, targetTimezone string) []map[string]string {
	results := make([]map[string]string, 0, len(timestamps))
	
	for _, ts := range timestamps {
		result := map[string]string{
			"input": ts,
			"format": format,
		}
		
		switch format {
		case "unix":
			result["converted"] = t.ConvertUnixTimestamp(ts, sourceTimezone, targetTimezone)
		case "ios":
			result["converted"] = t.ConvertIosTimestamp(ts, sourceTimezone, targetTimezone)
		case "chrome":
			result["converted"] = t.ConvertChromeTimestamp(ts, sourceTimezone, targetTimezone)
		case "windows":
			result["converted"] = t.ConvertWindowsFileTime(ts, sourceTimezone, targetTimezone)
		case "apple":
			result["converted"] = t.ConvertAppleTimestamp(ts, sourceTimezone, targetTimezone)
		case "nine":
			result["converted"] = t.ConvertNineDigitTimestamp(ts, sourceTimezone, targetTimezone)
		default:
			result["converted"] = "Unknown format"
		}
		
		results = append(results, result)
	}
	
	return results
}
