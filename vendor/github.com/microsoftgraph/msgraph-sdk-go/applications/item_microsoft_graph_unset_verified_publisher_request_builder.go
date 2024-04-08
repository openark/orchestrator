package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder provides operations to call the unsetVerifiedPublisher method.
type ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderInternal instantiates a new MicrosoftGraphUnsetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder) {
    m := &ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/microsoft.graph.unsetVerifiedPublisher";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder instantiates a new MicrosoftGraphUnsetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unset the verifiedPublisher previously set on an application, removing all verified publisher properties. For more information, see Publisher verification.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/application-unsetverifiedpublisher?view=graph-rest-1.0
func (m *ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unset the verifiedPublisher previously set on an application, removing all verified publisher properties. For more information, see Publisher verification.
func (m *ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
