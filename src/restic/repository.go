package restic

import "restic/crypto"

// Repository stores data in a backend. It provides high-level functions and
// transparently encrypts/decrypts data.
type Repository interface {

	// Backend returns the backend used by the repository
	Backend() Backend

	Key() *crypto.Key

	SetIndex(Index)

	Index() Index
	SaveFullIndex() error
	SaveIndex() error

	Config() Config

	LookupBlobSize(ID, BlobType) (uint, error)

	List(FileType, <-chan struct{}) <-chan ID
	ListPack(ID) ([]Blob, int64, error)

	Flush() error

	SaveJSON(BlobType, interface{}) (ID, error)
	SaveUnpacked(FileType, []byte) (ID, error)
	SaveAndEncrypt(BlobType, []byte, *ID) (ID, error)
	SaveJSONUnpacked(FileType, interface{}) (ID, error)

	LoadJSONPack(BlobType, ID, interface{}) error
	LoadJSONUnpacked(FileType, ID, interface{}) error
	LoadBlob(ID, BlobType, []byte) ([]byte, error)
	LoadAndDecrypt(FileType, ID) ([]byte, error)
}

// Deleter removes all data stored in a backend/repo.
type Deleter interface {
	Delete() error
}

// Lister allows listing files in a backend.
type Lister interface {
	List(FileType, <-chan struct{}) <-chan string
}

// Index keeps track of the blobs are stored within files.
type Index interface {
	Has(ID, BlobType) bool
	Lookup(ID, BlobType) ([]PackedBlob, error)
	Count(BlobType) uint
}
