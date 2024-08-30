package utils

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration parse duration
func ParseDuration(d string) (time.Duration, error) {
	// trim space
	d = strings.TrimSpace(d)
	// parse duration
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}

	// if contains "d", for ez use
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	// parse int64
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func StrToUInt(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	return uint(i), err
}
