// Package dspreport provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package dspreport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for CreateReportRequestBodyV3Dimensions.
const (
	CreateReportRequestBodyV3DimensionsBROWSERTYPE       CreateReportRequestBodyV3Dimensions = "BROWSER_TYPE"
	CreateReportRequestBodyV3DimensionsBROWSERVERSION    CreateReportRequestBodyV3Dimensions = "BROWSER_VERSION"
	CreateReportRequestBodyV3DimensionsCITY              CreateReportRequestBodyV3Dimensions = "CITY"
	CreateReportRequestBodyV3DimensionsCONVERSIONSOURCE  CreateReportRequestBodyV3Dimensions = "CONVERSION_SOURCE"
	CreateReportRequestBodyV3DimensionsCOUNTRY           CreateReportRequestBodyV3Dimensions = "COUNTRY"
	CreateReportRequestBodyV3DimensionsCREATIVE          CreateReportRequestBodyV3Dimensions = "CREATIVE"
	CreateReportRequestBodyV3DimensionsDEAL              CreateReportRequestBodyV3Dimensions = "DEAL"
	CreateReportRequestBodyV3DimensionsDEVICETYPE        CreateReportRequestBodyV3Dimensions = "DEVICE_TYPE"
	CreateReportRequestBodyV3DimensionsDMA               CreateReportRequestBodyV3Dimensions = "DMA"
	CreateReportRequestBodyV3DimensionsENVIRONMENTTYPE   CreateReportRequestBodyV3Dimensions = "ENVIRONMENT_TYPE"
	CreateReportRequestBodyV3DimensionsLINEITEM          CreateReportRequestBodyV3Dimensions = "LINE_ITEM"
	CreateReportRequestBodyV3DimensionsOPERATINGSYSTEM   CreateReportRequestBodyV3Dimensions = "OPERATING_SYSTEM"
	CreateReportRequestBodyV3DimensionsORDER             CreateReportRequestBodyV3Dimensions = "ORDER"
	CreateReportRequestBodyV3DimensionsPOSTALCODE        CreateReportRequestBodyV3Dimensions = "POSTAL_CODE"
	CreateReportRequestBodyV3DimensionsSITE              CreateReportRequestBodyV3Dimensions = "SITE"
	CreateReportRequestBodyV3DimensionsSTATECOUNTYREGION CreateReportRequestBodyV3Dimensions = "STATE_COUNTY_REGION"
	CreateReportRequestBodyV3DimensionsSUPPLY            CreateReportRequestBodyV3Dimensions = "SUPPLY"
)

// Defines values for CreateReportRequestBodyV3Format.
const (
	CreateReportRequestBodyV3FormatCSV  CreateReportRequestBodyV3Format = "CSV"
	CreateReportRequestBodyV3FormatJSON CreateReportRequestBodyV3Format = "JSON"
)

// Defines values for CreateReportRequestBodyV3TimeUnit.
const (
	DAILY   CreateReportRequestBodyV3TimeUnit = "DAILY"
	SUMMARY CreateReportRequestBodyV3TimeUnit = "SUMMARY"
)

// Defines values for CreateReportRequestBodyV3Type.
const (
	CreateReportRequestBodyV3TypeAUDIENCE         CreateReportRequestBodyV3Type = "AUDIENCE"
	CreateReportRequestBodyV3TypeCAMPAIGN         CreateReportRequestBodyV3Type = "CAMPAIGN"
	CreateReportRequestBodyV3TypeCONVERSIONSOURCE CreateReportRequestBodyV3Type = "CONVERSION_SOURCE"
	CreateReportRequestBodyV3TypeGEOGRAPHY        CreateReportRequestBodyV3Type = "GEOGRAPHY"
	CreateReportRequestBodyV3TypeINVENTORY        CreateReportRequestBodyV3Type = "INVENTORY"
	CreateReportRequestBodyV3TypePRODUCTS         CreateReportRequestBodyV3Type = "PRODUCTS"
	CreateReportRequestBodyV3TypeTECHNOLOGY       CreateReportRequestBodyV3Type = "TECHNOLOGY"
)

// Defines values for ReportMetadataV3Format.
const (
	ReportMetadataV3FormatCSV  ReportMetadataV3Format = "CSV"
	ReportMetadataV3FormatJSON ReportMetadataV3Format = "JSON"
)

// Defines values for ReportMetadataV3Status.
const (
	FAILURE    ReportMetadataV3Status = "FAILURE"
	INPROGRESS ReportMetadataV3Status = "IN_PROGRESS"
	SUCCESS    ReportMetadataV3Status = "SUCCESS"
)

