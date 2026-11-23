package utils

import (
	"github.com/google/uuid"
)

func UuidGenerator() string {
	machineID := uuid.NewString()
	return machineID
}
