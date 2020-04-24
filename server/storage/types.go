package storage

import "github.com/hunter1085/notary/tuf/data"

// MetaUpdate packages up the fields required to update a TUF record
type MetaUpdate struct {
	Role    data.RoleName
	Version int
	Data    []byte
}
