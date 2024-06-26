package probeservices

//
// bouncer.go - GET /api/v1/test-helpers
//

import (
	"context"

	"github.com/ooni/probe-cli/v3/internal/httpclientx"
	"github.com/ooni/probe-cli/v3/internal/model"
	"github.com/ooni/probe-cli/v3/internal/urlx"
)

// GetTestHelpers queries the /api/v1/test-helpers API.
func (c *Client) GetTestHelpers(ctx context.Context) (map[string][]model.OOAPIService, error) {
	// construct the URL to use
	URL, err := urlx.ResolveReference(c.BaseURL, "/api/v1/test-helpers", "")
	if err != nil {
		return nil, err
	}

	// get the response
	return httpclientx.GetJSON[map[string][]model.OOAPIService](ctx, URL, &httpclientx.Config{
		Client:    c.HTTPClient,
		Host:      c.Host,
		Logger:    c.Logger,
		UserAgent: c.UserAgent,
	})
}
