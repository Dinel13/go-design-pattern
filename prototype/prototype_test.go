package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyPerson(t *testing.T) {
	aku := person{"huddin", &address{"jl nangka", "gowa"}}

	dia := aku

	assert.Equal(t, aku.name, dia.name, "tidak oke")
}
