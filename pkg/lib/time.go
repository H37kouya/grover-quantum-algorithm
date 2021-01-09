package lib

import (
	"time"
)

const layoutForFileName = "2006-01-02_15-04-05"

func GetTimeForFileName() string {
	t := time.Now()
	return t.Format(layoutForFileName)
}
