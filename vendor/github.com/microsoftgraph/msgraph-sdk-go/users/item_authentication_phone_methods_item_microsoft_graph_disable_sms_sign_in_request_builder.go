package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder provides operations to call the disableSmsSignIn method.
type ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal instantiates a new MicrosoftGraphDisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) {
    m := &ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/phoneMethods/{phoneAuthenticationMethod%2Did}/microsoft.graph.disableSmsSignIn";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder instantiates a new MicrosoftGraphDisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post disable SMS sign-in for an existing `mobile` phone number registered to a user. The number will no longer be available for SMS sign-in, which can prevent your user from signing in.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/phoneauthenticationmethod-disablesmssignin?view=graph-rest-1.0
func (m *ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPhoneMethodsItemMicrosoftGraphDisableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
