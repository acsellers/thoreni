package router

import (
	"github.com/acsellers/thoreni"
)

type RoutingFunc func(*thoreni.Contextable)