// Defines values for ReportMetadataV3Type.
const (
	AUDIENCE         ReportMetadataV3Type = "AUDIENCE"
	CAMPAIGN         ReportMetadataV3Type = "CAMPAIGN"
	CONVERSIONSOURCE ReportMetadataV3Type = "CONVERSION_SOURCE"
	GEOGRAPHY        ReportMetadataV3Type = "GEOGRAPHY"
	INVENTORY        ReportMetadataV3Type = "INVENTORY"
	PRODUCTS         ReportMetadataV3Type = "PRODUCTS"
	TECHNOLOGY       ReportMetadataV3Type = "TECHNOLOGY"
)

// CreateReportRequestBodyV3 defines model for CreateReportRequestBodyV3.
type CreateReportRequestBodyV3 struct {
	// AdvertiserIds List of advertisers specified by identifier to include in the report. This should not be present if accountId is advertiser. To learn more about when to use advertiserIds, see [Reporting by account type](https://advertising.amazon.com/API/docs/en-us/reporting/dsp/reporting-by-account-type).
	AdvertiserIds *[]string `json:"advertiserIds,omitempty"`

	// Dimensions List of dimensions to include in the report. Specify one or many comma-delimited strings of dimensions. For example: ["ORDER", "LINE_ITEM", "CREATIVE"]. Adding a dimension in this array determines the aggregation level of the report data and also adds the fields for that dimension in the report. If the list is null or empty, the aggregation of the report data is at `ORDER` level. The allowed values can be used together in this array as an allowed value in which case the report aggregation will be at the lowest aggregation level and the report will contain the fields for all the dimensions included in the report. To see a list of metrics available by dimension, see [Dimensions](https://advertising.amazon.com/API/docs/en-us/reporting/dsp/dimensions).
	Dimensions *[]CreateReportRequestBodyV3Dimensions `json:"dimensions,omitempty"`

	// EndDate Date in yyyy-MM-dd format. The report contains only metrics generated on the specified date range between startDate and endDate. The maximum date range between startDate and endDate is 31 days. The endDate can be up to 90 days older from today.
	EndDate openapi_types.Date `json:"endDate"`

	// Format The report file format.
	Format *CreateReportRequestBodyV3Format `json:"format,omitempty"`

	// Metrics Specify a list of metrics field names to include in the report. For example: ["impressions", "clickThroughs", "CTR", "eCPC", "totalCost", "eCPM"]. If no metric field names are specified, only the default fields and selected `DIMENSION` fields are included by default. Specifying default fields returns an error. To view the metrics available by report type, see [DSP report types](https://advertising.amazon.com/API/docs/en-us/reporting/dsp/report-types)
	Metrics *[]string `json:"metrics,omitempty"`

	// OrderIds List of orders specified by identifier to include in the report.
	OrderIds *[]string `json:"orderIds,omitempty"`

	// StartDate Date in yyyy-MM-dd format. The report contains only metrics generated on the specified date range between startDate and endDate. The maximum date range between startDate and endDate is 31 days. The startDate can be up to 90 days older from today.
	StartDate openapi_types.Date `json:"startDate"`

	// TimeUnit Adding timeUnit determines the aggregation level (`SUMMARY` or `DAILY`) of the report data. If the timeUnit is null or empty, the aggregation of the report data is at the `SUMMARY` level and aggregated at the time period specified. `DAILY` timeUnit is not supported for `AUDIENCE` report type. The report will contain the fields based on timeUnit:<details><summary>`SUMMARY`</summary>intervalStart</br>intervalEnd</details></br><details><summary>`DAILY`</summary>Date</details>
	TimeUnit *CreateReportRequestBodyV3TimeUnit `json:"timeUnit,omitempty"`

	// Type The report type.
	Type *CreateReportRequestBodyV3Type `json:"type,omitempty"`
}

// CreateReportRequestBodyV3Dimensions defines model for CreateReportRequestBodyV3.Dimensions.
type CreateReportRequestBodyV3Dimensions string

// CreateReportRequestBodyV3Format The report file format.
type CreateReportRequestBodyV3Format string

// CreateReportRequestBodyV3TimeUnit Adding timeUnit determines the aggregation level (`SUMMARY` or `DAILY`) of the report data. If the timeUnit is null or empty, the aggregation of the report data is at the `SUMMARY` level and aggregated at the time period specified. `DAILY` timeUnit is not supported for `AUDIENCE` report type. The report will contain the fields based on timeUnit:<details><summary>`SUMMARY`</summary>intervalStart</br>intervalEnd</details></br><details><summary>`DAILY`</summary>Date</details>
type CreateReportRequestBodyV3TimeUnit string

