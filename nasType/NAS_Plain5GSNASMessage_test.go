package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nycu-ucr/nas/nasType"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	assert.NotNil(t, a)
}
