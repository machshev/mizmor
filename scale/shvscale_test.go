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
		want  uint8
	}{
		{"min value", 0, 10},
		{"max value", 0, 20},
		{"out of range", 20, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, s.fixed(tt.input), tt.want, "note incorrect")
		})
	}
	s.fixed(0)
}
