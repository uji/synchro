// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaLimaLocation  sync.Once
	cacheAmericaLimaLocation *time.Location
)

type AmericaLima struct{}

func (AmericaLima) Location() *time.Location {
	onceAmericaLimaLocation.Do(func() {
		loc, err := time.LoadLocation("America/Lima")
		if err != nil {
			panic(err)
		}
		cacheAmericaLimaLocation = loc
	})
	return cacheAmericaLimaLocation
}
