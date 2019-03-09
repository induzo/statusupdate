package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/induzo/fsm"
	"github.com/induzo/statusupdate"
	"github.com/rs/xid"
	"github.com/vincentserpoul/gohttperror"
)

// StatusUpdateHandler allows update of status for the entity
func StatusUpdateHandler(
	sumgr statusupdate.MgrI,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var errRender error
		defer func() {
			if errRender != nil {
				log.Fatalf("StatusUpdateHandler Render: %v", errRender)
			}
		}()

		id, errI := xid.FromString(chi.URLParam(r, "ID"))
		if id.IsNil() || errI != nil {
			errRender = render.Render(w, r,
				gohttperror.ErrBadRequest(errI),
			)
			return
		}

		actionID := fsm.Action(chi.URLParam(r, "ActionID"))

		var payload bytes.Buffer
		pl := io.TeeReader(r.Body, &payload)
		errJSON := json.NewDecoder(pl).Decode(&pl)
		if errJSON != nil && errJSON != io.EOF {
			errRender = render.Render(w, r,
				gohttperror.ErrBadRequest(
					fmt.Errorf("error decoding payload: %v", errJSON),
				),
			)
			return
		}
		if errJSON == io.EOF {
			payload = *bytes.NewBufferString(`{}`)
		}

		if err := sumgr.StatusUpdate(
			r.Context(),
			id,
			actionID,
			&payload,
		); err != nil {
			errRender = render.Render(w, r, sumgr.MapErrorToHTTPError(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
