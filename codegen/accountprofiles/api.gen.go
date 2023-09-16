// Package accountprofiles provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package accountprofiles

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

const (
	BearerAuthScopes              = "bearerAuth.Scopes"
	Oauth2AuthorizationCodeScopes = "oauth2AuthorizationCode.Scopes"
)

// Defines values for AccountInfoSubType.
const (
	AMAZONATTRIBUTION AccountInfoSubType = "AMAZON_ATTRIBUTION"
	KDPAUTHOR         AccountInfoSubType = "KDP_AUTHOR"
)

// Defines values for AccountType.
const (
	AccountTypeAgency AccountType = "agency"
	AccountTypeSeller AccountType = "seller"
	AccountTypeVendor AccountType = "vendor"
)

// Defines values for ProfileCurrencyCode.
const (
	AED ProfileCurrencyCode = "AED"
	AUD ProfileCurrencyCode = "AUD"
	BRL ProfileCurrencyCode = "BRL"
	CAD ProfileCurrencyCode = "CAD"
	EGP ProfileCurrencyCode = "EGP"
	EUR ProfileCurrencyCode = "EUR"
	GBP ProfileCurrencyCode = "GBP"
	INR ProfileCurrencyCode = "INR"
	JPY ProfileCurrencyCode = "JPY"
	MXN ProfileCurrencyCode = "MXN"
	PLN ProfileCurrencyCode = "PLN"
	SAR ProfileCurrencyCode = "SAR"
	SEK ProfileCurrencyCode = "SEK"
	SGD ProfileCurrencyCode = "SGD"
	TRY ProfileCurrencyCode = "TRY"
	USD ProfileCurrencyCode = "USD"
)

// Defines values for ProfileTimezone.
const (
	AfricaCairo       ProfileTimezone = "Africa/Cairo"
	AmericaLosAngeles ProfileTimezone = "America/Los_Angeles"
	AmericaSaoPaulo   ProfileTimezone = "America/Sao_Paulo"
	AsiaDubai         ProfileTimezone = "Asia/Dubai"
	AsiaKolkata       ProfileTimezone = "Asia/Kolkata"
	AsiaRiyadh        ProfileTimezone = "Asia/Riyadh"
	AsiaSingapore     ProfileTimezone = "Asia/Singapore"
	AsiaTokyo         ProfileTimezone = "Asia/Tokyo"
	AustraliaSydney   ProfileTimezone = "Australia/Sydney"
	EuropeAmsterdam   ProfileTimezone = "Europe/Amsterdam"
	EuropeIstanbul    ProfileTimezone = "Europe/Istanbul"
	EuropeLondon      ProfileTimezone = "Europe/London"
	EuropeParis       ProfileTimezone = "Europe/Paris"
	EuropeStockholm   ProfileTimezone = "Europe/Stockholm"
	EuropeWarsaw      ProfileTimezone = "Europe/Warsaw"
)

// Defines values for CountryCode.
const (
	AE CountryCode = "AE"
	AU CountryCode = "AU"
	BE CountryCode = "BE"
	BR CountryCode = "BR"
	CA CountryCode = "CA"
	DE CountryCode = "DE"
	EG CountryCode = "EG"
	ES CountryCode = "ES"
	FR CountryCode = "FR"
	IN CountryCode = "IN"
	IT CountryCode = "IT"
	JP CountryCode = "JP"
	MX CountryCode = "MX"
	NL CountryCode = "NL"
	PL CountryCode = "PL"
	SA CountryCode = "SA"
	SE CountryCode = "SE"
	SG CountryCode = "SG"
	TR CountryCode = "TR"
	UK CountryCode = "UK"
	US CountryCode = "US"
)

// Defines values for ListProfilesParamsApiProgram.
const (
	Account       ListProfilesParamsApiProgram = "account"
	Billing       ListProfilesParamsApiProgram = "billing"
	Campaign      ListProfilesParamsApiProgram = "campaign"
	PaymentMethod ListProfilesParamsApiProgram = "paymentMethod"
	Posts         ListProfilesParamsApiProgram = "posts"
	Report        ListProfilesParamsApiProgram = "report"
	Store         ListProfilesParamsApiProgram = "store"
)

// Defines values for ListProfilesParamsAccessLevel.
const (
	Edit ListProfilesParamsAccessLevel = "edit"
	View ListProfilesParamsAccessLevel = "view"
)

