// Package audienceinsights provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package audienceinsights

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// Defines values for InsightsGetAudiencesOverlappingAudiencesParamsAdType.
const (
	DSP InsightsGetAudiencesOverlappingAudiencesParamsAdType = "DSP"
	SD  InsightsGetAudiencesOverlappingAudiencesParamsAdType = "SD"
)

// InsightsAudiencesOverlapAudienceMetadata Information about any audience included in the response.
type InsightsAudiencesOverlapAudienceMetadata struct {
	// AudienceId The identifier of an audience.
	AudienceId string `json:"audienceId"`

	// Category The type of the overlapping audience.
	Category string `json:"category"`

	// ImpressionsForecastRange 30-day impression forecast provides an estimated range of how many times an ad could be displayed using the audience over the next 30 days. The impression forecast is for the individual audience. It can be defined by both upper and lower bound, or just by one of the bounds.
	ImpressionsForecastRange *struct {
		LowerBound *int `json:"lowerBound,omitempty"`
		UpperBound *int `json:"upperBound,omitempty"`
	} `json:"impressionsForecastRange,omitempty"`

	// Name Audience name
	Name string `json:"name"`

	// Size Audience size is a relative measure of how large the overlapping audience is compared to all Amazon audiences. A size of 6 indicates that the audience is larger than 60% of all Amazon audiences.
	Size int `json:"size"`
}

// InsightsAudiencesOverlapAudienceMetadataV2 Information about any audience included in the response.
type InsightsAudiencesOverlapAudienceMetadataV2 struct {
	AudienceForecast *struct {
		// DailyImpressions The available impressions across all inventory types (e.g. display, video, and mobile). This is an estimate and is not guaranteed. It will take at least 2 weeks for the impressions to become available for new audiences. Note: forecasted daily impressions are not available when SD adType is selected.
		DailyImpressions *struct {
			LowerBound *int `json:"lowerBound,omitempty"`
			UpperBound *int `json:"upperBound,omitempty"`
		} `json:"dailyImpressions,omitempty"`

		// DailyReach The unique devices reachable across all inventory types (e.g. display, video, and mobile). This is an estimate and is not guaranteed. It will take at least 2 weeks for the reach to become available for new audiences.
		DailyReach *struct {
			LowerBound *int `json:"lowerBound,omitempty"`
			UpperBound *int `json:"upperBound,omitempty"`
		} `json:"dailyReach,omitempty"`
	} `json:"audienceForecast,omitempty"`

	// AudienceId The identifier of an audience.
	AudienceId string `json:"audienceId"`

	// Category The type of the overlapping audience.
	Category string `json:"category"`

	// Name Audience name
	Name string `json:"name"`
}

// InsightsAudiencesOverlapEntry Contains information about the overlapping audiences and their affinity with the audience requested.
type InsightsAudiencesOverlapEntry struct {
	// Affinity Affinity is a measure of how similar the overlapping audience is to the audience selected to generate the audience overlap insights. An affinity of 2 indicates that the overlapping audience is twice as likely to fall into the audience selected than the average audience on Amazon.
	Affinity *float32 `json:"affinity,omitempty"`

	// AudienceMetadata Information about any audience included in the response.
	AudienceMetadata *InsightsAudiencesOverlapAudienceMetadata `json:"audienceMetadata,omitempty"`
}

// InsightsAudiencesOverlapEntryV2 Contains information about the overlapping audiences and their affinity with the provided audience.
type InsightsAudiencesOverlapEntryV2 struct {
	// Affinity Affinity is a measure of how similar the overlapping audience is to the audience selected to generate the audience overlap insights. An affinity of 2 indicates that the overlapping audience is twice as likely to fall into the audience selected than the average audience on Amazon.
	Affinity *float32 `json:"affinity,omitempty"`

	// AudienceMetadata Information about any audience included in the response.
	AudienceMetadata *InsightsAudiencesOverlapAudienceMetadataV2 `json:"audienceMetadata,omitempty"`
}

// InsightsError The error returned from the server.
type InsightsError struct {
	// Message Detailed information about the error that occurred.
	Message string `json:"message"`

	// RequestId A unique value generated by the server to identify the request.
	RequestId string `json:"requestId"`
}