// CreateReportRequestBodyV3Type The report type.
type CreateReportRequestBodyV3Type string

// DSPReportsError The error response object.
type DSPReportsError struct {
	// Errors A list of errors. Please check the values in this field for report validation errors.
	Errors *[]DSPReportsSubError `json:"errors,omitempty"`

	// Message A human-readable description of the response.
	Message string `json:"message"`

	// RequestId A unique identifier of the request.
	RequestId *string `json:"requestId,omitempty"`
}

// DSPReportsSubError The sub-error object.
type DSPReportsSubError struct {
	// ErrorType Enumerated error type.
	ErrorType string `json:"errorType"`

	// Field Request body field which is cause of the error.
	Field *string `json:"field,omitempty"`

	// Message Detailed error description
	Message string `json:"message"`
}

// ReportMetadataV3 defines model for ReportMetadataV3.
type ReportMetadataV3 struct {
	// Expiration The expiration time of the URI in the location property in date-time format(yyyy-MM-ddTHH:mm:ss). The expiration time is the time when the download link expires.
	Expiration *time.Time `json:"expiration,omitempty"`

	// Format The data format of the report.
	Format *ReportMetadataV3Format `json:"format,omitempty"`

	// Location The URI address of the report.
	Location *string `json:"location,omitempty"`

	// ReportId The identifier of the report.
	ReportId *string `json:"reportId,omitempty"`

	// Status The build status of the report.
	Status *ReportMetadataV3Status `json:"status,omitempty"`

	// StatusDetails  A human-readable description of the current status.
	StatusDetails *string `json:"statusDetails,omitempty"`

	// Type The type of report.
	Type *ReportMetadataV3Type `json:"type,omitempty"`
}

// ReportMetadataV3Format The data format of the report.
type ReportMetadataV3Format string

// ReportMetadataV3Status The build status of the report.
type ReportMetadataV3Status string

// ReportMetadataV3Type The type of report.
type ReportMetadataV3Type string

// DSPReportsBadGateway The error response object.
type DSPReportsBadGateway = DSPReportsError

// DSPReportsBadRequest The error response object.
type DSPReportsBadRequest = DSPReportsError

// DSPReportsEndpointRequestTimedOut The error response object.
type DSPReportsEndpointRequestTimedOut = DSPReportsError

// DSPReportsForbidden The error response object.
type DSPReportsForbidden = DSPReportsError

// DSPReportsInternalServerError The error response object.
type DSPReportsInternalServerError = DSPReportsError

// DSPReportsNotFound The error response object.
type DSPReportsNotFound = DSPReportsError

// DSPReportsServiceUnavailable The error response object.
type DSPReportsServiceUnavailable = DSPReportsError

// DSPReportsTooManyRequests The error response object.
type DSPReportsTooManyRequests = DSPReportsError

// DSPReportsUnauthorized The error response object.
type DSPReportsUnauthorized = DSPReportsError

// DSPReportsUnprocessableEntity The error response object.
type DSPReportsUnprocessableEntity = DSPReportsError

// DSPReportsUnsupportedMediaType The error response object.
type DSPReportsUnsupportedMediaType = DSPReportsError

// CreateReportV3Params defines parameters for CreateReportV3.
type CreateReportV3Params struct {
	// AmazonAdvertisingAPIClientId The client identifier. See [create API authorization and refresh tokens](https://advertising.amazon.com/API/docs/en-us/setting-up/generate-api-tokens) for more information.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`
}

// GetCampaignReportV3Params defines parameters for GetCampaignReportV3.
type GetCampaignReportV3Params struct {
	// AmazonAdvertisingAPIClientId The client identifier. See [create API authorization and refresh tokens](https://advertising.amazon.com/API/docs/en-us/setting-up/generate-api-tokens) for more information.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`
}

// CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody defines body for CreateReportV3 for application/vnd.dspcreatereports.v3+json ContentType.
type CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody = CreateReportRequestBodyV3

// RequestEditorFn is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// ResponseEditorFn is the function signature for the ResponseEditor callback function
type ResponseEditorFn func(ctx context.Context, rsp *http.Response) error

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

	// A callback for modifying response which are generated after receive from the network.
	ResponseEditors []ResponseEditorFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
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
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v2.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
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

