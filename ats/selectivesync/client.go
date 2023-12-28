// This file was auto-generated by Fern from our API Definition.

package selectivesync

import (
	context "context"
	fmt "fmt"
	ats "github.com/merge-api/merge-go-client/ats"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
	url "net/url"
)

type Client interface {
	ConfigurationsList(ctx context.Context) ([]*ats.LinkedAccountSelectiveSyncConfiguration, error)
	ConfigurationsUpdate(ctx context.Context, request *ats.LinkedAccountSelectiveSyncConfigurationListRequest) ([]*ats.LinkedAccountSelectiveSyncConfiguration, error)
	MetaList(ctx context.Context, request *ats.SelectiveSyncMetaListRequest) (*ats.PaginatedConditionSchemaList, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Get a linked account's selective syncs.
func (c *client) ConfigurationsList(ctx context.Context) ([]*ats.LinkedAccountSelectiveSyncConfiguration, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "selective-sync/configurations"

	var response []*ats.LinkedAccountSelectiveSyncConfiguration
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Replace a linked account's selective syncs.
func (c *client) ConfigurationsUpdate(ctx context.Context, request *ats.LinkedAccountSelectiveSyncConfigurationListRequest) ([]*ats.LinkedAccountSelectiveSyncConfiguration, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "selective-sync/configurations"

	var response []*ats.LinkedAccountSelectiveSyncConfiguration
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Get metadata for the conditions available to a linked account.
func (c *client) MetaList(ctx context.Context, request *ats.SelectiveSyncMetaListRequest) (*ats.PaginatedConditionSchemaList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "selective-sync/meta"

	queryParams := make(url.Values)
	if request.CommonModel != nil {
		queryParams.Add("common_model", fmt.Sprintf("%v", *request.CommonModel))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ats.PaginatedConditionSchemaList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
