package simple

import "time"

type ValidStruct struct {
	time.Time

	version int
}

type NoSpaceStruct struct {
	time.Time // want `missing space`
	version   int
}

type NotSortedStruct struct {
	version int

	time.Time // want `should be declared before 17`
}