// WithResponseEditorFn allows setting up a callback function, which will be
// called right after receive the response.
func WithResponseEditorFn(fn ResponseEditorFn) ClientOption {
	return func(c *Client) error {
		c.ResponseEditors = append(c.ResponseEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateReportV3WithBody request with any body
	CreateReportV3WithBody(ctx context.Context, accountId string, params *CreateReportV3Params, contentType string, body io.Reader) (*http.Response, error)

	CreateReportV3WithApplicationVndDspcreatereportsV3PlusJSONBody(ctx context.Context, accountId string, params *CreateReportV3Params, body CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody) (*http.Response, error)

	// GetCampaignReportV3 request
	GetCampaignReportV3(ctx context.Context, accountId string, reportId string, params *GetCampaignReportV3Params) (*http.Response, error)
}

func (c *Client) CreateReportV3WithBody(ctx context.Context, accountId string, params *CreateReportV3Params, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewCreateReportV3RequestWithBody(c.Server, accountId, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if err := c.applyReqEditors(ctx, req); err != nil {
		return nil, err
	}
	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := c.applyRspEditor(ctx, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) CreateReportV3WithApplicationVndDspcreatereportsV3PlusJSONBody(ctx context.Context, accountId string, params *CreateReportV3Params, body CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody) (*http.Response, error) {
	req, err := NewCreateReportV3RequestWithApplicationVndDspcreatereportsV3PlusJSONBody(c.Server, accountId, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if err := c.applyReqEditors(ctx, req); err != nil {
		return nil, err
	}
	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := c.applyRspEditor(ctx, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) GetCampaignReportV3(ctx context.Context, accountId string, reportId string, params *GetCampaignReportV3Params) (*http.Response, error) {
	req, err := NewGetCampaignReportV3Request(c.Server, accountId, reportId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if err := c.applyReqEditors(ctx, req); err != nil {
		return nil, err
	}
	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := c.applyRspEditor(ctx, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

// NewCreateReportV3RequestWithApplicationVndDspcreatereportsV3PlusJSONBody calls the generic CreateReportV3 builder with application/vnd.dspcreatereports.v3+json body
func NewCreateReportV3RequestWithApplicationVndDspcreatereportsV3PlusJSONBody(server string, accountId string, params *CreateReportV3Params, body CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateReportV3RequestWithBody(server, accountId, params, "application/vnd.dspcreatereports.v3+json", bodyReader)
}

// NewCreateReportV3RequestWithBody generates requests for CreateReportV3 with any type of body
func NewCreateReportV3RequestWithBody(server string, accountId string, params *CreateReportV3Params, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "accountId", runtime.ParamLocationPath, accountId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/accounts/%s/dsp/reports", pathParam0)
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

	}

	return req, nil
}

// NewGetCampaignReportV3Request generates requests for GetCampaignReportV3
func NewGetCampaignReportV3Request(server string, accountId string, reportId string, params *GetCampaignReportV3Params) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "accountId", runtime.ParamLocationPath, accountId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "reportId", runtime.ParamLocationPath, reportId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/accounts/%s/dsp/reports/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-ClientId", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIClientId)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-ClientId", headerParam0)

	}

	return req, nil
}

func (c *Client) applyReqEditors(ctx context.Context, req *http.Request) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) applyRspEditor(ctx context.Context, rsp *http.Response) error {
	for _, r := range c.ResponseEditors {
		if err := r(ctx, rsp); err != nil {
			return err
		}
	}
	return nil
} // ClientWithResponses builds on ClientInterface to offer response payloads
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
	// CreateReportV3WithBodyWithResponse request with any body
	CreateReportV3WithBodyWithResponse(ctx context.Context, accountId string, params *CreateReportV3Params, contentType string, body io.Reader) (*CreateReportV3Resp, error)

	CreateReportV3WithApplicationVndDspcreatereportsV3PlusJSONBodyWithResponse(ctx context.Context, accountId string, params *CreateReportV3Params, body CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody) (*CreateReportV3Resp, error)

	// GetCampaignReportV3WithResponse request
	GetCampaignReportV3WithResponse(ctx context.Context, accountId string, reportId string, params *GetCampaignReportV3Params) (*GetCampaignReportV3Resp, error)
}

type CreateReportV3Resp struct {
	Body                                    []byte
	HTTPResponse                            *http.Response
	ApplicationvndDspcreatereportsV3JSON200 *ReportMetadataV3
	ApplicationvndDsperrorV1JSON400         *DSPReportsBadRequest
	ApplicationvndDsperrorV1JSON401         *DSPReportsUnauthorized
	ApplicationvndDsperrorV1JSON403         *DSPReportsForbidden
	ApplicationvndDsperrorV1JSON404         *DSPReportsNotFound
	ApplicationvndDsperrorV1JSON415         *DSPReportsUnsupportedMediaType
	ApplicationvndDsperrorV1JSON422         *DSPReportsUnprocessableEntity
	ApplicationvndDsperrorV1JSON429         *DSPReportsTooManyRequests
	ApplicationvndDsperrorV1JSON500         *DSPReportsInternalServerError
	ApplicationvndDsperrorV1JSON502         *DSPReportsBadGateway
	ApplicationvndDsperrorV1JSON503         *DSPReportsServiceUnavailable
	ApplicationvndDsperrorV1JSON504         *DSPReportsEndpointRequestTimedOut
}

// Status returns HTTPResponse.Status
func (r CreateReportV3Resp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateReportV3Resp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCampaignReportV3Resp struct {
	Body                                 []byte
	HTTPResponse                         *http.Response
	ApplicationvndDspgetreportsV3JSON200 *ReportMetadataV3
	ApplicationvndDsperrorV1JSON400      *DSPReportsBadRequest
	ApplicationvndDsperrorV1JSON403      *DSPReportsForbidden
	ApplicationvndDsperrorV1JSON404      *DSPReportsNotFound
	ApplicationvndDsperrorV1JSON415      *DSPReportsUnsupportedMediaType
	ApplicationvndDsperrorV1JSON429      *DSPReportsTooManyRequests
	ApplicationvndDsperrorV1JSON500      *DSPReportsInternalServerError
	ApplicationvndDsperrorV1JSON502      *DSPReportsBadGateway
	ApplicationvndDsperrorV1JSON503      *DSPReportsServiceUnavailable
	ApplicationvndDsperrorV1JSON504      *DSPReportsEndpointRequestTimedOut
}

// Status returns HTTPResponse.Status
func (r GetCampaignReportV3Resp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCampaignReportV3Resp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateReportV3WithBodyWithResponse request with arbitrary body returning *CreateReportV3Resp
func (c *ClientWithResponses) CreateReportV3WithBodyWithResponse(ctx context.Context, accountId string, params *CreateReportV3Params, contentType string, body io.Reader) (*CreateReportV3Resp, error) {
	rsp, err := c.CreateReportV3WithBody(ctx, accountId, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportV3Resp(rsp)
}

func (c *ClientWithResponses) CreateReportV3WithApplicationVndDspcreatereportsV3PlusJSONBodyWithResponse(ctx context.Context, accountId string, params *CreateReportV3Params, body CreateReportV3ApplicationVndDspcreatereportsV3PlusJSONRequestBody) (*CreateReportV3Resp, error) {
	rsp, err := c.CreateReportV3WithApplicationVndDspcreatereportsV3PlusJSONBody(ctx, accountId, params, body)
	if err != nil {
		return nil, err
	}
	return ParseCreateReportV3Resp(rsp)
}

// GetCampaignReportV3WithResponse request returning *GetCampaignReportV3Resp
func (c *ClientWithResponses) GetCampaignReportV3WithResponse(ctx context.Context, accountId string, reportId string, params *GetCampaignReportV3Params) (*GetCampaignReportV3Resp, error) {
	rsp, err := c.GetCampaignReportV3(ctx, accountId, reportId, params)
	if err != nil {
		return nil, err
	}
	return ParseGetCampaignReportV3Resp(rsp)
}

// ParseCreateReportV3Resp parses an HTTP response from a CreateReportV3WithResponse call
func ParseCreateReportV3Resp(rsp *http.Response) (*CreateReportV3Resp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateReportV3Resp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ReportMetadataV3
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDspcreatereportsV3JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest DSPReportsBadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest DSPReportsUnauthorized
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest DSPReportsForbidden
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest DSPReportsNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest DSPReportsUnsupportedMediaType
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest DSPReportsUnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest DSPReportsTooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest DSPReportsInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest DSPReportsBadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON502 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest DSPReportsServiceUnavailable
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON503 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 504:
		var dest DSPReportsEndpointRequestTimedOut
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON504 = &dest

	}

	return response, nil
}

// ParseGetCampaignReportV3Resp parses an HTTP response from a GetCampaignReportV3WithResponse call
func ParseGetCampaignReportV3Resp(rsp *http.Response) (*GetCampaignReportV3Resp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetCampaignReportV3Resp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ReportMetadataV3
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDspgetreportsV3JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest DSPReportsBadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest DSPReportsForbidden
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest DSPReportsNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest DSPReportsUnsupportedMediaType
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest DSPReportsTooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest DSPReportsInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest DSPReportsBadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON502 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest DSPReportsServiceUnavailable
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON503 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 504:
		var dest DSPReportsEndpointRequestTimedOut
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndDsperrorV1JSON504 = &dest

	}

	return response, nil
}
