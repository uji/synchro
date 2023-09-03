// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAfricaAlgiersLocation  sync.Once
	cacheAfricaAlgiersLocation *time.Location
)

type AfricaAlgiers struct{}

func (AfricaAlgiers) Location() *time.Location {
	onceAfricaAlgiersLocation.Do(func() {
		loc, err := time.LoadLocation("Africa/Algiers")
		if err != nil {
			panic(err)
		}
		cacheAfricaAlgiersLocation = loc
	})
	return cacheAfricaAlgiersLocation
}
