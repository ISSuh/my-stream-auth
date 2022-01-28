package internal

type StorageType int

const (
	Local    = StorageType(0)
	DataBase = StorageType(1)
)

type Configure struct {
	storageType StorageType
}
