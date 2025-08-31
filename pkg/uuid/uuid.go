package uuid

import (
	uuid "github.com/google/uuid"
	puuid "github.com/google/uuid"
)

func NewUUID() uuid.UUID {
	return puuid.MustParse(uuid.NewString())
}

func MustParse(s string) uuid.UUID {
	return puuid.MustParse(s)
}