// Package brandMetrics provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package brandMetrics

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

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for BrandMetricsGenerateReportRequestFormat.
const (
	BrandMetricsGenerateReportRequestFormatCSV  BrandMetricsGenerateReportRequestFormat = "CSV"
	BrandMetricsGenerateReportRequestFormatJSON BrandMetricsGenerateReportRequestFormat = "JSON"
)

// Defines values for BrandMetricsGenerateReportRequestLookBackPeriod.
const (
	N1cm BrandMetricsGenerateReportRequestLookBackPeriod = "1cm"
	N1m  BrandMetricsGenerateReportRequestLookBackPeriod = "1m"
	N1w  BrandMetricsGenerateReportRequestLookBackPeriod = "1w"
)

// Defines values for BrandMetricsGenerateReportResponseFormat.
const (
	BrandMetricsGenerateReportResponseFormatCSV  BrandMetricsGenerateReportResponseFormat = "CSV"
	BrandMetricsGenerateReportResponseFormatJSON BrandMetricsGenerateReportResponseFormat = "JSON"
)

// Defines values for BrandMetricsGenerateReportResponseStatus.
const (
	BrandMetricsGenerateReportResponseStatusFAILURE    BrandMetricsGenerateReportResponseStatus = "FAILURE"
	BrandMetricsGenerateReportResponseStatusINPROGRESS BrandMetricsGenerateReportResponseStatus = "IN_PROGRESS"
	BrandMetricsGenerateReportResponseStatusSUCCESS    BrandMetricsGenerateReportResponseStatus = "SUCCESS"
)

// Defines values for BrandMetricsGetReportByIdResponseFormat.
const (
	CSV  BrandMetricsGetReportByIdResponseFormat = "CSV"
	JSON BrandMetricsGetReportByIdResponseFormat = "JSON"
)

// Defines values for BrandMetricsGetReportByIdResponseStatus.
const (
	BrandMetricsGetReportByIdResponseStatusFAILURE    BrandMetricsGetReportByIdResponseStatus = "FAILURE"
	BrandMetricsGetReportByIdResponseStatusINPROGRESS BrandMetricsGetReportByIdResponseStatus = "IN_PROGRESS"
	BrandMetricsGetReportByIdResponseStatusSUCCESS    BrandMetricsGetReportByIdResponseStatus = "SUCCESS"
)

