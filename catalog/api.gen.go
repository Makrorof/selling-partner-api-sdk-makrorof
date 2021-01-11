// Package catalog provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"gopkg.me/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListCatalogCategories request
	ListCatalogCategories(ctx context.Context, params *ListCatalogCategoriesParams) (*http.Response, error)

	// ListCatalogItems request
	ListCatalogItems(ctx context.Context, params *ListCatalogItemsParams) (*http.Response, error)

	// GetCatalogItem request
	GetCatalogItem(ctx context.Context, asin string, params *GetCatalogItemParams) (*http.Response, error)
}

func (c *Client) ListCatalogCategories(ctx context.Context, params *ListCatalogCategoriesParams) (*http.Response, error) {
	req, err := NewListCatalogCategoriesRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) ListCatalogItems(ctx context.Context, params *ListCatalogItemsParams) (*http.Response, error) {
	req, err := NewListCatalogItemsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetCatalogItem(ctx context.Context, asin string, params *GetCatalogItemParams) (*http.Response, error) {
	req, err := NewGetCatalogItemRequest(c.Endpoint, asin, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewListCatalogCategoriesRequest generates requests for ListCatalogCategories
func NewListCatalogCategoriesRequest(endpoint string, params *ListCatalogCategoriesParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/catalog/v0/categories")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.ASIN != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "ASIN", *params.ASIN); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SellerSKU != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "SellerSKU", *params.SellerSKU); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListCatalogItemsRequest generates requests for ListCatalogItems
func NewListCatalogItemsRequest(endpoint string, params *ListCatalogItemsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/catalog/v0/items")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Query != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "Query", *params.Query); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.QueryContextId != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "QueryContextId", *params.QueryContextId); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.SellerSKU != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "SellerSKU", *params.SellerSKU); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.UPC != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "UPC", *params.UPC); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.EAN != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "EAN", *params.EAN); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.ISBN != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "ISBN", *params.ISBN); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.JAN != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "JAN", *params.JAN); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetCatalogItemRequest generates requests for GetCatalogItem
func NewGetCatalogItemRequest(endpoint string, asin string, params *GetCatalogItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "asin", asin)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/catalog/v0/items/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListCatalogCategories request
	ListCatalogCategoriesWithResponse(ctx context.Context, params *ListCatalogCategoriesParams) (*ListCatalogCategoriesResp, error)

	// ListCatalogItems request
	ListCatalogItemsWithResponse(ctx context.Context, params *ListCatalogItemsParams) (*ListCatalogItemsResp, error)

	// GetCatalogItem request
	GetCatalogItemWithResponse(ctx context.Context, asin string, params *GetCatalogItemParams) (*GetCatalogItemResp, error)
}

type ListCatalogCategoriesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *ListCatalogCategoriesResponse
}

// Status returns HTTPResponse.Status
func (r ListCatalogCategoriesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListCatalogCategoriesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListCatalogItemsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *ListCatalogItemsResponse
}

// Status returns HTTPResponse.Status
func (r ListCatalogItemsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListCatalogItemsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCatalogItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetCatalogItemResponse
}

// Status returns HTTPResponse.Status
func (r GetCatalogItemResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCatalogItemResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListCatalogCategoriesWithResponse request returning *ListCatalogCategoriesResponse
func (c *ClientWithResponses) ListCatalogCategoriesWithResponse(ctx context.Context, params *ListCatalogCategoriesParams) (*ListCatalogCategoriesResp, error) {
	rsp, err := c.ListCatalogCategories(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseListCatalogCategoriesResp(rsp)
}

// ListCatalogItemsWithResponse request returning *ListCatalogItemsResponse
func (c *ClientWithResponses) ListCatalogItemsWithResponse(ctx context.Context, params *ListCatalogItemsParams) (*ListCatalogItemsResp, error) {
	rsp, err := c.ListCatalogItems(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseListCatalogItemsResp(rsp)
}

// GetCatalogItemWithResponse request returning *GetCatalogItemResponse
func (c *ClientWithResponses) GetCatalogItemWithResponse(ctx context.Context, asin string, params *GetCatalogItemParams) (*GetCatalogItemResp, error) {
	rsp, err := c.GetCatalogItem(ctx, asin, params)
	if err != nil {
		return nil, err
	}
	return ParseGetCatalogItemResp(rsp)
}

// ParseListCatalogCategoriesResp parses an HTTP response from a ListCatalogCategoriesWithResponse call
func ParseListCatalogCategoriesResp(rsp *http.Response) (*ListCatalogCategoriesResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListCatalogCategoriesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest ListCatalogCategoriesResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseListCatalogItemsResp parses an HTTP response from a ListCatalogItemsWithResponse call
func ParseListCatalogItemsResp(rsp *http.Response) (*ListCatalogItemsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListCatalogItemsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest ListCatalogItemsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetCatalogItemResp parses an HTTP response from a GetCatalogItemWithResponse call
func ParseGetCatalogItemResp(rsp *http.Response) (*GetCatalogItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetCatalogItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetCatalogItemResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
