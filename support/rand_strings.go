package support

import (
	"fmt"
	"math/rand"
	"time"
)

var seeded bool

func CheapRandString8() string {
	if !seeded {
		rand.Seed(time.Now().UnixNano())
		seeded = true
	}
	return fmt.Sprintf("%x", rand.Uint32())
}
func CheapRandString16() string {
	if !seeded {
		rand.Seed(time.Now().UnixNano())
		seeded = true
	}
	return fmt.Sprintf("%x%x", rand.Uint32(), rand.Uint32())
}