// Defines values for ListProfilesParamsProfileTypeFilter.
const (
	ListProfilesParamsProfileTypeFilterAgency ListProfilesParamsProfileTypeFilter = "agency"
	ListProfilesParamsProfileTypeFilterSeller ListProfilesParamsProfileTypeFilter = "seller"
	ListProfilesParamsProfileTypeFilterVendor ListProfilesParamsProfileTypeFilter = "vendor"
)

// Defines values for ListProfilesParamsValidPaymentMethodFilter.
const (
	False ListProfilesParamsValidPaymentMethodFilter = "false"
	True  ListProfilesParamsValidPaymentMethodFilter = "true"
)

// AccountInfo defines model for AccountInfo.
type AccountInfo struct {
	// Id Identifier for sellers and vendors. Note that this value is not unique and may be the same across marketplace.
	Id *string `json:"id,omitempty"`

	// MarketplaceStringId The identifier of the marketplace to which the account is associated.
	MarketplaceStringId *string `json:"marketplaceStringId,omitempty"`

	// Name Account name.
	Name *string `json:"name,omitempty"`

	// SubType The account subtype.
	SubType *AccountInfoSubType `json:"subType,omitempty"`

	// Type The `seller` and `vendor` account types are associated with Sponsored Ads APIs. The `agency` account type is associated with DSP and Data Provider APIs.
	Type *AccountType `json:"type,omitempty"`

	// ValidPaymentMethod Only present for Vendors, this returns whether the Advertiser has set up a valid payment method or not.
	ValidPaymentMethod *bool `json:"validPaymentMethod,omitempty"`
}

// AccountInfoSubType The account subtype.
type AccountInfoSubType string

// AccountType The `seller` and `vendor` account types are associated with Sponsored Ads APIs. The `agency` account type is associated with DSP and Data Provider APIs.
type AccountType string

