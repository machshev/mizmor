package scale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPsalmsScaleFixed(t *testing.T) {
	s := NewPsalmScaleSHV()
	var tests = []struct {
		name  string
		input uint8
		want  Note
	}{
		{"note 1", 0, 75},
		{"note 2", 1, 76},
		{"note 3", 2, 78},
		{"note 4", 3, 79},
		{"note 5", 4, 81},
		{"note 6", 5, 83},
		{"note 7", 6, 84},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := s.Fixed(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, note, tt.want, "note incorrect")
		})
	}
}
