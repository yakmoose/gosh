package stack

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/utils"
)

// Restart restarts a stack on the given server.
func (s *Client) Restart(ctx context.Context, request StopStartRestartRequest) (response StartStopRestartResponse, err error) {
	uri := "cloud/stack/restart.json"
	if err != nil {
	return s.stopStartRestart(ctx, request, uri)
}
