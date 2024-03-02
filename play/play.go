package play

import (
	"fmt"

	"github.com/coral/fluidsynth2"
)

func Play(filename string) error {
	s := fluidsynth2.NewSettings()

	synth := fluidsynth2.NewSynth(s)
	if _, err := synth.SFLoad("SuperHarp.sf2", false); err != nil {
		return fmt.Errorf("failed to play midi: %w", err)
	}

	player := fluidsynth2.NewPlayer(synth)

	if err := player.Add(filename + ".mid"); err != nil {
		return fmt.Errorf("failed to add '%s.mid' to player: %w", filename, err)
	}

	fluidsynth2.NewAudioDriver(s, synth)

	if err := player.Play(); err != nil {
		return fmt.Errorf("error playing '%s.mid': %w", filename, err)
	}

	player.Join()

	return nil
}
