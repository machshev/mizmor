package scale

import "fmt"

type scaleSHV struct {
	notes        []Note
	startNote    uint8
	baseScaleLen uint8
	ornaments    []Offsets
}

// Question: using this scale, notes + ornamental offsets the range of notes is
// 12. However Psa 92:4, Psa 33:2 and Psa 144:9 all refereing to 10 strings. So
// how many notes of this scale are used in the psalms? If more then maybe the
// ornaments are incorrect and should be modified to keep within the playable
// range?
//
// David invented instruments for the priests to use in the temple (ref?). So
// were these extended range to fit a 12 string scale? But then why refer to 10
// strings in the psalms themselves?

func NewPsalmScaleSHV() scaleSHV {
	scale := scaleSHV{
		notes: []Note{
			71, // B4
			72, // C5
			75, // D5#
			76, // E5
			78, // F5#
			79, // G5
			81, // A5
			83, // B5
			84, // C6
			87, // D6#
			88, // E6
			90, // F6
			91, // G6
		},
		startNote:    2,
		baseScaleLen: 7,
		ornaments: []Offsets{
			{1},          // pashta
			{2},          // geresh
			{-1},         // revia
			{4},          // illuy
			{2, 1},       // pazer
			{3, 0},       // ole veyored
			{-1, 1},      // tsinnor
			{-2, -1, -1}, // shalshelet
		},
	}
	return scale
}

func (s scaleSHV) Fixed(id uint8) (Note, error) {
	if id >= s.baseScaleLen {
		return 0, fmt.Errorf(
			"index '%d' is out of base scale range [0-%d]",
			id,
			s.baseScaleLen-1,
		)
	}
	return s.notes[s.startNote+id], nil
}

func (s scaleSHV) Relative(fixedId uint8, relId uint8) ([]Note, error) {
	relNotes := []Note{}

	if fixedId >= s.baseScaleLen {
		return []Note{}, fmt.Errorf(
			"index '%d' is out of base scale range [0-%d]",
			fixedId,
			s.baseScaleLen-1,
		)
	}
	baseId := uint8(s.startNote + fixedId)

	if relId >= uint8(len(s.ornaments)) {
		return []Note{}, fmt.Errorf(
			"index '%d' is out of ornament range [0-%d]",
			relId, len(s.ornaments),
		)
	}

	for i := range s.ornaments {
		relNotes = append(relNotes, s.notes[baseId+uint8(i)])
	}

	return relNotes, nil
}
