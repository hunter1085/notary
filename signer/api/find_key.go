package api

import (
	"github.com/hunter1085/notary/signer"
	"github.com/hunter1085/notary/trustmanager"
	"github.com/hunter1085/notary/tuf/data"

	pb "github.com/hunter1085/notary/proto"
)

// findKeyByID looks for the key with the given ID in each of the
// signing services in sigServices. It returns the first matching key it finds,
// or ErrInvalidKeyID if the key is not found in any of the signing services.
func findKeyByID(cryptoServices signer.CryptoServiceIndex, keyID *pb.KeyID) (data.PrivateKey, data.RoleName, error) {
	for _, service := range cryptoServices {
		key, role, err := service.GetPrivateKey(keyID.ID)
		if err == nil {
			return key, role, nil
		}
	}

	return nil, "", trustmanager.ErrKeyNotFound{KeyID: keyID.ID}
}
