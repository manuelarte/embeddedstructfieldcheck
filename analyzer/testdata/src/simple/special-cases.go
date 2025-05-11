package simple

import "time"

func myFunction() {
	type myType struct {
		version   int
		time.Time // want `embedded fields should be listed before regular fields`
	}
}

type DeclaredSameLineNoSpace struct {
	time.Time // want `there must be an empty line separating embedded fields from regular fields`
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
	Bar  // want `there must be an empty line separating embedded fields from regular fields`
	Data []string
}

type FooBar struct {
	Foo
	Data []string
	Bar  // want `embedded fields should be listed before regular fields`
}
