package simple

import "time"

func myFunction() {
	type myType struct {
		version   int
		time.Time // want `embedded types should be listed before non embedded types`
	}
}

type ValidStructWithSingleLineComments struct {
	// time.Time Single line comment
	time.Time

	// version Single line comment
	version int
}

type StructWithSingleLineComments struct {
	// time.Time Single line comment
	time.Time // want `there must be an empty line separating embedded fields from regular fields`
	// version Single line comment
	version int
}

type StructWithMultiLineComments struct {
	// time.Time Single line comment
	time.Time // want `there must be an empty line separating embedded fields from regular fields`
	// version Single line comment
	// very long comment
	version int
}
