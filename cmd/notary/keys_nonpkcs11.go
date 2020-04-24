// +build !pkcs11

package main

import (
	"errors"

	"github.com/hunter1085/notary"
	store "github.com/hunter1085/notary/storage"
	"github.com/hunter1085/notary/trustmanager"
)

func getYubiStore(fileKeyStore trustmanager.KeyStore, ret notary.PassRetriever) (trustmanager.KeyStore, error) {
	return nil, errors.New("Not built with hardware support")
}

func getImporters(baseDir string, _ notary.PassRetriever) ([]trustmanager.Importer, error) {
	fileStore, err := store.NewPrivateKeyFileStorage(baseDir, notary.KeyExtension)
	if err != nil {
		return nil, err
	}
	return []trustmanager.Importer{fileStore}, nil
}
