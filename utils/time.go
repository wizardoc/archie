package utils

import "time"

/**
 * get Unix time that's expressed at millisecond
 */
func Now() int64 {
	return ParseToMillisecond(time.Now().UnixNano())
}

func ParseToMillisecond(target int64) int64 {
	return target / int64(time.Millisecond)
}
