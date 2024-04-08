package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getDirectRoutingCalls method.
type CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters invoke function getDirectRoutingCalls
type CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new MicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getDirectRoutingCalls(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if fromDateTime != nil {
        urlTplParams["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        urlTplParams["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new MicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getDirectRoutingCalls
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable), nil
}
// ToGetRequestInformation invoke function getDirectRoutingCalls
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
