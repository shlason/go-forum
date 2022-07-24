package constants

import "time"

type cookie struct {
	SessionTokenName   string
	SessionTokenExpiry time.Duration
}

var Cookie = cookie{
	SessionTokenName:   "_SID",
	SessionTokenExpiry: time.Hour * 6,
}
