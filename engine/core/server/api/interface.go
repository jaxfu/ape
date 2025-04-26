package api

import (
	"net/http"

	"github.com/jaxfu/ape/engine/core/events"
	"github.com/jaxfu/ape/engine/core/server/api/internal"
)

type Api interface {
	CreateComponent(w http.ResponseWriter, r *http.Request)
}

func NewApi(bus *events.Bus) Api {
	return internal.DefaultApi(bus)
}
