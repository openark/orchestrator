package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder provides operations to call the image method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderInternal instantiates a new MicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fittingMode *string, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/microsoft.graph.image(width={width},height={height},fittingMode='{fittingMode}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if fittingMode != nil {
        urlTplParams["fittingMode"] = *fittingMode
    }
    if height != nil {
        urlTplParams["height"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*height), 10)
    }
    if width != nil {
        urlTplParams["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder instantiates a new MicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function image
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeImageWithWidthWithHeightWithFittingModeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeImageWithWidthWithHeightWithFittingModeResponseable), nil
}
// ToGetRequestInformation invoke function image
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
