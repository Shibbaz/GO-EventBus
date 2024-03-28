package helpers

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/events"
)

var BatchSize = 100
var ProcessNum = 1000

type Bus chan []Event
