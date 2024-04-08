package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationPhoneMethodsRequestBuilder provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationPhoneMethodsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationPhoneMethodsRequestBuilderGetQueryParameters retrieve a list of phone authentication method objects for a user. This will return up to three objects, as a user can have up to three phones usable for authentication. This method is available only for standard Azure AD and B2B users, but not B2C users.
type ItemAuthenticationPhoneMethodsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemAuthenticationPhoneMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPhoneMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationPhoneMethodsRequestBuilderGetQueryParameters
}
// ItemAuthenticationPhoneMethodsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPhoneMethodsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationPhoneMethodsRequestBuilderInternal instantiates a new PhoneMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationPhoneMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhoneMethodsRequestBuilder) {
    m := &ItemAuthenticationPhoneMethodsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/phoneMethods{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAuthenticationPhoneMethodsRequestBuilder instantiates a new PhoneMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationPhoneMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhoneMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPhoneMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemAuthenticationPhoneMethodsRequestBuilder) Count()(*ItemAuthenticationPhoneMethodsCountRequestBuilder) {
    return NewItemAuthenticationPhoneMethodsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve a list of phone authentication method objects for a user. This will return up to three objects, as a user can have up to three phones usable for authentication. This method is available only for standard Azure AD and B2B users, but not B2C users.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/authentication-list-phonemethods?view=graph-rest-1.0
func (m *ItemAuthenticationPhoneMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationPhoneMethodsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePhoneAuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodCollectionResponseable), nil
}
// Post add a new phone authentication method for a user. A user may only have one phone of each type, captured in the **phoneType** property. This means, for example, adding a `mobile` phone to a user with a preexisting `mobile` phone will fail. Additionally, a user must always have a `mobile` phone before adding an `alternateMobile` phone. Adding a phone number makes it available for use in both Azure multi-factor authentication (MFA) and self-service password reset (SSPR), if enabled. Additionally, if a user is enabled by policy to use SMS sign-in and a `mobile` number is added, the system will attempt to register the number for use in that system.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/authentication-post-phonemethods?view=graph-rest-1.0
func (m *ItemAuthenticationPhoneMethodsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodable, requestConfiguration *ItemAuthenticationPhoneMethodsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePhoneAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodable), nil
}
// ToGetRequestInformation retrieve a list of phone authentication method objects for a user. This will return up to three objects, as a user can have up to three phones usable for authentication. This method is available only for standard Azure AD and B2B users, but not B2C users.
func (m *ItemAuthenticationPhoneMethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPhoneMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation add a new phone authentication method for a user. A user may only have one phone of each type, captured in the **phoneType** property. This means, for example, adding a `mobile` phone to a user with a preexisting `mobile` phone will fail. Additionally, a user must always have a `mobile` phone before adding an `alternateMobile` phone. Adding a phone number makes it available for use in both Azure multi-factor authentication (MFA) and self-service password reset (SSPR), if enabled. Additionally, if a user is enabled by policy to use SMS sign-in and a `mobile` number is added, the system will attempt to register the number for use in that system.
func (m *ItemAuthenticationPhoneMethodsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PhoneAuthenticationMethodable, requestConfiguration *ItemAuthenticationPhoneMethodsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
