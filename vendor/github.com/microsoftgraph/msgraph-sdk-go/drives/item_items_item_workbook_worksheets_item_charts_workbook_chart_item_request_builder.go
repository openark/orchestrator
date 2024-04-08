package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters returns collection of charts that are part of the worksheet. Read-only.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Axes()(*ItemItemsItemWorkbookWorksheetsItemChartsItemAxesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) DataLabels()(*ItemItemsItemWorkbookWorksheetsItemChartsItemDataLabelsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property charts for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Format()(*ItemItemsItemWorkbookWorksheetsItemChartsItemFormatRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get returns collection of charts that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable), nil
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Legend()(*ItemItemsItemWorkbookWorksheetsItemChartsItemLegendRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImage provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphImage()(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImageWithWidth provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphImageWithWidth(width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width)
}
// MicrosoftGraphImageWithWidthWithHeight provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphImageWithWidthWithHeight(height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, height, width)
}
// MicrosoftGraphImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, fittingMode, height, width)
}
// MicrosoftGraphSetData provides operations to call the setData method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphSetData()(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetPosition provides operations to call the setPosition method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) MicrosoftGraphSetPosition()(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property charts in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable), nil
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Series()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SeriesById provides operations to manage the series property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) SeriesById(id string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesWorkbookChartSeriesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesWorkbookChartSeriesItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Title()(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property charts for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property charts in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
