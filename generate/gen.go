package generate

import (
	"fmt"

	"github.com/machshev/mizmor/scale"
	midi "gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

type Generator struct {
	scale scale.Scale
}

func NewGenerator(s scale.Scale) *Generator {
	return &Generator{
		scale: s,
	}
}

func (g Generator) GenMidi(filename string) error {
	var (
		clock          = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also a common choice
		general, piano smf.Track             // our tracks
	)

	// first track must have tempo and meter information
	general.Add(0, smf.MetaTrackSequenceName("general"))
	general.Add(0, smf.MetaMeter(3, 4))
	general.Add(0, smf.MetaTempo(140))
	general.Add(clock.Ticks4th()*6, smf.MetaTempo(130))
	general.Add(clock.Ticks4th(), smf.MetaTempo(135))
	general.Close(0) // don't forget to close a track

	piano.Add(0, smf.MetaInstrument("harp"))
	piano.Add(0, midi.ProgramChange(0, gm.Instr_OrchestralHarp.Value()))

	sequence := []uint8{0, 1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1, 0}

	for _, id := range sequence {
		note, err := g.scale.Fixed(id)
		if err != nil {
			return fmt.Errorf(
				"can't determine fixed scale note id '%d': %w", id, err,
			)
		}
		piano.Add(0, midi.NoteOn(0, uint8(note), 120))
		piano.Add(clock.Ticks4th(), midi.NoteOff(0, uint8(note)))
	}

	piano.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(general)
	s.Add(piano)

	err := s.WriteFile(filename + ".mid")
	if err != nil {
		return fmt.Errorf("failed to save midi file: %w", err)
	}

	return nil
}
