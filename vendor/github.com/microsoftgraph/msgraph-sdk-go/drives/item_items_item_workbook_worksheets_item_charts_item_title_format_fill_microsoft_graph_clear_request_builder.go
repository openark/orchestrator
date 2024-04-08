package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder provides operations to call the clear method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderInternal instantiates a new MicrosoftGraphClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/title/format/fill/microsoft.graph.clear";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder instantiates a new MicrosoftGraphClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the fill color of a chart element.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chartfill-clear?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear the fill color of a chart element.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillMicrosoftGraphClearRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
