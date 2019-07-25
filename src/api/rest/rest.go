package rest

// REST API Resources.

import (
	"net/http"

	chanstore "github.com/arush15june/revshellz/src/pkg/chanstore"
	utils "github.com/arush15june/revshellz/src/pkg/utils"
)

type GetChansResourcePayload struct {
	Connections []string `json:"connections"`
}

// GetChansResource returns a json list of all active connections.
func GetChansResource(w http.ResponseWriter, r *http.Request) {
	chans := chanstore.GetChans()
	conns := &GetChansResourcePayload{}

	for k := range chans {
		conns.Connections = append(conns.Connections, k)
	}

	utils.WriteSerializeJSON(w, conns)
}
