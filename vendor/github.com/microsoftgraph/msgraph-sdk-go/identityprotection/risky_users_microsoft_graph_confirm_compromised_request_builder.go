package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder provides operations to call the confirmCompromised method.
type RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderInternal instantiates a new MicrosoftGraphConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder) {
    m := &RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyUsers/microsoft.graph.confirmCompromised";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder instantiates a new MicrosoftGraphConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderInternal(urlParams, requestAdapter)
}
// Post confirm one or more riskyUser objects as compromised. This action sets the targeted user's risk level to high.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/riskyuser-confirmcompromised?view=graph-rest-1.0
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder) Post(ctx context.Context, body RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation confirm one or more riskyUser objects as compromised. This action sets the targeted user's risk level to high.
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyUsersMicrosoftGraphConfirmCompromisedRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
