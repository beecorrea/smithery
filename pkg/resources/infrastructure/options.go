package infrastructure

type InfraOptions struct {
	Tier     int64
	Provider Provider
}

type StorageType string

const (
	Blob     StorageType = "BLOB"
	Database StorageType = "DATABASE"
)
