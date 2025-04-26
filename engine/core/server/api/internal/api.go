package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/events"
)

func DefaultApi(bus *events.Bus) *Api {
	return &Api{
		Bus: bus,
	}
}

type Api struct {
	Bus *events.Bus
}

func (a *Api) CreateComponent(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into the struct
	var req components.Object
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Invalid request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// dev.PrettyPrint(req)

	a.Bus.Events <- events.Event{
		EventType: events.EventTypes.CREATE,
		Component: req,
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
