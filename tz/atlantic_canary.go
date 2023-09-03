// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAtlanticCanaryLocation  sync.Once
	cacheAtlanticCanaryLocation *time.Location
)

type AtlanticCanary struct{}

func (AtlanticCanary) Location() *time.Location {
	onceAtlanticCanaryLocation.Do(func() {
		loc, err := time.LoadLocation("Atlantic/Canary")
		if err != nil {
			panic(err)
		}
		cacheAtlanticCanaryLocation = loc
	})
	return cacheAtlanticCanaryLocation
}
