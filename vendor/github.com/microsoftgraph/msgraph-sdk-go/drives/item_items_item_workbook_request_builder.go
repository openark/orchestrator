package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookRequestBuilder provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
type ItemItemsItemWorkbookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookRequestBuilderGetQueryParameters for files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
type ItemItemsItemWorkbookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application provides operations to manage the application property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Application()(*ItemItemsItemWorkbookApplicationRequestBuilder) {
    return NewItemItemsItemWorkbookApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Comments provides operations to manage the comments property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Comments()(*ItemItemsItemWorkbookCommentsRequestBuilder) {
    return NewItemItemsItemWorkbookCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CommentsById provides operations to manage the comments property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) CommentsById(id string)(*ItemItemsItemWorkbookCommentsWorkbookCommentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookComment%2Did"] = id
    }
    return NewItemItemsItemWorkbookCommentsWorkbookCommentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemsItemWorkbookRequestBuilderInternal instantiates a new WorkbookRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookRequestBuilder) {
    m := &ItemItemsItemWorkbookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookRequestBuilder instantiates a new WorkbookRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property workbook for drives
func (m *ItemItemsItemWorkbookRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(error) {
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
// Functions provides operations to manage the functions property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Functions()(*ItemItemsItemWorkbookFunctionsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get for files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
func (m *ItemItemsItemWorkbookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable), nil
}
// MicrosoftGraphCloseSession provides operations to call the closeSession method.
func (m *ItemItemsItemWorkbookRequestBuilder) MicrosoftGraphCloseSession()(*ItemItemsItemWorkbookMicrosoftGraphCloseSessionRequestBuilder) {
    return NewItemItemsItemWorkbookMicrosoftGraphCloseSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateSession provides operations to call the createSession method.
func (m *ItemItemsItemWorkbookRequestBuilder) MicrosoftGraphCreateSession()(*ItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilder) {
    return NewItemItemsItemWorkbookMicrosoftGraphCreateSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRefreshSession provides operations to call the refreshSession method.
func (m *ItemItemsItemWorkbookRequestBuilder) MicrosoftGraphRefreshSession()(*ItemItemsItemWorkbookMicrosoftGraphRefreshSessionRequestBuilder) {
    return NewItemItemsItemWorkbookMicrosoftGraphRefreshSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSessionInfoResourceWithKey provides operations to call the sessionInfoResource method.
func (m *ItemItemsItemWorkbookRequestBuilder) MicrosoftGraphSessionInfoResourceWithKey(key *string)(*ItemItemsItemWorkbookMicrosoftGraphSessionInfoResourceWithKeyRequestBuilder) {
    return NewItemItemsItemWorkbookMicrosoftGraphSessionInfoResourceWithKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, key)
}
// MicrosoftGraphTableRowOperationResultWithKey provides operations to call the tableRowOperationResult method.
func (m *ItemItemsItemWorkbookRequestBuilder) MicrosoftGraphTableRowOperationResultWithKey(key *string)(*ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyRequestBuilder) {
    return NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, key)
}
// Names provides operations to manage the names property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Names()(*ItemItemsItemWorkbookNamesRequestBuilder) {
    return NewItemItemsItemWorkbookNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NamesById provides operations to manage the names property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) NamesById(id string)(*ItemItemsItemWorkbookNamesWorkbookNamedItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem%2Did"] = id
    }
    return NewItemItemsItemWorkbookNamesWorkbookNamedItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Operations()(*ItemItemsItemWorkbookOperationsRequestBuilder) {
    return NewItemItemsItemWorkbookOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) OperationsById(id string)(*ItemItemsItemWorkbookOperationsWorkbookOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookOperation%2Did"] = id
    }
    return NewItemItemsItemWorkbookOperationsWorkbookOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property workbook in drives
func (m *ItemItemsItemWorkbookRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *ItemItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable), nil
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Tables()(*ItemItemsItemWorkbookTablesRequestBuilder) {
    return NewItemItemsItemWorkbookTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TablesById provides operations to manage the tables property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) TablesById(id string)(*ItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable%2Did"] = id
    }
    return NewItemItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property workbook for drives
func (m *ItemItemsItemWorkbookRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation for files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
func (m *ItemItemsItemWorkbookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property workbook in drives
func (m *ItemItemsItemWorkbookRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Workbookable, requestConfiguration *ItemItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Worksheets provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) Worksheets()(*ItemItemsItemWorkbookWorksheetsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WorksheetsById provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
func (m *ItemItemsItemWorkbookRequestBuilder) WorksheetsById(id string)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookWorksheet%2Did"] = id
    }
    return NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
