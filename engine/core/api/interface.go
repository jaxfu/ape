package api

import (
	"context"

	"github.com/jaxfu/ape/engine/core/api/internal"
)

type ApiInterface interface {
	StartServer(ctx context.Context) error
}

func NewServer(url string, port uint, clientDirFp string) (ApiInterface, error) {
	return internal.DefaultServer(url, port, clientDirFp)
}