// BrandMetricsError The error response object.
type BrandMetricsError struct {
	// Code The HTTP status code of the response.
	Code *string `json:"code,omitempty"`

	// Details A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// BrandMetricsGenerateReportRequest Request object to generate the Brand Metrics Report
type BrandMetricsGenerateReportRequest struct {
	// BrandName Optional. Brand Name. If no Brand Name is passed, then all data available for all brands belonging to the entity are retrieved.
	BrandName *string `json:"brandName,omitempty"`

	// CategoryNodePath Optional. The hierarchical path that leads to a node starting with the root node. If no Category Node Name is passed, then all data available for all brands belonging to the entity are retrieved.
	CategoryNodePath *[]string `json:"categoryNodePath,omitempty"`

	// CategoryNodeTreeName Optional. The node at the top of a browse tree. It is the start node of a tree
	CategoryNodeTreeName *string `json:"categoryNodeTreeName,omitempty"`

	// Format Format of the report
	Format *BrandMetricsGenerateReportRequestFormat `json:"format,omitempty"`

	// LookBackPeriod Currently supported values: "1w" (one week), "1m" (one month) and  "1cm" (one calendar month). This defines the period of time used to determine the number of shoppers in the metrics computation.
	LookBackPeriod *BrandMetricsGenerateReportRequestLookBackPeriod `json:"lookBackPeriod,omitempty"`

	// Metrics Optional. Specify an array of string of metrics field names to include in the report. If no metric field names are specified, all metrics are returned.
	Metrics *[]string `json:"metrics,omitempty"`

	// ReportEndDate Optional. Retrieves metrics with metricsComputationDate between reportStartDate and reportEndDate  (inclusive). The date will be in the Coordinated Universal Time (UTC) timezone in YYYY-MM-DD format. If no date is passed in reportEndDate, all available metrics with metricsComputationDate from the reportStartDate will be provided. If no date is passed for either reportStartDate or reportEndDate, the metrics with the most receont metricsComputationDate will be returned.
	ReportEndDate *openapi_types.Date `json:"reportEndDate,omitempty"`

	// ReportStartDate Optional. Retrieves metrics with metricsComputationDate between reportStartDate and reportEndDate  (inclusive). The date will be in the Coordinated Universal Time (UTC) timezone in YYYY-MM-DD format. If no date is passed in reportStartDate, all available metrics with metricsComputationDate till the reportEndDate will be provided. If no date is passed for either reportStartDate or reportEndDate, the metrics with the most receont metricsComputationDate will be returned.
	ReportStartDate *openapi_types.Date `json:"reportStartDate,omitempty"`
}

// BrandMetricsGenerateReportRequestFormat Format of the report
type BrandMetricsGenerateReportRequestFormat string

// BrandMetricsGenerateReportRequestLookBackPeriod Currently supported values: "1w" (one week), "1m" (one month) and  "1cm" (one calendar month). This defines the period of time used to determine the number of shoppers in the metrics computation.
type BrandMetricsGenerateReportRequestLookBackPeriod string

// BrandMetricsGenerateReportResponse Response object containing Brand Metrics Report metadata
type BrandMetricsGenerateReportResponse struct {
	// Expiration The expiration time of the URI in the location property in milliseconds. The expiration time is the interval between the time the response was generated and the time the URI expires.
	Expiration int64 `json:"expiration"`

	// Format Format of the report
	Format BrandMetricsGenerateReportResponseFormat `json:"format"`

	// Location The URI address of the report.
	Location *string `json:"location,omitempty"`

	// ReportId The identifier of the report.
	ReportId string `json:"reportId"`

	// Status The build status of the report.
	Status BrandMetricsGenerateReportResponseStatus `json:"status"`

	// StatusDetails A human-readable description of the current status.
	StatusDetails string `json:"statusDetails"`
}

// BrandMetricsGenerateReportResponseFormat Format of the report
type BrandMetricsGenerateReportResponseFormat string

// BrandMetricsGenerateReportResponseStatus The build status of the report.
type BrandMetricsGenerateReportResponseStatus string

// BrandMetricsGetReportByIdResponse Response object containing Brand Metrics Report status metadata
type BrandMetricsGetReportByIdResponse struct {
	// BrandsInfo List of first 200 brands for which the Brand Metrics report is generated. The report may contain more than 200 brands. This list is only populated with brands if the Brand Metrics are available for the brands that an advertiser has access to.
	BrandsInfo *[]struct {
		// Id Id
		Id *string `json:"id,omitempty"`

		// Name Brand Name
		Name *string `json:"name,omitempty"`
	} `json:"brandsInfo,omitempty"`

	// Expiration The expiration time of the URI in the location property in milliseconds. The expiration time is the interval between the time the response was generated and the time the URI expires.
	Expiration int64 `json:"expiration"`

	// Format Format of the report
	Format BrandMetricsGetReportByIdResponseFormat `json:"format"`

	// Location The URI address of the report. Only available if the report is generated successfully. The location is empty if the Brand Metrics are not available or if the report is not generated successfully.
	Location *string `json:"location,omitempty"`

	// ReportId The identifier of the report.
	ReportId string `json:"reportId"`

	// Status The build status of the report.
	Status BrandMetricsGetReportByIdResponseStatus `json:"status"`

	// StatusDetails A human-readable description of the current status.
	StatusDetails string `json:"statusDetails"`
}

// BrandMetricsGetReportByIdResponseFormat Format of the report
type BrandMetricsGetReportByIdResponseFormat string

// BrandMetricsGetReportByIdResponseStatus The build status of the report.
type BrandMetricsGetReportByIdResponseStatus string

// GenerateBrandMetricsReportParams defines parameters for GenerateBrandMetricsReport.
type GenerateBrandMetricsReportParams struct {
	// AmazonAdvertisingAPIScope The profile Id, for example, 195213312458027.
	AmazonAdvertisingAPIScope string `json:"Amazon-Advertising-API-Scope"`

	// AmazonAdvertisingAPIClientID The client Id, for example, amzn1.application-oa2-client.8baa9caa3eac48eab89780e73ce03b19.
	AmazonAdvertisingAPIClientID string `json:"Amazon-Advertising-API-ClientID"`
}

// GetBrandMetricsReportParams defines parameters for GetBrandMetricsReport.
type GetBrandMetricsReportParams struct {
	// AmazonAdvertisingAPIScope The profile Id, for example, 195213312458027.
	AmazonAdvertisingAPIScope string `json:"Amazon-Advertising-API-Scope"`

	// AmazonAdvertisingAPIClientID The client Id, for example, amzn1.application-oa2-client.8baa9caa3eac48eab89780e73ce03b19.
	AmazonAdvertisingAPIClientID string `json:"Amazon-Advertising-API-ClientID"`
}

// GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody defines body for GenerateBrandMetricsReport for application/vnd.insightsBrandMetrics.v1+json ContentType.
type GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody = BrandMetricsGenerateReportRequest

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
		client.UserAgent = fmt.Sprintf("go-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
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
	// GenerateBrandMetricsReportWithBody request with any body
	GenerateBrandMetricsReportWithBody(ctx context.Context, params *GenerateBrandMetricsReportParams, contentType string, body io.Reader) (*http.Response, error)

	GenerateBrandMetricsReportWithApplicationVndInsightsBrandMetricsV1PlusJSONBody(ctx context.Context, params *GenerateBrandMetricsReportParams, body GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody) (*http.Response, error)

	// GetBrandMetricsReport request
	GetBrandMetricsReport(ctx context.Context, reportId string, params *GetBrandMetricsReportParams) (*http.Response, error)
}

func (c *Client) GenerateBrandMetricsReportWithBody(ctx context.Context, params *GenerateBrandMetricsReportParams, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewGenerateBrandMetricsReportRequestWithBody(c.Server, params, contentType, body)
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

func (c *Client) GenerateBrandMetricsReportWithApplicationVndInsightsBrandMetricsV1PlusJSONBody(ctx context.Context, params *GenerateBrandMetricsReportParams, body GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody) (*http.Response, error) {
	req, err := NewGenerateBrandMetricsReportRequestWithApplicationVndInsightsBrandMetricsV1PlusJSONBody(c.Server, params, body)
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

func (c *Client) GetBrandMetricsReport(ctx context.Context, reportId string, params *GetBrandMetricsReportParams) (*http.Response, error) {
	req, err := NewGetBrandMetricsReportRequest(c.Server, reportId, params)
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

// NewGenerateBrandMetricsReportRequestWithApplicationVndInsightsBrandMetricsV1PlusJSONBody calls the generic GenerateBrandMetricsReport builder with application/vnd.insightsBrandMetrics.v1+json body
func NewGenerateBrandMetricsReportRequestWithApplicationVndInsightsBrandMetricsV1PlusJSONBody(server string, params *GenerateBrandMetricsReportParams, body GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGenerateBrandMetricsReportRequestWithBody(server, params, "application/vnd.insightsBrandMetrics.v1+json", bodyReader)
}

// NewGenerateBrandMetricsReportRequestWithBody generates requests for GenerateBrandMetricsReport with any type of body
func NewGenerateBrandMetricsReportRequestWithBody(server string, params *GenerateBrandMetricsReportParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/insights/brandMetrics/report")
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

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-Scope", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIScope)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-Scope", headerParam0)

		var headerParam1 string

		headerParam1, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-ClientID", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIClientID)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-ClientID", headerParam1)

	}

	return req, nil
}

// NewGetBrandMetricsReportRequest generates requests for GetBrandMetricsReport
func NewGetBrandMetricsReportRequest(server string, reportId string, params *GetBrandMetricsReportParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "reportId", runtime.ParamLocationPath, reportId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/insights/brandMetrics/report/%s", pathParam0)
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

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-Scope", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIScope)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-Scope", headerParam0)

		var headerParam1 string

		headerParam1, err = runtime.StyleParamWithLocation("simple", false, "Amazon-Advertising-API-ClientID", runtime.ParamLocationHeader, params.AmazonAdvertisingAPIClientID)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Amazon-Advertising-API-ClientID", headerParam1)

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
	// GenerateBrandMetricsReportWithBodyWithResponse request with any body
	GenerateBrandMetricsReportWithBodyWithResponse(ctx context.Context, params *GenerateBrandMetricsReportParams, contentType string, body io.Reader) (*GenerateBrandMetricsReportResp, error)

	GenerateBrandMetricsReportWithApplicationVndInsightsBrandMetricsV1PlusJSONBodyWithResponse(ctx context.Context, params *GenerateBrandMetricsReportParams, body GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody) (*GenerateBrandMetricsReportResp, error)

	// GetBrandMetricsReportWithResponse request
	GetBrandMetricsReportWithResponse(ctx context.Context, reportId string, params *GetBrandMetricsReportParams) (*GetBrandMetricsReportResp, error)
}

