package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphForceDeleteRequestBuilder provides operations to call the forceDelete method.
type ItemMicrosoftGraphForceDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphForceDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphForceDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphForceDeleteRequestBuilderInternal instantiates a new MicrosoftGraphForceDeleteRequestBuilder and sets the default values.
func NewItemMicrosoftGraphForceDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphForceDeleteRequestBuilder) {
    m := &ItemMicrosoftGraphForceDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain%2Did}/microsoft.graph.forceDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphForceDeleteRequestBuilder instantiates a new MicrosoftGraphForceDeleteRequestBuilder and sets the default values.
func NewItemMicrosoftGraphForceDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphForceDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphForceDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post deletes a domain using an asynchronous long-running operation. Prior to calling forceDelete, you must update or remove any references to **Exchange** as the provisioning service. The following actions are performed as part of this operation: After the domain deletion completes, API operations for the deleted domain will return a HTTP 404 status code. To verify deletion of a domain, you can perform a get domain operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/domain-forcedelete?view=graph-rest-1.0
func (m *ItemMicrosoftGraphForceDeleteRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphForceDeleteForceDeletePostRequestBodyable, requestConfiguration *ItemMicrosoftGraphForceDeleteRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation deletes a domain using an asynchronous long-running operation. Prior to calling forceDelete, you must update or remove any references to **Exchange** as the provisioning service. The following actions are performed as part of this operation: After the domain deletion completes, API operations for the deleted domain will return a HTTP 404 status code. To verify deletion of a domain, you can perform a get domain operation.
func (m *ItemMicrosoftGraphForceDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphForceDeleteForceDeletePostRequestBodyable, requestConfiguration *ItemMicrosoftGraphForceDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
