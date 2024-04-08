package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemInstancesRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarViewItemInstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarViewItemInstancesRequestBuilderGetQueryParameters get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
type ItemCalendarViewItemInstancesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarViewItemInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarViewItemInstancesRequestBuilderGetQueryParameters
}
// NewItemCalendarViewItemInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewItemCalendarViewItemInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemInstancesRequestBuilder) {
    m := &ItemCalendarViewItemInstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/instances{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarViewItemInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewItemCalendarViewItemInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemCalendarViewItemInstancesRequestBuilder) Count()(*ItemCalendarViewItemInstancesCountRequestBuilder) {
    return NewItemCalendarViewItemInstancesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-list-instances?view=graph-rest-1.0
func (m *ItemCalendarViewItemInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarViewItemInstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable), nil
}
// MicrosoftGraphDelta provides operations to call the delta method.
func (m *ItemCalendarViewItemInstancesRequestBuilder) MicrosoftGraphDelta()(*ItemCalendarViewItemInstancesMicrosoftGraphDeltaRequestBuilder) {
    return NewItemCalendarViewItemInstancesMicrosoftGraphDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
func (m *ItemCalendarViewItemInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarViewItemInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
