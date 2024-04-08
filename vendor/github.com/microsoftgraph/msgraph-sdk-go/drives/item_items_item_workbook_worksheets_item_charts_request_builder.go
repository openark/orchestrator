package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
type ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetQueryParameters retrieve a list of chart objects.
type ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal instantiates a new ChartsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder instantiates a new ChartsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) Count()(*ItemItemsItemWorkbookWorksheetsItemChartsCountRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve a list of chart objects.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/worksheet-list-charts?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartCollectionResponseable), nil
}
// MicrosoftGraphAdd provides operations to call the add method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) MicrosoftGraphAdd()(*ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCount provides operations to call the count method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) MicrosoftGraphCount()(*ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphCountRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphItemAtWithIndex provides operations to call the itemAt method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) MicrosoftGraphItemAtWithIndex(index *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index)
}
// MicrosoftGraphItemWithName provides operations to call the item method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) MicrosoftGraphItemWithName(name *string)(*ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphItemWithNameRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphItemWithNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, name)
}
// Post use this API to create a new Chart.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/worksheet-post-charts?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve a list of chart objects.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation use this API to create a new Chart.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
