package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Charts provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Charts()(*ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChartsById provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ChartsById(id string)(*ItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemChartsWorkbookChartItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property worksheets for drives
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents a collection of worksheets associated with the workbook. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable), nil
}
// MicrosoftGraphCellWithRowWithColumn provides operations to call the cell method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) MicrosoftGraphCellWithRowWithColumn(column *int32, row *int32)(*ItemItemsItemWorkbookWorksheetsItemMicrosoftGraphCellWithRowWithColumnRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemMicrosoftGraphCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column, row)
}
// MicrosoftGraphRange provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) MicrosoftGraphRange()(*ItemItemsItemWorkbookWorksheetsItemMicrosoftGraphRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemMicrosoftGraphRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRangeWithAddress provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) MicrosoftGraphRangeWithAddress(address *string)(*ItemItemsItemWorkbookWorksheetsItemMicrosoftGraphRangeWithAddressRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemMicrosoftGraphRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address)
}
// MicrosoftGraphUsedRange provides operations to call the usedRange method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) MicrosoftGraphUsedRange()(*ItemItemsItemWorkbookWorksheetsItemMicrosoftGraphUsedRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemMicrosoftGraphUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUsedRangeWithValuesOnly provides operations to call the usedRange method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) MicrosoftGraphUsedRangeWithValuesOnly(valuesOnly *bool)(*ItemItemsItemWorkbookWorksheetsItemMicrosoftGraphUsedRangeWithValuesOnlyRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemMicrosoftGraphUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly)
}
// Names provides operations to manage the names property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Names()(*ItemItemsItemWorkbookWorksheetsItemNamesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NamesById provides operations to manage the names property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) NamesById(id string)(*ItemItemsItemWorkbookWorksheetsItemNamesWorkbookNamedItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemNamesWorkbookNamedItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property worksheets in drives
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable), nil
}
// PivotTables provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) PivotTables()(*ItemItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PivotTablesById provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) PivotTablesById(id string)(*ItemItemsItemWorkbookWorksheetsItemPivotTablesWorkbookPivotTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemPivotTablesWorkbookPivotTableItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Protection provides operations to manage the protection property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Protection()(*ItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Tables()(*ItemItemsItemWorkbookWorksheetsItemTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TablesById provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) TablesById(id string)(*ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property worksheets for drives
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property worksheets in drives
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
