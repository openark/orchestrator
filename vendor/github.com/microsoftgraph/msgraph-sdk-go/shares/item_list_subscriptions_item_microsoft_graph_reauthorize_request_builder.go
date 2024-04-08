package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder provides operations to call the reauthorize method.
type ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderInternal instantiates a new MicrosoftGraphReauthorizeRequestBuilder and sets the default values.
func NewItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder) {
    m := &ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list/subscriptions/{subscription%2Did}/microsoft.graph.reauthorize";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder instantiates a new MicrosoftGraphReauthorizeRequestBuilder and sets the default values.
func NewItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action reauthorize
func (m *ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action reauthorize
func (m *ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListSubscriptionsItemMicrosoftGraphReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
