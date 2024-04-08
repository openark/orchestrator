package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters retrieve the filter applied to the column. Read-only.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder instantiates a new FilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filter for drives
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApply()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyBottomItemsFilter provides operations to call the applyBottomItemsFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyBottomItemsFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyBottomPercentFilter provides operations to call the applyBottomPercentFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyBottomPercentFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyCellColorFilter provides operations to call the applyCellColorFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyCellColorFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCellColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyCustomFilter provides operations to call the applyCustomFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyCustomFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyDynamicFilter provides operations to call the applyDynamicFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyDynamicFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyDynamicFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyFontColorFilter provides operations to call the applyFontColorFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyFontColorFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyFontColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyIconFilter provides operations to call the applyIconFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyIconFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyIconFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyTopItemsFilter provides operations to call the applyTopItemsFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyTopItemsFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyTopPercentFilter provides operations to call the applyTopPercentFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyTopPercentFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyTopPercentFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyValuesFilter provides operations to call the applyValuesFilter method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphApplyValuesFilter()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyValuesFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClear provides operations to call the clear method.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) MicrosoftGraphClear()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphClearRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphClearRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property filter in drives
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