// Profile defines model for Profile.
type Profile struct {
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`

	// CountryCode The countryCode for a given country
	// |Region|`countryCode`|Country Name|
	// |------|-----|-------|
	// |NA|BR|Brazil|
	// |NA|CA|Canada|
	// |NA|MX|Mexico|
	// |NA|US|United States|
	// |EU|AE|United Arab Emirates|
	// |EU|BE|Belgium|
	// |EU|DE|Germany|
	// |EU|EG|Egypt|
	// |EU|ES|Spain|
	// |EU|FR|France|
	// |EU|IN|India|
	// |EU|IT|Italy|
	// |EU|NL|The Netherlands|
	// |EU|PL|Poland|
	// |EU|SA|Saudi Arabia|
	// |EU|SE|Sweden|
	// |EU|TR|Turkey|
	// |EU|UK|United Kingdom|
	// |FE|AU|Australia|
	// |FE|JP|Japan|
	// |FE|SG|Singapore|
	CountryCode *CountryCode `json:"countryCode,omitempty"`

	// CurrencyCode The currency used for all monetary values for entities under this profile.
	// |Region|`countryCode`|Country Name|`currencyCode`|
	// |-----|------|------|------|
	// |NA|BR|Brazil|BRL|
	// |NA|CA|Canada|CAD|
	// |NA|MX|Mexico|MXN|
	// |NA|US|United States|USD|
	// |EU|AE|United Arab Emirates|AED|
	// |EU|BE|Belgium|EUR|
	// |EU|DE|Germany|EUR|
	// |EU|EG|Egypt|EGP|
	// |EU|ES|Spain|EUR|
	// |EU|FR|France|EUR|
	// |EU|IN|India|INR|
	// |EU|IT|Italy|EUR|
	// |EU|NL|The Netherlands|EUR|
	// |EU|PL|Poland|PLN|
	// |EU|SA|Saudi Arabia|SAR|
	// |EU|SE|Sweden|SEK|
	// |EU|TR|Turkey|TRY|
	// |EU|UK|United Kingdom|GBP|
	// |FE|AU|Australia|AUD|
	// |FE|JP|Japan|JPY|
	// |FE|SG|Singapore|SGD|
	CurrencyCode *ProfileCurrencyCode `json:"currencyCode,omitempty"`

	// DailyBudget Note that this field applies to Sponsored Product campaigns for seller type accounts only. Not supported for vendor type accounts.
	DailyBudget *float32 `json:"dailyBudget,omitempty"`
	ProfileId   *int64   `json:"profileId,omitempty"`

	// Timezone The time zone used for all date-based campaign management and reporting.
	// |Region|`countryCode`|Country Name|`timezone`|
	// |------|-----|-----|------|
	// |NA|BR|Brazil|America/Sao_Paulo|
	// |NA|CA|Canada|America/Los_Angeles|
	// |NA|MX|Mexico|America/Los_Angeles|
	// |NA|US|United States|America/Los_Angeles|
	// |EU|AE|United Arab Emirates|Asia/Dubai|
	// |EU|BE|Belgium|Europe/Paris|
	// |EU|DE|Germany|Europe/Paris|
	// |EU|EG|Egypt|Africa/Cairo|
	// |EU|ES|Spain|Europe/Paris|
	// |EU|FR|France|Europe/Paris|
	// |EU|IN|India|Asia/Kolkata|
	// |EU|IT|Italy|Europe/Paris|
	// |EU|NL|The Netherlands|Europe/Amsterdam|
	// |EU|PL|Poland|Europe/Warsaw|
	// |EU|SA|Saudi Arabia|Asia/Riyadh|
	// |EU|SE|Sweden|Europe/Stockholm|
	// |EU|TR|Turkey|Europe/Istanbul|
	// |EU|UK|United Kingdom|Europe/London|
	// |FE|AU|Australia|Australia/Sydney|
	// |FE|JP|Japan|Asia/Tokyo|
	// |FE|SG|Singapore|Asia/Singapore|
	Timezone *ProfileTimezone `json:"timezone,omitempty"`
}

// ProfileCurrencyCode The currency used for all monetary values for entities under this profile.
// |Region|`countryCode`|Country Name|`currencyCode`|
// |-----|------|------|------|
// |NA|BR|Brazil|BRL|
// |NA|CA|Canada|CAD|
// |NA|MX|Mexico|MXN|
// |NA|US|United States|USD|
// |EU|AE|United Arab Emirates|AED|
// |EU|BE|Belgium|EUR|
// |EU|DE|Germany|EUR|
// |EU|EG|Egypt|EGP|
// |EU|ES|Spain|EUR|
// |EU|FR|France|EUR|
// |EU|IN|India|INR|
// |EU|IT|Italy|EUR|
// |EU|NL|The Netherlands|EUR|
// |EU|PL|Poland|PLN|
// |EU|SA|Saudi Arabia|SAR|
// |EU|SE|Sweden|SEK|
// |EU|TR|Turkey|TRY|
// |EU|UK|United Kingdom|GBP|
// |FE|AU|Australia|AUD|
// |FE|JP|Japan|JPY|
// |FE|SG|Singapore|SGD|
type ProfileCurrencyCode string

// ProfileTimezone The time zone used for all date-based campaign management and reporting.
// |Region|`countryCode`|Country Name|`timezone`|
// |------|-----|-----|------|
// |NA|BR|Brazil|America/Sao_Paulo|
// |NA|CA|Canada|America/Los_Angeles|
// |NA|MX|Mexico|America/Los_Angeles|
// |NA|US|United States|America/Los_Angeles|
// |EU|AE|United Arab Emirates|Asia/Dubai|
// |EU|BE|Belgium|Europe/Paris|
// |EU|DE|Germany|Europe/Paris|
// |EU|EG|Egypt|Africa/Cairo|
// |EU|ES|Spain|Europe/Paris|
// |EU|FR|France|Europe/Paris|
// |EU|IN|India|Asia/Kolkata|
// |EU|IT|Italy|Europe/Paris|
// |EU|NL|The Netherlands|Europe/Amsterdam|
// |EU|PL|Poland|Europe/Warsaw|
// |EU|SA|Saudi Arabia|Asia/Riyadh|
// |EU|SE|Sweden|Europe/Stockholm|
// |EU|TR|Turkey|Europe/Istanbul|
// |EU|UK|United Kingdom|Europe/London|
// |FE|AU|Australia|Australia/Sydney|
// |FE|JP|Japan|Asia/Tokyo|
// |FE|SG|Singapore|Asia/Singapore|
type ProfileTimezone string

// ProfileResponse defines model for ProfileResponse.
type ProfileResponse struct {
	Code      *string `json:"code,omitempty"`
	Details   *string `json:"details,omitempty"`
	ProfileId *int64  `json:"profileId,omitempty"`
}

// CountryCode The countryCode for a given country
// |Region|`countryCode`|Country Name|
// |------|-----|-------|
// |NA|BR|Brazil|
// |NA|CA|Canada|
// |NA|MX|Mexico|
// |NA|US|United States|
// |EU|AE|United Arab Emirates|
// |EU|BE|Belgium|
// |EU|DE|Germany|
// |EU|EG|Egypt|
// |EU|ES|Spain|
// |EU|FR|France|
// |EU|IN|India|
// |EU|IT|Italy|
// |EU|NL|The Netherlands|
// |EU|PL|Poland|
// |EU|SA|Saudi Arabia|
// |EU|SE|Sweden|
// |EU|TR|Turkey|
// |EU|UK|United Kingdom|
// |FE|AU|Australia|
// |FE|JP|Japan|
// |FE|SG|Singapore|
type CountryCode string

// ListProfilesParams defines parameters for ListProfiles.
type ListProfilesParams struct {
	// ApiProgram Filters response to include profiles that have permissions for the specified Advertising API program only. Setting `apiProgram=billing` filters the response to include only profiles to which the user and application associated with the access token have permission to view or edit billing information.
	ApiProgram *ListProfilesParamsApiProgram `form:"apiProgram,omitempty" json:"apiProgram,omitempty"`

	// AccessLevel Filters response to include profiles that have specified permissions for the specified Advertising API program only. Currently, the only supported access level is `view` and `edit`. Setting `accessLevel=view` filters the response to include only profiles to which the user and application associated with the access token have view permission to the provided api program.
	AccessLevel *ListProfilesParamsAccessLevel `form:"accessLevel,omitempty" json:"accessLevel,omitempty"`

	// ProfileTypeFilter Filters response to include profiles that are of the specified types in the comma-delimited list. Default is all types. Note that this filter performs an inclusive AND operation on the types.
	ProfileTypeFilter *ListProfilesParamsProfileTypeFilter `form:"profileTypeFilter,omitempty" json:"profileTypeFilter,omitempty"`

	// ValidPaymentMethodFilter Filter response to include profiles that have valid payment methods. Default is to include all profiles. Setting this filter to `true` returns only profiles with either no `validPaymentMethod` field, or the `validPaymentMethod` field set to `true`.  Setting this to `false` returns profiles with the `validPaymentMethod` field set to `false` only.
	ValidPaymentMethodFilter *ListProfilesParamsValidPaymentMethodFilter `form:"validPaymentMethodFilter,omitempty" json:"validPaymentMethodFilter,omitempty"`

	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`
}

