package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder provides operations to call the stopHoldMusic method.
type CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderInternal instantiates a new MicrosoftGraphStopHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder) {
    m := &CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}/participants/{participant%2Did}/microsoft.graph.stopHoldMusic";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder instantiates a new MicrosoftGraphStopHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reincorporate a participant previously put on hold to the call.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/participant-stopholdmusic?view=graph-rest-1.0
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder) Post(ctx context.Context, body CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StopHoldMusicOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStopHoldMusicOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StopHoldMusicOperationable), nil
}
// ToPostRequestInformation reincorporate a participant previously put on hold to the call.
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
