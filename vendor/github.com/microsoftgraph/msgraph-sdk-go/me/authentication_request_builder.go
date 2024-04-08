package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationRequestBuilderGetQueryParameters
}
// AuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AuthenticationRequestBuilder) EmailMethods()(*AuthenticationEmailMethodsRequestBuilder) {
    return NewAuthenticationEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmailMethodsById provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*AuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Fido2Methods provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Fido2Methods()(*AuthenticationFido2MethodsRequestBuilder) {
    return NewAuthenticationFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Fido2MethodsById provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*AuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
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
func (m *AuthenticationRequestBuilder) Methods()(*AuthenticationMethodsRequestBuilder) {
    return NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MethodsById provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*AuthenticationMethodsAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return NewAuthenticationMethodsAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftAuthenticatorMethods provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*AuthenticationMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftAuthenticatorMethodsById provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*AuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Operations()(*AuthenticationOperationsRequestBuilder) {
    return NewAuthenticationOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*AuthenticationOperationsLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return NewAuthenticationOperationsLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// PasswordMethods provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordMethods()(*AuthenticationPasswordMethodsRequestBuilder) {
    return NewAuthenticationPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PasswordMethodsById provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*AuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
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
func (m *AuthenticationRequestBuilder) PhoneMethods()(*AuthenticationPhoneMethodsRequestBuilder) {
    return NewAuthenticationPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PhoneMethodsById provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*AuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SoftwareOathMethods provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*AuthenticationSoftwareOathMethodsRequestBuilder) {
    return NewAuthenticationSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SoftwareOathMethodsById provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*AuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TemporaryAccessPassMethods provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*AuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    return NewAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemporaryAccessPassMethodsById provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*AuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AuthenticationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*AuthenticationWindowsHelloForBusinessMethodsRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsHelloForBusinessMethodsById provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*AuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return NewAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
