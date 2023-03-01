package server

import (
	"context"
)

// List information about the server.
func (s *Client) List(ctx context.Context) (response ListResponse, err error) {
	// u := fmt.Sprintf("server/get_server.json?name=%v", request.ServerName)
	u := "server/list_servers.json"
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return response, err
	}

	if err := s.client.Do(ctx, req, &response); err != nil {
		return response, err
	}

	return response, nil
}
