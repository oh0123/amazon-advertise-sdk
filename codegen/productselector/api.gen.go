// Package productselector provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package productselector

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
)

// Defines values for ProductMetadataRequestAdType.
const (
	SB ProductMetadataRequestAdType = "SB"
	SD ProductMetadataRequestAdType = "SD"
	SP ProductMetadataRequestAdType = "SP"
)

// Defines values for ProductMetadataRequestSortBy.
const (
	CREATEDDATE ProductMetadataRequestSortBy = "CREATED_DATE"
	SUGGESTED   ProductMetadataRequestSortBy = "SUGGESTED"
)

// Defines values for ProductMetadataRequestSortOrder.
const (
	ASC  ProductMetadataRequestSortOrder = "ASC"
	DESC ProductMetadataRequestSortOrder = "DESC"
)

// ProductMetadataModel defines model for ProductMetadataModel.
type ProductMetadataModel struct {
	// Asin ASIN of the item
	Asin *string `json:"asin,omitempty"`

	// Availability Stock availability:
	//  * IN_STOCK - The item is in stock.
	//  * IN_STOCK_SCARCE - The item is in stock, but stock levels are limited.
	//  * OUT_OF_STOCK - The item is currently out of stock.
	//  * PREORDER - The item is not yet available, but can be pre-ordered.
	//  * LEADTIME - The item is only available after some amount of lead time.
	//  * AVAILABLE_DATE - The item is not available, but will be available on a future date.
	Availability *string `json:"availability,omitempty"`

	// BasisPrice The basis price before the savings are calculated
	BasisPrice *BasisPrice `json:"basisPrice,omitempty"`

	// BestSellerRank Best seller rank position in the category
	BestSellerRank *string `json:"bestSellerRank,omitempty"`

	// Brand Brand name of the item
	Brand *string `json:"brand,omitempty"`

	// Category Category (browse node) name of the ASIN
	Category *string `json:"category,omitempty"`

	// CreatedDate Date the item was first available on Amazon
	CreatedDate *string `json:"createdDate,omitempty"`

	// EligibilityStatus Eligibility status for advertising:
	//  * ELIGIBLE - Eligible for advertising
	//  * INELIGIBLE - Ineligible for advertising
	EligibilityStatus *string `json:"eligibilityStatus,omitempty"`

	// ImageUrl Url to the product image
	ImageUrl *string `json:"imageUrl,omitempty"`

	// IneligibilityCodes List of ineligible status identifier
	IneligibilityCodes *[]string `json:"ineligibilityCodes,omitempty"`

	// IneligibilityReasons List of reasons that made this item ineligible to be advertised
	IneligibilityReasons *[]string `json:"ineligibilityReasons,omitempty"`

	// PriceToPay The price customer would pay for the buying option
	PriceToPay *PriceToPay `json:"priceToPay,omitempty"`

	// Sku sku of the item
	Sku *string `json:"sku,omitempty"`

	// Title Product title of the item
	Title *string `json:"title,omitempty"`

	// VariationList List of ASIN variations of the current item
	VariationList *[]string `json:"variationList,omitempty"`
}

// ProductMetadataRequest defines model for ProductMetadataRequest.
type ProductMetadataRequest struct {
	// AdType Program type. Required if checks advertising eligibility:
	//  * SP - Sponsored Product
	//  * SB - Sponsored Brand
	//  * SD - Sponsored Display
	AdType *ProductMetadataRequestAdType `json:"adType,omitempty"`

	// Asins Specific asins to search for in the advertiser's inventory. Cannot use together with skus or searchStr input types.
	Asins *[]string `json:"asins,omitempty"`

	// CheckEligibility Whether advertising eligibility info is required
	CheckEligibility *bool `json:"checkEligibility,omitempty"`

	// CheckItemDetails Whether item details such as name, image, and price is required.
	CheckItemDetails *bool `json:"checkItemDetails,omitempty"`

	// CursorToken Pagination token used for the suggested sort type or for author merchant
	CursorToken *string `json:"cursorToken,omitempty"`

	// Locale Optional locale for detail and eligibility response strings. Default to the marketplace locale.
	Locale *string `json:"locale,omitempty"`

	// PageIndex Index of the page to be returned; For author, this value will be ignored, should use cursorToken instead. For seller, results are capped at 10k(pageIndex * pageSize). For vendor, results are capped at 50k.
	PageIndex int32 `json:"pageIndex"`

	// PageSize Number of items to be returned on this page index.
	PageSize int32 `json:"pageSize"`

	// SearchStr Specific string in the item title to search for in the advertiser's inventory. Case insensitive. Cannot use together with asins or skus input types.
	SearchStr *string `json:"searchStr,omitempty"`

	// Skus Specific SKUs to search for in the advertiser's inventory. Currently only support SP program type for sellers. Cannot use together with asins or searchStr input types.
	Skus *[]string `json:"skus,omitempty"`

	// SortBy Sort option for the result. Currently only support SP program type for sellers:
	//  * SUGGESTED - Suggested products are those most likely to engage customers, and have a higher chance of generating clicks if advertised.
	//  * CREATED_DATE - Date the item listing was created
	SortBy *ProductMetadataRequestSortBy `json:"sortBy,omitempty"`

	// SortOrder Sort order (has to be DESC for the suggested sort type):
	//  * ASC - Ascending, from A to Z
	//  * DESC - Descending, from Z to A
	SortOrder *ProductMetadataRequestSortOrder `json:"sortOrder,omitempty"`
}

