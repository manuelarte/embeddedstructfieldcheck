package simple

import "time"

func myFunction() {
	type myType struct {
		version   int
		time.Time // want `embedded fields should be listed before regular fields`
	}
}

type DeclaredSameLineNoSpace struct {
	time.Time
	lat, long int
}

type DeclaredSameLineEmbeddedLast struct {
	lat, long int
	time.Time // want `embedded fields should be listed before regular fields`
}

type Foo struct {
	Name string
}

type Bar struct {
	Value int
}

type MultipleEmbeddedTypes struct {
	Foo
	Bar
	Data []string
}
