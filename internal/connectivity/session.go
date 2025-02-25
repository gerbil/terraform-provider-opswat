package opswatClient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetSession - Returns session config
func (c *Client) GetSession(ctx context.Context) (*Session, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/session", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/session", c.HostURL))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateSession - Updates session config
func (c *Client) UpdateSession(ctx context.Context, config Session) (*Session, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/session", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/session, request body: %s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateSession - Creates session config
func (c *Client) CreateSession(ctx context.Context, config Session) (*Session, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/session", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/session, request body: %s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
