package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestValidator(t *testing.T) {
	ts := testStruct{}
	err := c.Validator.Validate(ts)
	assert.Error(t, err)

	ts.Name = "Professor Hubert J. Farnsworth"
	ts.Email = "hjf@email"
	assert.Error(t, c.Validator.Validate(ts))

	ts.Email = "hjf@planetexpress.com"
	assert.NoError(t, c.Validator.Validate(ts))
}
