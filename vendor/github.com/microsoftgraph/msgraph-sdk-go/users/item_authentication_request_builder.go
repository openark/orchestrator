package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type ItemAuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type ItemAuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationRequestBuilderGetQueryParameters
}
// ItemAuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewItemAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationRequestBuilder) {
    m := &ItemAuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewItemAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authentication for users
func (m *ItemAuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EmailMethods provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) EmailMethods()(*ItemAuthenticationEmailMethodsRequestBuilder) {
    return NewItemAuthenticationEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmailMethodsById provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) EmailMethodsById(id string)(*ItemAuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Fido2Methods provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) Fido2Methods()(*ItemAuthenticationFido2MethodsRequestBuilder) {
    return NewItemAuthenticationFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Fido2MethodsById provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) Fido2MethodsById(id string)(*ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the authentication methods that are supported for the user.
func (m *ItemAuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable), nil
}
// Methods provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) Methods()(*ItemAuthenticationMethodsRequestBuilder) {
    return NewItemAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MethodsById provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) MethodsById(id string)(*ItemAuthenticationMethodsAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationMethodsAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftAuthenticatorMethods provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*ItemAuthenticationMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftAuthenticatorMethodsById provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*ItemAuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) Operations()(*ItemAuthenticationOperationsRequestBuilder) {
    return NewItemAuthenticationOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) OperationsById(id string)(*ItemAuthenticationOperationsLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return NewItemAuthenticationOperationsLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// PasswordMethods provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) PasswordMethods()(*ItemAuthenticationPasswordMethodsRequestBuilder) {
    return NewItemAuthenticationPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PasswordMethodsById provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) PasswordMethodsById(id string)(*ItemAuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property authentication in users
func (m *ItemAuthenticationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *ItemAuthenticationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable), nil
}
// PhoneMethods provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) PhoneMethods()(*ItemAuthenticationPhoneMethodsRequestBuilder) {
    return NewItemAuthenticationPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PhoneMethodsById provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) PhoneMethodsById(id string)(*ItemAuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SoftwareOathMethods provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) SoftwareOathMethods()(*ItemAuthenticationSoftwareOathMethodsRequestBuilder) {
    return NewItemAuthenticationSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SoftwareOathMethodsById provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*ItemAuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TemporaryAccessPassMethods provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    return NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemporaryAccessPassMethodsById provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*ItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property authentication for users
func (m *ItemAuthenticationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the authentication methods that are supported for the user.
func (m *ItemAuthenticationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authentication in users
func (m *ItemAuthenticationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *ItemAuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// WindowsHelloForBusinessMethods provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*ItemAuthenticationWindowsHelloForBusinessMethodsRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsHelloForBusinessMethodsById provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *ItemAuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
