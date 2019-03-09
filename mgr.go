package statusupdate

import (
	"context"
	"io"

	"github.com/induzo/fsm"
	"github.com/rs/xid"
	"github.com/vincentserpoul/gohttperror"
)

// MgrI is the interface to initialize the new entity mgr
type MgrI interface {
	StatusUpdate(
		context.Context,
		xid.ID,
		fsm.Action,
		io.Reader,
	) error
	IsActionAuthorized(
		ctx context.Context,
		subject string,
		action fsm.Action,
	) error
	GetOutcomeStatus(status fsm.Status, action fsm.Action) (fsm.Status, error)
	MapErrorToHTTPError(error) *gohttperror.ErrResponse
}
