package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal instantiates a new MicrosoftGraphCheckMemberObjectsRequestBuilder and sets the default values.
func NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    m := &DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/deletedItems/{directoryObject%2Did}/microsoft.graph.checkMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder instantiates a new MicrosoftGraphCheckMemberObjectsRequestBuilder and sets the default values.
func NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration)(DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable, requestConfiguration *DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
