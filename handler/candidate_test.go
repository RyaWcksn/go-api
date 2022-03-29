package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCandidate(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		phone   string
		cv      string
		details string
	}{
		{
			name:    "test",
			email:   "Pramudya",
			phone:   "081234",
			cv:      "Pramudya",
			details: "Pramudya",
		},
		{
			name:    "test",
			email:   "Pramudya",
			phone:   "",
			cv:      "Pramudya",
			details: "Pramudya",
		},
		{
			name:    "",
			email:   "Pramudya",
			phone:   "081234",
			cv:      "Pramudya",
			details: "Pramudya",
		},
	}

	for _, v := range tests {
		assert.NotEmpty(t, v.name, "name should not be empty")
		assert.NotEmpty(t, v.email, "email should not be empty")
		assert.NotEmpty(t, v.phone, "name should not be empty")
	}
}