type GenerateBrandMetricsReportResp struct {
	Body                                             []byte
	HTTPResponse                                     *http.Response
	ApplicationvndInsightsBrandMetricsV1JSON200      *BrandMetricsGenerateReportResponse
	ApplicationvndInsightsBrandMetricsErrorV1JSON400 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON401 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON403 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON422 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON424 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON429 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON500 *BrandMetricsError
}

// Status returns HTTPResponse.Status
func (r GenerateBrandMetricsReportResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GenerateBrandMetricsReportResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetBrandMetricsReportResp struct {
	Body                                             []byte
	HTTPResponse                                     *http.Response
	ApplicationvndInsightsBrandMetricsV1JSON200      *BrandMetricsGetReportByIdResponse
	ApplicationvndInsightsBrandMetricsErrorV1JSON400 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON401 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON403 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON422 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON424 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON429 *BrandMetricsError
	ApplicationvndInsightsBrandMetricsErrorV1JSON500 *BrandMetricsError
}

// Status returns HTTPResponse.Status
func (r GetBrandMetricsReportResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBrandMetricsReportResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GenerateBrandMetricsReportWithBodyWithResponse request with arbitrary body returning *GenerateBrandMetricsReportResp
func (c *ClientWithResponses) GenerateBrandMetricsReportWithBodyWithResponse(ctx context.Context, params *GenerateBrandMetricsReportParams, contentType string, body io.Reader) (*GenerateBrandMetricsReportResp, error) {
	rsp, err := c.GenerateBrandMetricsReportWithBody(ctx, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseGenerateBrandMetricsReportResp(rsp)
}

func (c *ClientWithResponses) GenerateBrandMetricsReportWithApplicationVndInsightsBrandMetricsV1PlusJSONBodyWithResponse(ctx context.Context, params *GenerateBrandMetricsReportParams, body GenerateBrandMetricsReportApplicationVndInsightsBrandMetricsV1PlusJSONRequestBody) (*GenerateBrandMetricsReportResp, error) {
	rsp, err := c.GenerateBrandMetricsReportWithApplicationVndInsightsBrandMetricsV1PlusJSONBody(ctx, params, body)
	if err != nil {
		return nil, err
	}
	return ParseGenerateBrandMetricsReportResp(rsp)
}

// GetBrandMetricsReportWithResponse request returning *GetBrandMetricsReportResp
func (c *ClientWithResponses) GetBrandMetricsReportWithResponse(ctx context.Context, reportId string, params *GetBrandMetricsReportParams) (*GetBrandMetricsReportResp, error) {
	rsp, err := c.GetBrandMetricsReport(ctx, reportId, params)
	if err != nil {
		return nil, err
	}
	return ParseGetBrandMetricsReportResp(rsp)
}

// ParseGenerateBrandMetricsReportResp parses an HTTP response from a GenerateBrandMetricsReportWithResponse call
func ParseGenerateBrandMetricsReportResp(rsp *http.Response) (*GenerateBrandMetricsReportResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GenerateBrandMetricsReportResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BrandMetricsGenerateReportResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsV1JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 424:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON424 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON500 = &dest

	}

	return response, nil
}

// ParseGetBrandMetricsReportResp parses an HTTP response from a GetBrandMetricsReportWithResponse call
func ParseGetBrandMetricsReportResp(rsp *http.Response) (*GetBrandMetricsReportResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetBrandMetricsReportResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BrandMetricsGetReportByIdResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsV1JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 424:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON424 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest BrandMetricsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsBrandMetricsErrorV1JSON500 = &dest

	}

	return response, nil
}
