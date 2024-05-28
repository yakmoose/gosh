package user

import (
	"context"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/utils"
	"net/url"
)

// Add creates a new SSH user.
func (s *Client) Add(ctx context.Context, request AddRequest) (response UpdateResponse, err error) {
	uri := "cloud/ssh/user/add.json"
	keys := []string{
		"client_id",
		"server_name",
		"username",
		"params[password]",
		"params[containers][]",
		"params[volumes][]",
		"params[ssh_keys][]",
		"params[read_only_config]",
	}

	values := url.Values{}
	values.Add("client_id", s.client.ClientID)
	values.Add("server_name", request.ServerName)
	values.Add("username", request.Username)
	values.Add("params[password]", request.Password)
	values.Add("params[read_only_config]", fmt.Sprint(utils.BoolToInt(request.ReadOnlyConfig)))

	for _, c := range request.Containers {
		values.Add("params[containers][]", c)
	}

	for _, k := range request.SSHKeys {
		values.Add("params[ssh_keys][]", k)
	}

	for _, k := range request.Volumes {
		values.Add("params[volumes][]", k)
	}

	req, err := s.client.NewRequest("POST", uri, utils.Encode(values, keys))
	if err != nil {
		return response, err
	}

	if err := s.client.Do(ctx, req, &response); err != nil {
		return response, err
	}

	return response, nil
}