// ProductMetadataRequestAdType Program type. Required if checks advertising eligibility:
//   - SP - Sponsored Product
//   - SB - Sponsored Brand
//   - SD - Sponsored Display
type ProductMetadataRequestAdType string

// ProductMetadataRequestSortBy Sort option for the result. Currently only support SP program type for sellers:
//   - SUGGESTED - Suggested products are those most likely to engage customers, and have a higher chance of generating clicks if advertised.
//   - CREATED_DATE - Date the item listing was created
type ProductMetadataRequestSortBy string

// ProductMetadataRequestSortOrder Sort order (has to be DESC for the suggested sort type):
//   - ASC - Ascending, from A to Z
//   - DESC - Descending, from Z to A
type ProductMetadataRequestSortOrder string

// ProductMetadataResponse defines model for ProductMetadataResponse.
type ProductMetadataResponse struct {
	ProductMetadataList *[]ProductMetadataModel `json:"ProductMetadataList,omitempty"`

	// CursorToken Pagination token for later requests with specific sort type to use as the page index instead. Empty cursorToken means no further data is present at Server side.
	CursorToken *string `json:"cursorToken,omitempty"`
}

// BasisPrice The basis price before the savings are calculated
type BasisPrice struct {
	// Amount Price amount
	Amount *float32 `json:"amount,omitempty"`

	// Currency Currency of the price
	Currency *string `json:"currency,omitempty"`
}

// Error Error response object.
type Error struct {
	// Code Enumerated error type.
	Code *string `json:"code,omitempty"`

	// Details A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// PriceToPay The price customer would pay for the buying option
type PriceToPay struct {
	// Amount Price amount
	Amount *float32 `json:"amount,omitempty"`

	// Currency Currency of the price
	Currency *string `json:"currency,omitempty"`
}

// ProductMetadataParams defines parameters for ProductMetadata.
type ProductMetadataParams struct {
	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`

	// AmazonAdvertisingAPIScope The identifier of a profile associated with the advertiser account. Use `GET` method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id `profileId` from the response to pass it as input.
	AmazonAdvertisingAPIScope string `json:"Amazon-Advertising-API-Scope"`
}

// ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody defines body for ProductMetadata for application/vnd.productmetadatarequest.v1+json ContentType.
type ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody = ProductMetadataRequest

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
	// ProductMetadataWithBody request with any body
	ProductMetadataWithBody(ctx context.Context, params *ProductMetadataParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ProductMetadataWithApplicationVndProductmetadatarequestV1PlusJSONBody(ctx context.Context, params *ProductMetadataParams, body ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ProductMetadataWithBody(ctx context.Context, params *ProductMetadataParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewProductMetadataRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ProductMetadataWithApplicationVndProductmetadatarequestV1PlusJSONBody(ctx context.Context, params *ProductMetadataParams, body ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewProductMetadataRequestWithApplicationVndProductmetadatarequestV1PlusJSONBody(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewProductMetadataRequestWithApplicationVndProductmetadatarequestV1PlusJSONBody calls the generic ProductMetadata builder with application/vnd.productmetadatarequest.v1+json body
func NewProductMetadataRequestWithApplicationVndProductmetadatarequestV1PlusJSONBody(server string, params *ProductMetadataParams, body ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewProductMetadataRequestWithBody(server, params, "application/vnd.productmetadatarequest.v1+json", bodyReader)
}

// NewProductMetadataRequestWithBody generates requests for ProductMetadata with any type of body
func NewProductMetadataRequestWithBody(server string, params *ProductMetadataParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/product/metadata")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
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
	// ProductMetadataWithBodyWithResponse request with any body
	ProductMetadataWithBodyWithResponse(ctx context.Context, params *ProductMetadataParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ProductMetadataResp, error)

	ProductMetadataWithApplicationVndProductmetadatarequestV1PlusJSONBodyWithResponse(ctx context.Context, params *ProductMetadataParams, body ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*ProductMetadataResp, error)
}

type ProductMetadataResp struct {
	Body                                           []byte
	HTTPResponse                                   *http.Response
	ApplicationvndProductmetadataresponseV1JSON200 *ProductMetadataResponse
	JSON400                                        *Error
	JSON401                                        *Error
	JSON403                                        *Error
	JSON404                                        *Error
	JSON429                                        *Error
	JSON500                                        *Error
}

// Status returns HTTPResponse.Status
func (r ProductMetadataResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ProductMetadataResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ProductMetadataWithBodyWithResponse request with arbitrary body returning *ProductMetadataResp
func (c *ClientWithResponses) ProductMetadataWithBodyWithResponse(ctx context.Context, params *ProductMetadataParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ProductMetadataResp, error) {
	rsp, err := c.ProductMetadataWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseProductMetadataResp(rsp)
}

func (c *ClientWithResponses) ProductMetadataWithApplicationVndProductmetadatarequestV1PlusJSONBodyWithResponse(ctx context.Context, params *ProductMetadataParams, body ProductMetadataApplicationVndProductmetadatarequestV1PlusJSONRequestBody, reqEditors ...RequestEditorFn) (*ProductMetadataResp, error) {
	rsp, err := c.ProductMetadataWithApplicationVndProductmetadatarequestV1PlusJSONBody(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseProductMetadataResp(rsp)
}

// ParseProductMetadataResp parses an HTTP response from a ProductMetadataWithResponse call
func ParseProductMetadataResp(rsp *http.Response) (*ProductMetadataResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ProductMetadataResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ProductMetadataResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndProductmetadataresponseV1JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}