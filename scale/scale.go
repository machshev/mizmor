// Package scale implements the API for interacting with different scales.
//
// In the method proposed by Susan Heik-Vantoura there are two scales used in
// the tanak; the common scale and then another one used for the psalms and a
// few other specific passages.
package scale

// Note in the scale
type Note uint8

// Offeset sequence providing relative note changes to go from a fixed note to
// final note
type Offsets []int8

// Scale provides access to a Fixed note scale and orenamental relative notes.
//
// Such a scale would be used for example to represent music for chironomy
// where one hand gestures the fixed notes and the other hand gestures the
// relative notes.
type Scale interface {
	// get a note by id from the fixed note scale
	Fixed(id uint8) (Note, error)

	// Given a fixed note ID (fixedId) get a relative note by ID (relId)
	Relative(fixedId uint8, relId uint8) ([]Note, error)
}
