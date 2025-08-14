// https://github.com/manuelarte/embeddedstructfieldcheck/issues/21
package block_comments

import "time"

type BlockComments struct {
	time.Time

	/**
	 * Other important fields
	 */

	// the version
	version int

	// other field
	name string
}

type MoreComments struct {
	time.Time

	// some fields

	// the version
	version int

	// other field
	name string

	// some other fields

	// myfield
	age int
}
