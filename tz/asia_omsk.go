// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaOmskLocation  sync.Once
	cacheAsiaOmskLocation *time.Location
)

type AsiaOmsk struct{}

func (AsiaOmsk) Location() *time.Location {
	onceAsiaOmskLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Omsk")
		if err != nil {
			panic(err)
		}
		cacheAsiaOmskLocation = loc
	})
	return cacheAsiaOmskLocation
}
