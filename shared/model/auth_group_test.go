package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthGroupValidation(t *testing.T) {
	ag := RandomAuthGroup()
	assert.NoError(t, ag.Validate())

	ag.ID = ""
	assert.Error(t, ag.Validate())

	ag = RandomAuthGroup()
	ag.Users = []string{""}
	assert.Error(t, ag.Validate())
}

func TestAuthGroupCreateValidation(t *testing.T) {
	agc := RandomAuthGroupCreate()
	assert.NoError(t, agc.Validate())

	agc.Name = ""
	assert.Error(t, agc.Validate())

	agc = RandomAuthGroupCreate()
	agc.Users = []string{""}
	assert.Error(t, agc.Validate())
}

func TestAuthGroupUpdateValidation(t *testing.T) {
	agu := RandomAuthGroupUpdate()
	assert.NoError(t, agu.Validate())

	agu.Users = []string{""}
	assert.Error(t, agu.Validate())
}

func TestRandomAuthGroupGeneration(t *testing.T) {
	ag := RandomAuthGroup()
	assert.NotEmpty(t, ag.ID)
	assert.NotEmpty(t, ag.Name)
	assert.NotEmpty(t, ag.Users)
}
