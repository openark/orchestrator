package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder provides operations to call the clear method.
type ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderInternal instantiates a new MicrosoftGraphClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/sort/microsoft.graph.clear";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder instantiates a new MicrosoftGraphClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clears the sorting that is currently on the table. While this doesn't modify the table's ordering, it clears the state of the header buttons.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/tablesort-clear?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clears the sorting that is currently on the table. While this doesn't modify the table's ordering, it clears the state of the header buttons.
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphClearRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
