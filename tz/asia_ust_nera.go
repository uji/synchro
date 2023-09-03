// Code generated by tzgen. DO NOT EDIT.

package tz

import "time"
import "sync"

var (
	onceAsiaUstNeraLocation  sync.Once
	cacheAsiaUstNeraLocation *time.Location
)

type AsiaUstNera struct{}

func (AsiaUstNera) Location() *time.Location {
	onceAsiaUstNeraLocation.Do(func() {
		loc, err := time.LoadLocation("Asia/Ust-Nera")
		if err != nil {
			panic(err)
		}
		cacheAsiaUstNeraLocation = loc
	})
	return cacheAsiaUstNeraLocation
}