// InsightsGetAudiencesOverlappingAudiencesParams defines parameters for InsightsGetAudiencesOverlappingAudiences.
type InsightsGetAudiencesOverlappingAudiencesParams struct {
	// AdType The advertising program.
	AdType InsightsGetAudiencesOverlappingAudiencesParamsAdType `form:"adType" json:"adType"`

	// AdvertiserId The identifier of the advertiser you'd like to retrieve overlapping audiences for. This parameter is required for the DSP adType, but is optional for the SD adType.
	AdvertiserId *string `form:"advertiserId,omitempty" json:"advertiserId,omitempty"`

	// MinimumAudienceSize If specified, the sizes of all returned overlapping audiences will be at least the provided size. This parameter is supported only for request to return `application/vnd.insightsaudiencesoverlap.v1+json`.
	MinimumAudienceSize *int `form:"minimumAudienceSize,omitempty" json:"minimumAudienceSize,omitempty"`

	// MaximumAudienceSize If specified, the sizes of all returned overlapping audiences will be at most the provided size. This parameter is supported only for request to return `application/vnd.insightsaudiencesoverlap.v1+json`.
	MaximumAudienceSize *int `form:"maximumAudienceSize,omitempty" json:"maximumAudienceSize,omitempty"`

	// MinimumOverlapAffinity If specified, the affinities of all returned overlapping audiences will be at least the provided affinity.
	MinimumOverlapAffinity *float32 `form:"minimumOverlapAffinity,omitempty" json:"minimumOverlapAffinity,omitempty"`

	// MaximumOverlapAffinity If specified, the affinities of all returned overlapping audiences will be at most the provided affinity.
	MaximumOverlapAffinity *float32 `form:"maximumOverlapAffinity,omitempty" json:"maximumOverlapAffinity,omitempty"`

	// AudienceCategory If specified, the categories of all returned overlapping audiences will be one of the provided categories.
	AudienceCategory *[]string `form:"audienceCategory,omitempty" json:"audienceCategory,omitempty"`

	// MaxResults Sets the maximum number of overlapping audiences in the response. This parameter is supported only for request to return `application/vnd.insightsaudiencesoverlap.v2+json`.
	MaxResults *int `form:"maxResults,omitempty" json:"maxResults,omitempty"`

	// NextToken Token to be used to request additional overlapping audiences. If not provided, the top 30 overlapping audiences are returned. Note: subsequent calls must be made using the same parameters as used in previous requests.
	NextToken *string `form:"nextToken,omitempty" json:"nextToken,omitempty"`

	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`

	// AmazonAdvertisingAPIScope The identifier of a profile associated with the advertiser account. Use `GET` method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header.
	AmazonAdvertisingAPIScope string `json:"Amazon-Advertising-API-Scope"`
}

// InsightsGetAudiencesOverlappingAudiencesParamsAdType defines parameters for InsightsGetAudiencesOverlappingAudiences.
type InsightsGetAudiencesOverlappingAudiencesParamsAdType string

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
	// InsightsGetAudiencesOverlappingAudiences request
	InsightsGetAudiencesOverlappingAudiences(ctx context.Context, audienceId string, params *InsightsGetAudiencesOverlappingAudiencesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) InsightsGetAudiencesOverlappingAudiences(ctx context.Context, audienceId string, params *InsightsGetAudiencesOverlappingAudiencesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewInsightsGetAudiencesOverlappingAudiencesRequest(c.Server, audienceId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewInsightsGetAudiencesOverlappingAudiencesRequest generates requests for InsightsGetAudiencesOverlappingAudiences
func NewInsightsGetAudiencesOverlappingAudiencesRequest(server string, audienceId string, params *InsightsGetAudiencesOverlappingAudiencesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "audienceId", runtime.ParamLocationPath, audienceId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/insights/audiences/%s/overlappingAudiences", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "adType", runtime.ParamLocationQuery, params.AdType); err != nil {
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

		if params.AdvertiserId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "advertiserId", runtime.ParamLocationQuery, *params.AdvertiserId); err != nil {
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

		if params.MinimumAudienceSize != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "minimumAudienceSize", runtime.ParamLocationQuery, *params.MinimumAudienceSize); err != nil {
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

		if params.MaximumAudienceSize != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maximumAudienceSize", runtime.ParamLocationQuery, *params.MaximumAudienceSize); err != nil {
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

		if params.MinimumOverlapAffinity != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "minimumOverlapAffinity", runtime.ParamLocationQuery, *params.MinimumOverlapAffinity); err != nil {
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

		if params.MaximumOverlapAffinity != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maximumOverlapAffinity", runtime.ParamLocationQuery, *params.MaximumOverlapAffinity); err != nil {
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

		if params.AudienceCategory != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "audienceCategory", runtime.ParamLocationQuery, *params.AudienceCategory); err != nil {
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

		if params.MaxResults != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
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

		if params.NextToken != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nextToken", runtime.ParamLocationQuery, *params.NextToken); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
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
	// InsightsGetAudiencesOverlappingAudiencesWithResponse request
	InsightsGetAudiencesOverlappingAudiencesWithResponse(ctx context.Context, audienceId string, params *InsightsGetAudiencesOverlappingAudiencesParams, reqEditors ...RequestEditorFn) (*InsightsGetAudiencesOverlappingAudiencesResp, error)
}

type InsightsGetAudiencesOverlappingAudiencesResp struct {
	Body                                            []byte
	HTTPResponse                                    *http.Response
	ApplicationvndInsightsaudiencesoverlapV1JSON200 *struct {
		// Marketplace The locale used to generate the overlapping audiences.
		Marketplace string `json:"marketplace"`

		// NextToken If present, there are more overlapping audiences than initially returned. Use this token to call the operation again and have the additional overlapping audiences returned. The token is valid for 8 hours from the initial request. Note: subsequent calls must be made using the same parameters as used in previous requests.
		NextToken *string `json:"nextToken,omitempty"`

		// OverlappingAudiences The list of overlapping audiences.
		OverlappingAudiences []InsightsAudiencesOverlapEntry `json:"overlappingAudiences"`

		// RequestedAudienceMetadata Information about any audience included in the response.
		RequestedAudienceMetadata *InsightsAudiencesOverlapAudienceMetadata `json:"requestedAudienceMetadata,omitempty"`
	}
	ApplicationvndInsightsaudiencesoverlapV2JSON200 *struct {
		// Marketplace The locale used to generate the overlapping audiences.
		Marketplace string `json:"marketplace"`

		// NextToken If present, there are more overlapping audiences than initially returned. Use this token to call the operation again and have the additional overlapping audiences returned. The token is valid for 8 hours from the initial request. Note: subsequent calls must be made using the same parameters as used in previous requests.
		NextToken *string `json:"nextToken,omitempty"`

		// OverlappingAudiences The list of overlapping audiences.
		OverlappingAudiences []InsightsAudiencesOverlapEntryV2 `json:"overlappingAudiences"`

		// RequestedAudienceMetadata Information about any audience included in the response.
		RequestedAudienceMetadata InsightsAudiencesOverlapAudienceMetadataV2 `json:"requestedAudienceMetadata"`
	}
	ApplicationvndInsightserrorV1JSON400 *InsightsError
	ApplicationvndInsightserrorV1JSON403 *InsightsError
	ApplicationvndInsightserrorV1JSON404 *InsightsError
}

// Status returns HTTPResponse.Status
func (r InsightsGetAudiencesOverlappingAudiencesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r InsightsGetAudiencesOverlappingAudiencesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// InsightsGetAudiencesOverlappingAudiencesWithResponse request returning *InsightsGetAudiencesOverlappingAudiencesResp
func (c *ClientWithResponses) InsightsGetAudiencesOverlappingAudiencesWithResponse(ctx context.Context, audienceId string, params *InsightsGetAudiencesOverlappingAudiencesParams, reqEditors ...RequestEditorFn) (*InsightsGetAudiencesOverlappingAudiencesResp, error) {
	rsp, err := c.InsightsGetAudiencesOverlappingAudiences(ctx, audienceId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseInsightsGetAudiencesOverlappingAudiencesResp(rsp)
}

// ParseInsightsGetAudiencesOverlappingAudiencesResp parses an HTTP response from a InsightsGetAudiencesOverlappingAudiencesWithResponse call
func ParseInsightsGetAudiencesOverlappingAudiencesResp(rsp *http.Response) (*InsightsGetAudiencesOverlappingAudiencesResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &InsightsGetAudiencesOverlappingAudiencesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.Header.Get("Content-Type") == "application/vnd.insightsaudiencesoverlap.v1+json" && rsp.StatusCode == 200:
		var dest struct {
			// Marketplace The locale used to generate the overlapping audiences.
			Marketplace string `json:"marketplace"`

			// NextToken If present, there are more overlapping audiences than initially returned. Use this token to call the operation again and have the additional overlapping audiences returned. The token is valid for 8 hours from the initial request. Note: subsequent calls must be made using the same parameters as used in previous requests.
			NextToken *string `json:"nextToken,omitempty"`

			// OverlappingAudiences The list of overlapping audiences.
			OverlappingAudiences []InsightsAudiencesOverlapEntry `json:"overlappingAudiences"`

			// RequestedAudienceMetadata Information about any audience included in the response.
			RequestedAudienceMetadata *InsightsAudiencesOverlapAudienceMetadata `json:"requestedAudienceMetadata,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsaudiencesoverlapV1JSON200 = &dest

	case rsp.Header.Get("Content-Type") == "application/vnd.insightsaudiencesoverlap.v2+json" && rsp.StatusCode == 200:
		var dest struct {
			// Marketplace The locale used to generate the overlapping audiences.
			Marketplace string `json:"marketplace"`

			// NextToken If present, there are more overlapping audiences than initially returned. Use this token to call the operation again and have the additional overlapping audiences returned. The token is valid for 8 hours from the initial request. Note: subsequent calls must be made using the same parameters as used in previous requests.
			NextToken *string `json:"nextToken,omitempty"`

			// OverlappingAudiences The list of overlapping audiences.
			OverlappingAudiences []InsightsAudiencesOverlapEntryV2 `json:"overlappingAudiences"`

			// RequestedAudienceMetadata Information about any audience included in the response.
			RequestedAudienceMetadata InsightsAudiencesOverlapAudienceMetadataV2 `json:"requestedAudienceMetadata"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightsaudiencesoverlapV2JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest InsightsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightserrorV1JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest InsightsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightserrorV1JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest InsightsError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationvndInsightserrorV1JSON404 = &dest

	}

	return response, nil
}
