package rest

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/induzo/fsm"
	"github.com/rs/xid"
	"github.com/induzo/gohttperror"
)

var (
	// ErrNotFound is used for Get and GetList when no entity is found
	ErrNotFound = errors.New("entity not found")
	// ErrBadRequest is used when the request is malformed or wrong
	ErrBadRequest = errors.New("bad request")
)

// mgrMock is a mock for the mgr interface
type mgrMock struct {
	wantStatusUpdateError bool
	EntityList            map[xid.ID]*entityMock
}

// entityMock is respecting the CRUD interface
type entityMock struct {
	ID       xid.ID `json:"id"`
	StatusID string `json:"status_id"`
}

func newMgrMock() *mgrMock {
	return &mgrMock{}
}

func (m *mgrMock) StatusUpdate(
	ctx context.Context,
	id xid.ID,
	actionID fsm.Action,
	payload io.Reader,
) error {
	if m.wantStatusUpdateError {
		return fmt.Errorf("Error create")
	}
	return nil
}

func (m *mgrMock) MapErrorToHTTPError(e error) *gohttperror.ErrResponse {
	switch e {
	case ErrNotFound:
		return gohttperror.ErrNotFound
	case ErrBadRequest:
		return gohttperror.ErrBadRequest(e)
	default:
		return gohttperror.ErrInternal(e)
	}
}

func (m *mgrMock) IsActionAuthorized(
	context.Context,
	string,
	fsm.Action,
) error {
	return nil
}

func (m *mgrMock) GetOutcomeStatus(
	fsm.Status,
	fsm.Action,
) (fsm.Status, error) {
	return "ok", nil
}
