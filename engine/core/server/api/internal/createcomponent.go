package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/events"
)

func (a *Api) CreateComponent(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	var req components.Object
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Invalid request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// dev.PrettyPrint(req)

	event := events.Event{
		EventType: events.EventTypes.CREATE_COMPONENT,
		Component: req,
	}

	result := a.ProcessEvent(event)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error processing event: %+v", result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
