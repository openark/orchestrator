package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters represents a collection of all the columns in the table. Read-only.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal instantiates a new WorkbookTableColumnItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder instantiates a new WorkbookTableColumnItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property columns for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Filter provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Filter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get represents a collection of all the columns in the table. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable), nil
}
// MicrosoftGraphDataBodyRange provides operations to call the dataBodyRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) MicrosoftGraphDataBodyRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHeaderRowRange provides operations to call the headerRowRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) MicrosoftGraphHeaderRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRange provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) MicrosoftGraphRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTotalRowRange provides operations to call the totalRowRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) MicrosoftGraphTotalRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemMicrosoftGraphTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property columns in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable), nil
}
// ToDeleteRequestInformation delete navigation property columns for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property columns in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
