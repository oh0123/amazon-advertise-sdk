// Package moderation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package moderation

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
)

// Defines values for IdType.
const (
	ADID IdType = "AD_ID"
)

// Defines values for ModerationResultsAccessDeniedErrorCode.
const (
	ACCESSDENIED ModerationResultsAccessDeniedErrorCode = "ACCESS_DENIED"
)

// Defines values for ModerationResultsAdProgramType.
const (
	SBPRODUCTCOLLECTION ModerationResultsAdProgramType = "SB_PRODUCT_COLLECTION"
	SBSTORESPOTLIGHT    ModerationResultsAdProgramType = "SB_STORE_SPOTLIGHT"
	SBVIDEO             ModerationResultsAdProgramType = "SB_VIDEO"
	SPONSOREDPRODUCTS   ModerationResultsAdProgramType = "SPONSORED_PRODUCTS"
)

// Defines values for ModerationResultsBadRequestErrorCode.
const (
	BADREQUEST ModerationResultsBadRequestErrorCode = "BAD_REQUEST"
)

// Defines values for ModerationResultsInternalServerErrorCode.
const (
	INTERNALERROR ModerationResultsInternalServerErrorCode = "INTERNAL_ERROR"
)

// Defines values for ModerationResultsNotFoundErrorCode.
const (
	NOTFOUND ModerationResultsNotFoundErrorCode = "NOT_FOUND"
)

// Defines values for ModerationResultsThrottlingErrorCode.
const (
	THROTTLED ModerationResultsThrottlingErrorCode = "THROTTLED"
)

// Defines values for ModerationStatus.
const (
	APPROVED   ModerationStatus = "APPROVED"
	FAILED     ModerationStatus = "FAILED"
	INPROGRESS ModerationStatus = "IN_PROGRESS"
	REJECTED   ModerationStatus = "REJECTED"
)

// Id The unique identifier of the ad which can be obtained after the ad is created using create APIs.
type Id = string

// IdType The unique identifiers type based on the adProgram of the ad.
type IdType string

// ImageCrop defines model for ImageCrop.
type ImageCrop struct {
	// Height Policy violated region's height in pixel.
	Height *int64 `json:"height,omitempty"`

	// TopLeftX Policy violated region's top left X-axis pixel value.
	TopLeftX *int64 `json:"topLeftX,omitempty"`

	// TopLeftY Policy violated region's top left Y-axis pixel value.
	TopLeftY *int64 `json:"topLeftY,omitempty"`

	// Width Policy violated region's width in pixel.
	Width *int64 `json:"width,omitempty"`
}

// ModerationResult defines model for ModerationResult.
type ModerationResult struct {
	// EtaForModeration Expected date and time by which moderation will be complete. The format is ISO 8601 in UTC time zone. Note that this field is present in the response only when moderationStatus is set to IN_PROGRESS.
	EtaForModeration *string `json:"etaForModeration,omitempty"`

	// Id The unique identifier of the ad which can be obtained after the ad is created using create APIs.
	Id *Id `json:"id,omitempty"`

	// IdType The unique identifiers type based on the adProgram of the ad.
	IdType *IdType `json:"idType,omitempty"`

	// ModerationStatus The moderation status of the ad.
	ModerationStatus *ModerationStatus `json:"moderationStatus,omitempty"`

	// PolicyViolations A list of policy violations for a campaign that has failed moderation. Note that this field is present in the response only when moderationStatus is set to REJECTED.
	PolicyViolations *[]PolicyViolation `json:"policyViolations,omitempty"`

	// VersionId The version identifier that helps to keep track of multiple versions of a submitted ad. In case of Sponsored Brands this is the creative version id.
	VersionId *VersionId `json:"versionId,omitempty"`
}

