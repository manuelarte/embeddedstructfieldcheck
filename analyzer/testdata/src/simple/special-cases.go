package simple

import "time"

func myFunction() {
	type myType struct {
		version   int
		time.Time // want `should be declared before line 7`
	}
}
