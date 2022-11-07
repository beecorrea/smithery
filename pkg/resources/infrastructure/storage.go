package infrastructure

import (
	"github.com/google/uuid"
)

type StorageInfra struct {
	Id      uuid.UUID
	Type    StorageType
	Options InfraOptions
}
