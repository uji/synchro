// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAmericaSantiagoLocation  sync.Once
	cacheAmericaSantiagoLocation *time.Location
)

type AmericaSantiago struct{}

func (AmericaSantiago) Location() *time.Location {
	onceAmericaSantiagoLocation.Do(func() {
		loc, err := time.LoadLocation("America/Santiago")
		if err != nil {
			panic(err)
		}
		cacheAmericaSantiagoLocation = loc
	})
	return cacheAmericaSantiagoLocation
}
