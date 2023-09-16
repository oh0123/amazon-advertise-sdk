// Package dspaudiences provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package dspaudiences

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for DspAudienceCreateRequestItemAudienceType.
const (
	PRODUCTPURCHASES          DspAudienceCreateRequestItemAudienceType = "PRODUCT_PURCHASES"
	PRODUCTSEARCH             DspAudienceCreateRequestItemAudienceType = "PRODUCT_SEARCH"
	PRODUCTSIMS               DspAudienceCreateRequestItemAudienceType = "PRODUCT_SIMS"
	PRODUCTVIEWS              DspAudienceCreateRequestItemAudienceType = "PRODUCT_VIEWS"
	WHOLEFOODSMARKETPURCHASES DspAudienceCreateRequestItemAudienceType = "WHOLE_FOODS_MARKET_PURCHASES"
)

// Defines values for DspAudienceErrorItemErrorErrorType.
const (
	OTHER           DspAudienceErrorItemErrorErrorType = "OTHER"
	VALUEINVALID    DspAudienceErrorItemErrorErrorType = "VALUE_INVALID"
	VALUEOUTOFRANGE DspAudienceErrorItemErrorErrorType = "VALUE_OUT_OF_RANGE"
)

// Defines values for DspAudienceRuleAttributeType.
const (
	ASIN DspAudienceRuleAttributeType = "ASIN"
)

// Defines values for DspAudienceRuleClause.
const (
	INCLUDE DspAudienceRuleClause = "INCLUDE"
)

// Defines values for DspAudienceRuleOperator.
const (
	ONEOF DspAudienceRuleOperator = "ONE_OF"
)

// DspAudienceCreateRequestItem Complete audience model to be used for creation of the audience.
type DspAudienceCreateRequestItem struct {
	// AudienceType Type of audience to create.
	AudienceType DspAudienceCreateRequestItemAudienceType `json:"audienceType"`

	// Description The audience description.
	Description string `json:"description"`

	// IdempotencyKey The unique UUID for this requested audience.
	IdempotencyKey openapi_types.UUID `json:"idempotencyKey"`

	// Lookback The specified time period (in days) to include those who performed the action in the audience.
	// Lookback Constraints Table: Provides available valid values of lookback allowed for given audienceType
	//   | audienceType | lookback range |
	//   |------------------------------|-------|
	//   | PRODUCT_PURCHASES            | 1-365 |
	//   | PRODUCT_VIEWS                |  1-90 |
	//   | PRODUCT_SEARCH               |  1-90 |
	//   | PRODUCT_SIMS                 |  1-90 |
	//   | WHOLE_FOODS_MARKET_PURCHASES | 1-365 |
	Lookback int `json:"lookback"`

	// Name The audience name.
	Name string `json:"name"`

	// Rules The set of rules defining an audience; these rules will be ORed.
	Rules []DspAudienceRule `json:"rules"`
}

// DspAudienceCreateRequestItemAudienceType Type of audience to create.
type DspAudienceCreateRequestItemAudienceType string

// DspAudienceError The error object.
type DspAudienceError struct {
	Message string `json:"message"`
}

// DspAudienceErrorItem The error response object.
type DspAudienceErrorItem struct {
	Errors []DspAudienceErrorItemError `json:"errors"`

	// IdempotencyKey The UUID provided in the request for creation of this audience.
	IdempotencyKey string `json:"idempotencyKey"`

	// Index The index of the DspAudienceCreateRequestItem from the request, e.g. 1st item in the batch request will correspond to index 0 in the response.
	Index int `json:"index"`

	// Message A human-readable description of the response.
	Message string `json:"message"`
}

// DspAudienceErrorItemError The error object.
type DspAudienceErrorItemError struct {
	ErrorType DspAudienceErrorItemErrorErrorType `json:"errorType"`
	FieldName *string                            `json:"fieldName,omitempty"`
	Message   string                             `json:"message"`
}

// DspAudienceErrorItemErrorErrorType defines model for DspAudienceErrorItemError.ErrorType.
type DspAudienceErrorItemErrorErrorType string

