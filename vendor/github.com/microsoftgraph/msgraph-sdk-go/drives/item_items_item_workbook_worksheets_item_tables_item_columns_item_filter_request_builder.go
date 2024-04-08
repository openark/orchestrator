package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters retrieve the filter applied to the column. Read-only.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filter for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the filter applied to the column. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable), nil
}
// MicrosoftGraphApply provides operations to call the apply method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApply()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyBottomItemsFilter provides operations to call the applyBottomItemsFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyBottomItemsFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyBottomItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyBottomPercentFilter provides operations to call the applyBottomPercentFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyBottomPercentFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyCellColorFilter provides operations to call the applyCellColorFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyCellColorFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyCellColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyCustomFilter provides operations to call the applyCustomFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyCustomFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyDynamicFilter provides operations to call the applyDynamicFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyDynamicFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyDynamicFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyFontColorFilter provides operations to call the applyFontColorFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyFontColorFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyFontColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyIconFilter provides operations to call the applyIconFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyIconFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyIconFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyTopItemsFilter provides operations to call the applyTopItemsFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyTopItemsFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyTopPercentFilter provides operations to call the applyTopPercentFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyTopPercentFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyValuesFilter provides operations to call the applyValuesFilter method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyValuesFilter()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyValuesFilterRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClear provides operations to call the clear method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphClear()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphClearRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphClearRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property filter in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable), nil
}
// ToDeleteRequestInformation delete navigation property filter for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the filter applied to the column. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property filter in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