// ListProfilesParamsApiProgram defines parameters for ListProfiles.
type ListProfilesParamsApiProgram string

// ListProfilesParamsAccessLevel defines parameters for ListProfiles.
type ListProfilesParamsAccessLevel string

// ListProfilesParamsProfileTypeFilter defines parameters for ListProfiles.
type ListProfilesParamsProfileTypeFilter string

// ListProfilesParamsValidPaymentMethodFilter defines parameters for ListProfiles.
type ListProfilesParamsValidPaymentMethodFilter string

// UpdateProfilesJSONBody defines parameters for UpdateProfiles.
type UpdateProfilesJSONBody = []Profile

// UpdateProfilesParams defines parameters for UpdateProfiles.
type UpdateProfilesParams struct {
	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`
}

// GetProfileByIdParams defines parameters for GetProfileById.
type GetProfileByIdParams struct {
	// AmazonAdvertisingAPIClientId The identifier of a client associated with a "Login with Amazon" account.
	AmazonAdvertisingAPIClientId string `json:"Amazon-Advertising-API-ClientId"`
}

// UpdateProfilesJSONRequestBody defines body for UpdateProfiles for application/json ContentType.
type UpdateProfilesJSONRequestBody = UpdateProfilesJSONBody

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
	// ListProfiles request
	ListProfiles(ctx context.Context, params *ListProfilesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateProfilesWithBody request with any body
	UpdateProfilesWithBody(ctx context.Context, params *UpdateProfilesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateProfiles(ctx context.Context, params *UpdateProfilesParams, body UpdateProfilesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProfileById request
	GetProfileById(ctx context.Context, profileId int64, params *GetProfileByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListProfiles(ctx context.Context, params *ListProfilesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListProfilesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateProfilesWithBody(ctx context.Context, params *UpdateProfilesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateProfilesRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateProfiles(ctx context.Context, params *UpdateProfilesParams, body UpdateProfilesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateProfilesRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProfileById(ctx context.Context, profileId int64, params *GetProfileByIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProfileByIdRequest(c.Server, profileId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListProfilesRequest generates requests for ListProfiles
func NewListProfilesRequest(server string, params *ListProfilesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/profiles")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ApiProgram != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "apiProgram", runtime.ParamLocationQuery, *params.ApiProgram); err != nil {
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

		if params.AccessLevel != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "accessLevel", runtime.ParamLocationQuery, *params.AccessLevel); err != nil {
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

		if params.ProfileTypeFilter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "profileTypeFilter", runtime.ParamLocationQuery, *params.ProfileTypeFilter); err != nil {
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

		if params.ValidPaymentMethodFilter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "validPaymentMethodFilter", runtime.ParamLocationQuery, *params.ValidPaymentMethodFilter); err != nil {
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

	}

	return req, nil
}

// NewUpdateProfilesRequest calls the generic UpdateProfiles builder with application/json body
func NewUpdateProfilesRequest(server string, params *UpdateProfilesParams, body UpdateProfilesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateProfilesRequestWithBody(server, params, "application/json", bodyReader)
}

// NewUpdateProfilesRequestWithBody generates requests for UpdateProfiles with any type of body
func NewUpdateProfilesRequestWithBody(server string, params *UpdateProfilesParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/profiles")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
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

// NewGetProfileByIdRequest generates requests for GetProfileById
func NewGetProfileByIdRequest(server string, profileId int64, params *GetProfileByIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "profileId", runtime.ParamLocationPath, profileId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v2/profiles/%s", pathParam0)
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
	// ListProfilesWithResponse request
	ListProfilesWithResponse(ctx context.Context, params *ListProfilesParams, reqEditors ...RequestEditorFn) (*ListProfilesResp, error)

	// UpdateProfilesWithBodyWithResponse request with any body
	UpdateProfilesWithBodyWithResponse(ctx context.Context, params *UpdateProfilesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateProfilesResp, error)

	UpdateProfilesWithResponse(ctx context.Context, params *UpdateProfilesParams, body UpdateProfilesJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateProfilesResp, error)

	// GetProfileByIdWithResponse request
	GetProfileByIdWithResponse(ctx context.Context, profileId int64, params *GetProfileByIdParams, reqEditors ...RequestEditorFn) (*GetProfileByIdResp, error)
}

type ListProfilesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Profile
}

// Status returns HTTPResponse.Status
func (r ListProfilesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListProfilesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateProfilesResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]ProfileResponse
}

// Status returns HTTPResponse.Status
func (r UpdateProfilesResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateProfilesResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProfileByIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Profile
}

// Status returns HTTPResponse.Status
func (r GetProfileByIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProfileByIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListProfilesWithResponse request returning *ListProfilesResp
func (c *ClientWithResponses) ListProfilesWithResponse(ctx context.Context, params *ListProfilesParams, reqEditors ...RequestEditorFn) (*ListProfilesResp, error) {
	rsp, err := c.ListProfiles(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListProfilesResp(rsp)
}

// UpdateProfilesWithBodyWithResponse request with arbitrary body returning *UpdateProfilesResp
func (c *ClientWithResponses) UpdateProfilesWithBodyWithResponse(ctx context.Context, params *UpdateProfilesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateProfilesResp, error) {
	rsp, err := c.UpdateProfilesWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateProfilesResp(rsp)
}

func (c *ClientWithResponses) UpdateProfilesWithResponse(ctx context.Context, params *UpdateProfilesParams, body UpdateProfilesJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateProfilesResp, error) {
	rsp, err := c.UpdateProfiles(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateProfilesResp(rsp)
}

// GetProfileByIdWithResponse request returning *GetProfileByIdResp
func (c *ClientWithResponses) GetProfileByIdWithResponse(ctx context.Context, profileId int64, params *GetProfileByIdParams, reqEditors ...RequestEditorFn) (*GetProfileByIdResp, error) {
	rsp, err := c.GetProfileById(ctx, profileId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProfileByIdResp(rsp)
}

// ParseListProfilesResp parses an HTTP response from a ListProfilesWithResponse call
func ParseListProfilesResp(rsp *http.Response) (*ListProfilesResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListProfilesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Profile
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateProfilesResp parses an HTTP response from a UpdateProfilesWithResponse call
func ParseUpdateProfilesResp(rsp *http.Response) (*UpdateProfilesResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateProfilesResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ProfileResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetProfileByIdResp parses an HTTP response from a GetProfileByIdWithResponse call
func ParseGetProfileByIdResp(rsp *http.Response) (*GetProfileByIdResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetProfileByIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Profile
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}