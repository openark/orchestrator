package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemEventsMicrosoftGraphDeltaRequestBuilder provides operations to call the delta method.
type ItemEventsMicrosoftGraphDeltaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemEventsMicrosoftGraphDeltaRequestBuilderGetQueryParameters get a set of event resources that have been added, deleted, or updated in a **calendarView** (a range of events defined by start and end dates) of the user's primary calendar. Typically, synchronizing events in a **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified **calendarView**, without having to fetch all the events of that calendar from the server every time.
type ItemEventsMicrosoftGraphDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemEventsMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsMicrosoftGraphDeltaRequestBuilderGetQueryParameters
}
// NewItemEventsMicrosoftGraphDeltaRequestBuilderInternal instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewItemEventsMicrosoftGraphDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsMicrosoftGraphDeltaRequestBuilder) {
    m := &ItemEventsMicrosoftGraphDeltaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEventsMicrosoftGraphDeltaRequestBuilder instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewItemEventsMicrosoftGraphDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsMicrosoftGraphDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsMicrosoftGraphDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of event resources that have been added, deleted, or updated in a **calendarView** (a range of events defined by start and end dates) of the user's primary calendar. Typically, synchronizing events in a **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified **calendarView**, without having to fetch all the events of that calendar from the server every time.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-delta?view=graph-rest-1.0
func (m *ItemEventsMicrosoftGraphDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(ItemEventsMicrosoftGraphDeltaDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemEventsMicrosoftGraphDeltaDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEventsMicrosoftGraphDeltaDeltaResponseable), nil
}
// ToGetRequestInformation get a set of event resources that have been added, deleted, or updated in a **calendarView** (a range of events defined by start and end dates) of the user's primary calendar. Typically, synchronizing events in a **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified **calendarView**, without having to fetch all the events of that calendar from the server every time.
func (m *ItemEventsMicrosoftGraphDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
