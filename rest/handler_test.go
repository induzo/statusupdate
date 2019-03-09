package rest

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
)

func TestStatusUpdateHandler(t *testing.T) {

	tests := []struct {
		name                  string
		withPayloadError      bool
		withStatusUpdateError bool
		withBadID             bool
		withBadAction         bool
		wantedStatus          int
	}{
		{
			name:         "working StatusUpdate",
			wantedStatus: http.StatusNoContent,
		},
		{
			name:             "bad payload, non working StatusUpdate",
			withPayloadError: true,
			wantedStatus:     http.StatusBadRequest,
		},
		{
			name:         "bad id, non working StatusUpdate",
			withBadID:    true,
			wantedStatus: http.StatusBadRequest,
		},
		{
			name:         "bad action, non working StatusUpdate",
			withBadID:    true,
			wantedStatus: http.StatusBadRequest,
		},
		{
			name:                  "internal error, non working StatusUpdate",
			withStatusUpdateError: true,
			wantedStatus:          http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m := newMgrMock()
			id := xid.New()
			if tt.withBadID {
				id = xid.ID{}
			}
			actionID := "10"
			if tt.withBadAction {
				actionID = "bad"
			}
			payload := []byte(`{"well": 2}`)
			if tt.withPayloadError {
				payload = payload[1:]
			}
			m.wantStatusUpdateError = tt.withStatusUpdateError

			rr := httptest.NewRecorder()
			req := httptest.NewRequest(
				"POST", `http://dummy/entity`,
				bytes.NewBuffer(payload),
			)
			// Set the URL param
			ctxR := chi.NewRouteContext()
			ctxR.URLParams.Add("ID", id.String())
			ctxR.URLParams.Add("ActionID", actionID)
			newCtx := context.WithValue(ctx, chi.RouteCtxKey, ctxR)
			req = req.WithContext(newCtx)

			StatusUpdateHandler(m)(rr, req)
			resp := rr.Result()
			defer resp.Body.Close()

			if status := resp.StatusCode; status != tt.wantedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantedStatus)
				return
			}
		})
	}
}

func BenchmarkPATCHHandler(b *testing.B) {
	ctx := context.Background()
	m := newMgrMock()
	payload := []byte(`{"status_id": 2}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("PATCH", `http://dummy/ent`,
			bytes.NewBuffer(payload))
		ctxR := chi.NewRouteContext()
		ctxR.URLParams.Add("ID", xid.New().String())
		ctxR.URLParams.Add("ActionID", "123")
		newCtx := context.WithValue(ctx, chi.RouteCtxKey, ctxR)
		req = req.WithContext(newCtx)
		b.StartTimer()
		StatusUpdateHandler(m)(rr, req)
	}
}
