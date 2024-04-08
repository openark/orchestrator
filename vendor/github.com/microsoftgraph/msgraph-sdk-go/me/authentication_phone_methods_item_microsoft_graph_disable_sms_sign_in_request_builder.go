package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder provides operations to call the disableSmsSignIn method.
type AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal instantiates a new MicrosoftGraphDisableSmsSignInRequestBuilder and sets the default values.
func NewAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) {
    m := &AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/phoneMethods/{phoneAuthenticationMethod%2Did}/microsoft.graph.disableSmsSignIn";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder instantiates a new MicrosoftGraphDisableSmsSignInRequestBuilder and sets the default values.
func NewAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post disable SMS sign-in for an existing `mobile` phone number registered to a user. The number will no longer be available for SMS sign-in, which can prevent your user from signing in.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/phoneauthenticationmethod-disablesmssignin?view=graph-rest-1.0
func (m *AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation disable SMS sign-in for an existing `mobile` phone number registered to a user. The number will no longer be available for SMS sign-in, which can prevent your user from signing in.
func (m *AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
