// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceEuropeMaltaLocation  sync.Once
	cacheEuropeMaltaLocation *time.Location
)

type EuropeMalta struct{}

func (EuropeMalta) Location() *time.Location {
	onceEuropeMaltaLocation.Do(func() {
		loc, err := time.LoadLocation("Europe/Malta")
		if err != nil {
			panic(err)
		}
		cacheEuropeMaltaLocation = loc
	})
	return cacheEuropeMaltaLocation
}
