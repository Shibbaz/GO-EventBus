package helpers

import (
	. "events"
)

var BatchSize = 100
var ProcessNum = 1000

type Bus chan []Event
