package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriverOk(t *testing.T) {
	car := NewCarProxy(&Driver{25})
	assert.Equal(t, car.Drive(), successDriven, "fungsi salah")
}

func TestDriverToYoung(t *testing.T) {
	car := NewCarProxy(&Driver{15})
	assert.Equal(t, car.Drive(), failDriven)
}
