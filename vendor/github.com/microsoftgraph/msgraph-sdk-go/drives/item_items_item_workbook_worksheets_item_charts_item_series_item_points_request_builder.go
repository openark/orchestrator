package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder provides operations to manage the points property of the microsoft.graph.workbookChartSeries entity.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetQueryParameters retrieve a list of chartpoints objects.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetQueryParameters struct {
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
// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderInternal instantiates a new PointsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/series/{workbookChartSeries%2Did}/points{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder instantiates a new PointsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) Count()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve a list of chartpoints objects.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chartseries-list-points?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartPointCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointCollectionResponseable), nil
}
// MicrosoftGraphCount provides operations to call the count method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) MicrosoftGraphCount()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsMicrosoftGraphCountRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsMicrosoftGraphCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphItemAtWithIndex provides operations to call the itemAt method.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) MicrosoftGraphItemAtWithIndex(index *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsMicrosoftGraphItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsMicrosoftGraphItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index)
}
// Post use this API to create a new ChartPoints.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chartseries-post-points?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartPointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable), nil
}
// ToGetRequestInformation retrieve a list of chartpoints objects.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation use this API to create a new ChartPoints.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
