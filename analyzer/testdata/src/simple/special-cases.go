package simple

import "time"

func myFunction() {
	type myType struct {
		version   int
		time.Time // want `embedded types should be listed before non embedded types`
	}
}

type DeclaredSameLineNoSpace struct {
	time.Time // want `there must be an empty line separating embedded fields from regular fields`
	lat, long int
}

type DeclaredSameLineEmbeddedLast struct {
	lat, long int
	time.Time // want `embedded types should be listed before non embedded types`
}
