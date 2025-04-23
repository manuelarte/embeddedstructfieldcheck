package simple

import "time"

type NoSpaceStruct struct {
	time.Time // want `missing space`
	version   int
}
