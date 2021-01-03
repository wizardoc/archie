package utils

import "time"

/**
 * get Unix time that's expressed at millisecond
 */
func Now() int32 {
	return ParseToMillisecond(time.Now().UnixNano())
}

func ParseToMillisecond(target int64) int32 {
	return int32(target / int64(time.Millisecond))
}
