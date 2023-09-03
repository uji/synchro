// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaJakartaLocation  sync.Once
	cacheAsiaJakartaLocation *time.Location
)

type AsiaJakarta struct{}

func (AsiaJakarta) Location() *time.Location {
	onceAsiaJakartaLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}
		cacheAsiaJakartaLocation = loc
	})
	return cacheAsiaJakartaLocation
}
