// This file was auto-generated by Fern from our API Definition.

package bankinfo

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	hris "github.com/merge-api/merge-go-client/hris"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	List(ctx context.Context, request *hris.BankInfoListRequest) (*hris.PaginatedBankInfoList, error)
	Retrieve(ctx context.Context, id string, request *hris.BankInfoRetrieveRequest) (*hris.BankInfo, error)
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

// Returns a list of `BankInfo` objects.
func (c *client) List(ctx context.Context, request *hris.BankInfoListRequest) (*hris.PaginatedBankInfoList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bank-info"

	queryParams := make(url.Values)
	if request.AccountType != nil {
		queryParams.Add("account_type", fmt.Sprintf("%v", *request.AccountType))
	}
	if request.BankName != nil {
		queryParams.Add("bank_name", fmt.Sprintf("%v", *request.BankName))
	}
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EmployeeId != nil {
		queryParams.Add("employee_id", fmt.Sprintf("%v", *request.EmployeeId))
	}
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.OrderBy != nil {
		queryParams.Add("order_by", fmt.Sprintf("%v", *request.OrderBy))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", *request.RemoteFields))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", *request.ShowEnumOrigins))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hris.PaginatedBankInfoList
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

// Returns a `BankInfo` object with the given `id`.
func (c *client) Retrieve(ctx context.Context, id string, request *hris.BankInfoRetrieveRequest) (*hris.BankInfo, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bank-info/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.RemoteFields != nil {
		queryParams.Add("remote_fields", fmt.Sprintf("%v", *request.RemoteFields))
	}
	if request.ShowEnumOrigins != nil {
		queryParams.Add("show_enum_origins", fmt.Sprintf("%v", *request.ShowEnumOrigins))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *hris.BankInfo
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