// DspAudienceResponse This holds an array of successful items and an array of error items from the request.
type DspAudienceResponse struct {
	// Error The items in this array represent items in the request that failed.
	Error []DspAudienceErrorItem `json:"error"`

	// Success The items in this array represent items in the request that were successful.
	Success []DspAudienceSuccessItem `json:"success"`
}

// DspAudienceRule A rule for defining an audience.
//
//	**Rule Constraints Table**: Provides available valid combinations of parameters allowed in DspAudienceRule
//	 | audienceType | attributeType | attributeValues | max attribute values  |  max rules |
//	 |------------------------------|-----------|-----------------------------------------------------------------------------------------------------|------|---|
//	 | PRODUCT_PURCHASES            |      ASIN | product IDs (ASINs) e.g. B08V4T57R2                                                                 | 1000 | 1 |
//	 | PRODUCT_VIEWS                |      ASIN | product IDs (ASINs) e.g. B08V4T57R2                                                                 | 1000 | 1 |
//	 | PRODUCT_SEARCH               |      ASIN | product IDs (ASINs) e.g. B08V4T57R2                                                                 | 1000 | 1 |
//	 | PRODUCT_SIMS                 |      ASIN | product IDs (ASINs) e.g. B08V4T57R2                                                                 | 1000 | 1 |
//	 | WHOLE_FOODS_MARKET_PURCHASES |      ASIN | Whole Foods Market product IDs (ASINs) e.g. B01B2OVUAG                                              |  500 | 1 |
type DspAudienceRule struct {
	// AttributeType For a given audienceType, the type of the attributes being supplied.
	AttributeType DspAudienceRuleAttributeType `json:"attributeType"`

	// AttributeValues For a given audienceType and attributeType combination, the attribute values being supplied.
	AttributeValues []string `json:"attributeValues"`

	// Clause This parameter is used to include or exclude this particular rule. Currently only include is supported.
	Clause DspAudienceRuleClause `json:"clause"`

	// Operator For a given attributeType, the operator used for attributeValues.
	Operator DspAudienceRuleOperator `json:"operator"`
}

// DspAudienceRuleAttributeType For a given audienceType, the type of the attributes being supplied.
type DspAudienceRuleAttributeType string

// DspAudienceRuleClause This parameter is used to include or exclude this particular rule. Currently only include is supported.
type DspAudienceRuleClause string

// DspAudienceRuleOperator For a given attributeType, the operator used for attributeValues.
type DspAudienceRuleOperator string

// DspAudienceSuccessItem The success response object.
type DspAudienceSuccessItem struct {
	// AudienceId The audience identifier.
	AudienceId string `json:"audienceId"`

	// IdempotencyKey The UUID provided in the request for creation of this audience.
	IdempotencyKey string `json:"idempotencyKey"`

	// Index The index of the DspAudienceCreateRequestItem from the request, e.g. 1st item in the batch request will correspond to index 0 in the response.
	Index int `json:"index"`
}

// DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONBody defines parameters for DspCreateAudiencesPost.
type DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONBody = []DspAudienceCreateRequestItem

