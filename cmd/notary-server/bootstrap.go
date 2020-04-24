package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/hunter1085/notary"
	"github.com/hunter1085/notary/storage"
)

func bootstrap(ctx context.Context) error {
	s := ctx.Value(notary.CtxKeyMetaStore)
	if s == nil {
		return fmt.Errorf("no store set during bootstrapping")
	}
	store, ok := s.(storage.Bootstrapper)
	if !ok {
		return fmt.Errorf("store does not support bootstrapping")
	}
	return store.Bootstrap()
}