// ModerationResultsAccessDeniedError defines model for ModerationResultsAccessDeniedError.
type ModerationResultsAccessDeniedError struct {
	// Code Access denied error code.
	Code *ModerationResultsAccessDeniedErrorCode `json:"code,omitempty"`

	// Details A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// ModerationResultsAccessDeniedErrorCode Access denied error code.
type ModerationResultsAccessDeniedErrorCode string

// ModerationResultsAdProgramType The program type of the ad.
type ModerationResultsAdProgramType string

// ModerationResultsBadRequestError defines model for ModerationResultsBadRequestError.
type ModerationResultsBadRequestError struct {
	// Code Bad request error code.
	Code *ModerationResultsBadRequestErrorCode `json:"code,omitempty"`

	// Details A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// ModerationResultsBadRequestErrorCode Bad request error code.
type ModerationResultsBadRequestErrorCode string

// ModerationResultsInternalServerError defines model for ModerationResultsInternalServerError.
type ModerationResultsInternalServerError struct {
	// Code Internal error code.
	Code *ModerationResultsInternalServerErrorCode `json:"code,omitempty"`

	// Details A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// ModerationResultsInternalServerErrorCode Internal error code.
type ModerationResultsInternalServerErrorCode string

// ModerationResultsNotFoundError defines model for ModerationResultsNotFoundError.
type ModerationResultsNotFoundError struct {
	// Code Not found error code.
	Code *ModerationResultsNotFoundErrorCode `json:"code,omitempty"`

	// Details A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// ModerationResultsNotFoundErrorCode Not found error code.
type ModerationResultsNotFoundErrorCode string

// ModerationResultsRequest defines model for ModerationResultsRequest.
type ModerationResultsRequest struct {
	// AdProgramType The program type of the ad.
	AdProgramType ModerationResultsAdProgramType `json:"adProgramType"`

	// Id The unique identifier of the ad which can be obtained after the ad is created using create APIs.
	Id Id `json:"id"`

	// IdType The unique identifiers type based on the adProgram of the ad.
	IdType IdType `json:"idType"`

	// MaxResults Sets a limit on the number of results returned by an operation.
	MaxResults int32 `json:"maxResults"`

	// ModerationStatusFilter Filter by specific moderation status.
	ModerationStatusFilter *[]ModerationStatus `json:"moderationStatusFilter,omitempty"`

	// NextToken Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *NextToken `json:"nextToken,omitempty"`

	// VersionIdFilter Filter by specific version id of the ad. The API will return the ad's all versions moderation status if this field is empty.
	VersionIdFilter *[]VersionId `json:"versionIdFilter,omitempty"`
}

// ModerationResultsResponse defines model for ModerationResultsResponse.
type ModerationResultsResponse struct {
	ModerationResults *[]ModerationResult `json:"moderationResults,omitempty"`

	// NextToken Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *NextToken `json:"nextToken,omitempty"`
}

// ModerationResultsThrottlingError defines model for ModerationResultsThrottlingError.
type ModerationResultsThrottlingError struct {
	// Code Throttled error code.
	Code *ModerationResultsThrottlingErrorCode `json:"code,omitempty"`

	// Details A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// ModerationResultsThrottlingErrorCode Throttled error code.
type ModerationResultsThrottlingErrorCode string

// ModerationStatus The moderation status of the ad.
type ModerationStatus string

// NextToken Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
type NextToken = string

// PolicyViolation defines model for PolicyViolation.
type PolicyViolation struct {
	// PolicyDescription A human-readable description of the policy.
	PolicyDescription *string `json:"policyDescription,omitempty"`

	// PolicyLinkUrl Address of the policy documentation. Follow the link to learn more about the specified policy.
	PolicyLinkUrl          *string                  `json:"policyLinkUrl,omitempty"`
	ViolatingAsinContents  *[]ViolatingAsinContent  `json:"violatingAsinContents,omitempty"`
	ViolatingImageContents *[]ViolatingImageContent `json:"violatingImageContents,omitempty"`
	ViolatingTextContents  *[]ViolatingTextContent  `json:"violatingTextContents,omitempty"`
	ViolatingVideoContents *[]ViolatingVideoContent `json:"violatingVideoContents,omitempty"`
}

// TextPosition defines model for TextPosition.
type TextPosition struct {
	// End Zero-based index into the text in reviewedText where the text specified in violatingText ends.
	End *int64 `json:"end,omitempty"`

	// Start Zero-based index into the text in reviewedText where the text specified in violatingText starts.
	Start *int64 `json:"start,omitempty"`
}

// VersionId The version identifier that helps to keep track of multiple versions of a submitted ad. In case of Sponsored Brands this is the creative version id.
type VersionId = string

// VideoPosition defines model for VideoPosition.
type VideoPosition struct {
	// End End time of the video having the policy violation.
	End *int64 `json:"end,omitempty"`

	// Start Start time of the video having the policy violation.
	Start *int64 `json:"start,omitempty"`
}

// ViolatingAsinContent defines model for ViolatingAsinContent.
type ViolatingAsinContent struct {
	// ModeratedComponent Moderation component which marked the policy violation.
	ModeratedComponent     *string                  `json:"moderatedComponent,omitempty"`
	ViolatingAsinEvidences *[]ViolatingAsinEvidence `json:"violatingAsinEvidences,omitempty"`
}

// ViolatingAsinEvidence defines model for ViolatingAsinEvidence.
type ViolatingAsinEvidence struct {
	// Asin ASIN which has the ad policy violation.
	Asin *string `json:"asin,omitempty"`
}

// ViolatingImageContent defines model for ViolatingImageContent.
type ViolatingImageContent struct {
	// ModeratedComponent Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`

	// ReviewedImageUrl URL of the image which has the ad policy violation.
	ReviewedImageUrl        *string                   `json:"reviewedImageUrl,omitempty"`
	ViolatingImageEvidences *[]ViolatingImageEvidence `json:"violatingImageEvidences,omitempty"`
}

// ViolatingImageEvidence defines model for ViolatingImageEvidence.
type ViolatingImageEvidence struct {
	ViolatingImageCrop *ImageCrop `json:"violatingImageCrop,omitempty"`
}

// ViolatingTextContent Information about the specific text that violates the specified policy in the campaign.
type ViolatingTextContent struct {
	// ModeratedComponent Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`

	// ReviewedText The actual text on which the moderation was done.
	ReviewedText           *string                  `json:"reviewedText,omitempty"`
	ViolatingTextEvidences *[]ViolatingTextEvidence `json:"violatingTextEvidences,omitempty"`
}

// ViolatingTextEvidence defines model for ViolatingTextEvidence.
type ViolatingTextEvidence struct {
	// ViolatingText The specific text determined to violate the specified policy in reviewedText.
	ViolatingText         *string       `json:"violatingText,omitempty"`
	ViolatingTextPosition *TextPosition `json:"violatingTextPosition,omitempty"`
}

// ViolatingVideoContent defines model for ViolatingVideoContent.
type ViolatingVideoContent struct {
	// ModeratedComponent Moderation component which marked the policy violation.
	ModeratedComponent *string `json:"moderatedComponent,omitempty"`

	// ReviewedVideoUrl URL of the video which has the ad policy violation.
	ReviewedVideoUrl        *string                   `json:"reviewedVideoUrl,omitempty"`
	ViolatingVideoEvidences *[]ViolatingVideoEvidence `json:"violatingVideoEvidences,omitempty"`
}

// ViolatingVideoEvidence defines model for ViolatingVideoEvidence.
type ViolatingVideoEvidence struct {
	ViolatingVideoPosition *VideoPosition `json:"violatingVideoPosition,omitempty"`
}

// ClientHeader defines model for clientHeader.
type ClientHeader = string

// ProfileHeader defines model for profileHeader.
type ProfileHeader = string

// ModerationResultsParams defines parameters for ModerationResults.
type ModerationResultsParams struct {
	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId ClientHeader `json:"Amazon-Advertising-API-ClientId"`

	// AmazonAdvertisingAPIScope The identifier of a profile associated with the advertiser account. Use `GET` method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id `profileId` from the response to pass it as input.
	AmazonAdvertisingAPIScope ProfileHeader `json:"Amazon-Advertising-API-Scope"`
}

// ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody defines body for ModerationResults for application/vnd.moderationresultsrequest.v4.1+json ContentType.
type ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody = ModerationResultsRequest

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
	// ModerationResultsWithBody request with any body
	ModerationResultsWithBody(ctx context.Context, params *ModerationResultsParams, contentType string, body io.Reader) (*http.Response, error)

	ModerationResultsWithApplicationVndModerationresultsrequestV41PlusJSONBody(ctx context.Context, params *ModerationResultsParams, body ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody) (*http.Response, error)
}

func (c *Client) ModerationResultsWithBody(ctx context.Context, params *ModerationResultsParams, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewModerationResultsRequestWithBody(c.Server, params, contentType, body)
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

func (c *Client) ModerationResultsWithApplicationVndModerationresultsrequestV41PlusJSONBody(ctx context.Context, params *ModerationResultsParams, body ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody) (*http.Response, error) {
	req, err := NewModerationResultsRequestWithApplicationVndModerationresultsrequestV41PlusJSONBody(c.Server, params, body)
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

// NewModerationResultsRequestWithApplicationVndModerationresultsrequestV41PlusJSONBody calls the generic ModerationResults builder with application/vnd.moderationresultsrequest.v4.1+json body
func NewModerationResultsRequestWithApplicationVndModerationresultsrequestV41PlusJSONBody(server string, params *ModerationResultsParams, body ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewModerationResultsRequestWithBody(server, params, "application/vnd.moderationresultsrequest.v4.1+json", bodyReader)
}

// NewModerationResultsRequestWithBody generates requests for ModerationResults with any type of body
func NewModerationResultsRequestWithBody(server string, params *ModerationResultsParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/moderation/results")
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
	// ModerationResultsWithBodyWithResponse request with any body
	ModerationResultsWithBodyWithResponse(ctx context.Context, params *ModerationResultsParams, contentType string, body io.Reader) (*ModerationResultsResp, error)

	ModerationResultsWithApplicationVndModerationresultsrequestV41PlusJSONBodyWithResponse(ctx context.Context, params *ModerationResultsParams, body ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody) (*ModerationResultsResp, error)
}

type ModerationResultsResp struct {
	Body                                                         []byte
	HTTPResponse                                                 *http.Response
	ApplicationvndModerationresultsresponseV40JSON200            *ModerationResultsResponse
	ApplicationvndModerationresultsbadrequesterrorV40JSON400     *ModerationResultsBadRequestError
	ApplicationvndModerationresultsaccessdeniederrorV40JSON403   *ModerationResultsAccessDeniedError
	ApplicationvndModerationresultsnotfounderrorV40JSON404       *ModerationResultsNotFoundError
	ApplicationvndModerationresultsthrottlingerrorV40JSON429     *ModerationResultsThrottlingError
	ApplicationvndModerationresultsinternalservererrorV40JSON500 *ModerationResultsInternalServerError
}

// Status returns HTTPResponse.Status
func (r ModerationResultsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ModerationResultsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ModerationResultsWithBodyWithResponse request with arbitrary body returning *ModerationResultsResp
func (c *ClientWithResponses) ModerationResultsWithBodyWithResponse(ctx context.Context, params *ModerationResultsParams, contentType string, body io.Reader) (*ModerationResultsResp, error) {
	rsp, err := c.ModerationResultsWithBody(ctx, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseModerationResultsResp(rsp)
}

func (c *ClientWithResponses) ModerationResultsWithApplicationVndModerationresultsrequestV41PlusJSONBodyWithResponse(ctx context.Context, params *ModerationResultsParams, body ModerationResultsApplicationVndModerationresultsrequestV41PlusJSONRequestBody) (*ModerationResultsResp, error) {
	rsp, err := c.ModerationResultsWithApplicationVndModerationresultsrequestV41PlusJSONBody(ctx, params, body)
	if err != nil {
		return nil, err
	}
	return ParseModerationResultsResp(rsp)
}

// ParseModerationResultsResp parses an HTTP response from a ModerationResultsWithResponse call
func ParseModerationResultsResp(rsp *http.Response) (*ModerationResultsResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ModerationResultsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ModerationResultsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsresponseV40JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ModerationResultsBadRequestError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsbadrequesterrorV40JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ModerationResultsAccessDeniedError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsaccessdeniederrorV40JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ModerationResultsNotFoundError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsnotfounderrorV40JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ModerationResultsThrottlingError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsthrottlingerrorV40JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ModerationResultsInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndModerationresultsinternalservererrorV40JSON500 = &dest

	}

	return response, nil
}
