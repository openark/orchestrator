package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyUsersMicrosoftGraphDismissRequestBuilder provides operations to call the dismiss method.
type RiskyUsersMicrosoftGraphDismissRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyUsersMicrosoftGraphDismissRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyUsersMicrosoftGraphDismissRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyUsersMicrosoftGraphDismissRequestBuilderInternal instantiates a new MicrosoftGraphDismissRequestBuilder and sets the default values.
func NewRiskyUsersMicrosoftGraphDismissRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersMicrosoftGraphDismissRequestBuilder) {
    m := &RiskyUsersMicrosoftGraphDismissRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyUsers/microsoft.graph.dismiss";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRiskyUsersMicrosoftGraphDismissRequestBuilder instantiates a new MicrosoftGraphDismissRequestBuilder and sets the default values.
func NewRiskyUsersMicrosoftGraphDismissRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersMicrosoftGraphDismissRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUsersMicrosoftGraphDismissRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss the risk of one or more riskyUser objects. This action sets the targeted user's risk level to none.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/riskyuser-dismiss?view=graph-rest-1.0
func (m *RiskyUsersMicrosoftGraphDismissRequestBuilder) Post(ctx context.Context, body RiskyUsersMicrosoftGraphDismissDismissPostRequestBodyable, requestConfiguration *RiskyUsersMicrosoftGraphDismissRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation dismiss the risk of one or more riskyUser objects. This action sets the targeted user's risk level to none.
func (m *RiskyUsersMicrosoftGraphDismissRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyUsersMicrosoftGraphDismissDismissPostRequestBodyable, requestConfiguration *RiskyUsersMicrosoftGraphDismissRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