// DspCreateAudiencesPostParams defines parameters for DspCreateAudiencesPost.
type DspCreateAudiencesPostParams struct {
	// AdvertiserId The advertiser to create audience for.
	AdvertiserId string `form:"AdvertiserId" json:"AdvertiserId"`

	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`

	// AmazonAdvertisingAPIScope The identifier of a profile associated with the advertiser account. Use `GET` method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id `profileId` from the response to pass it as input.
	AmazonAdvertisingAPIScope string `json:"Amazon-Advertising-API-Scope"`
}

// DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody defines body for DspCreateAudiencesPost for application/vnd.dspaudiences.v1+json ContentType.
type DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody = DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

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
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
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

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// DspCreateAudiencesPostWithBody request with any body
	DspCreateAudiencesPostWithBody(ctx context.Context, params *DspCreateAudiencesPostParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	DspCreateAudiencesPostWithApplicationVndDspaudiencesV1PlusJSONBody(ctx context.Context, params *DspCreateAudiencesPostParams, body DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DspCreateAudiencesPostWithBody(ctx context.Context, params *DspCreateAudiencesPostParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDspCreateAudiencesPostRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DspCreateAudiencesPostWithApplicationVndDspaudiencesV1PlusJSONBody(ctx context.Context, params *DspCreateAudiencesPostParams, body DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDspCreateAudiencesPostRequestWithApplicationVndDspaudiencesV1PlusJSONBody(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDspCreateAudiencesPostRequestWithApplicationVndDspaudiencesV1PlusJSONBody calls the generic DspCreateAudiencesPost builder with application/vnd.dspaudiences.v1+json body
func NewDspCreateAudiencesPostRequestWithApplicationVndDspaudiencesV1PlusJSONBody(server string, params *DspCreateAudiencesPostParams, body DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDspCreateAudiencesPostRequestWithBody(server, params, "application/vnd.dspaudiences.v1+json", bodyReader)
}

// NewDspCreateAudiencesPostRequestWithBody generates requests for DspCreateAudiencesPost with any type of body
func NewDspCreateAudiencesPostRequestWithBody(server string, params *DspCreateAudiencesPostParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dsp/audiences")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "AdvertiserId", runtime.ParamLocationQuery, params.AdvertiserId); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-ClientId", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIClientId)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-ClientId", headerParam0)

		var headerParam1 string

		headerParam1, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-Scope", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIScope)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-Scope", headerParam1)

	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
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
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// DspCreateAudiencesPostWithBodyWithResponse request with any body
	DspCreateAudiencesPostWithBodyWithResponse(ctx context.Context, params *DspCreateAudiencesPostParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DspCreateAudiencesPostResp, error)

	DspCreateAudiencesPostWithApplicationVndDspaudiencesV1PlusJSONBodyWithResponse(ctx context.Context, params *DspCreateAudiencesPostParams, body DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*DspCreateAudiencesPostResp, error)
}

type DspCreateAudiencesPostResp struct {
	Body                                        []byte
	HTTPResponse                                *http.Response
	ApplicationvndDspaudiencesresponseV1JSON207 *DspAudienceResponse
	ApplicationvndDspaudienceserrorV1JSON400    *DspAudienceError
	ApplicationvndDspaudienceserrorV1JSON403    *DspAudienceError
}

// Status returns HTTPResponse.Status
func (r DspCreateAudiencesPostResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DspCreateAudiencesPostResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DspCreateAudiencesPostWithBodyWithResponse request with arbitrary body returning *DspCreateAudiencesPostResp
func (c *ClientWithResponses) DspCreateAudiencesPostWithBodyWithResponse(ctx context.Context, params *DspCreateAudiencesPostParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DspCreateAudiencesPostResp, error) {
	rsp, err := c.DspCreateAudiencesPostWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDspCreateAudiencesPostResp(rsp)
}

func (c *ClientWithResponses) DspCreateAudiencesPostWithApplicationVndDspaudiencesV1PlusJSONBodyWithResponse(ctx context.Context, params *DspCreateAudiencesPostParams, body DspCreateAudiencesPostApplicationVndDspaudiencesV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*DspCreateAudiencesPostResp, error) {
	rsp, err := c.DspCreateAudiencesPostWithApplicationVndDspaudiencesV1PlusJSONBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDspCreateAudiencesPostResp(rsp)
}

// ParseDspCreateAudiencesPostResp parses an HTTP response from a DspCreateAudiencesPostWithResponse call
func ParseDspCreateAudiencesPostResp(rsp *http.Response) (*DspCreateAudiencesPostResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DspCreateAudiencesPostResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 207:
		var dest DspAudienceResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDspaudiencesresponseV1JSON207 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest DspAudienceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDspaudienceserrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest DspAudienceError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDspaudienceserrorV1JSON403 = &dest

	}

	return response, nil
}