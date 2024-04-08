package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsCallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
type CallsCallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsCallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallsCallItemRequestBuilderGetQueryParameters get calls from communications
type CallsCallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallsCallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallsCallItemRequestBuilderGetQueryParameters
}
// CallsCallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AudioRoutingGroups provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) AudioRoutingGroups()(*CallsItemAudioRoutingGroupsRequestBuilder) {
    return NewCallsItemAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AudioRoutingGroupsById provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) AudioRoutingGroupsById(id string)(*CallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return NewCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewCallsCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    m := &CallsCallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ContentSharingSessions()(*CallsItemContentSharingSessionsRequestBuilder) {
    return NewCallsItemContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentSharingSessionsById provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ContentSharingSessionsById(id string)(*CallsItemContentSharingSessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return NewCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property calls for communications
func (m *CallsCallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get calls from communications
func (m *CallsCallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable), nil
}
// MicrosoftGraphAddLargeGalleryView provides operations to call the addLargeGalleryView method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphAddLargeGalleryView()(*CallsItemMicrosoftGraphAddLargeGalleryViewRequestBuilder) {
    return NewCallsItemMicrosoftGraphAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAnswer provides operations to call the answer method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphAnswer()(*CallsItemMicrosoftGraphAnswerRequestBuilder) {
    return NewCallsItemMicrosoftGraphAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancelMediaProcessing provides operations to call the cancelMediaProcessing method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphCancelMediaProcessing()(*CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder) {
    return NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChangeScreenSharingRole provides operations to call the changeScreenSharingRole method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphChangeScreenSharingRole()(*CallsItemMicrosoftGraphChangeScreenSharingRoleRequestBuilder) {
    return NewCallsItemMicrosoftGraphChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphKeepAlive provides operations to call the keepAlive method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphKeepAlive()(*CallsItemMicrosoftGraphKeepAliveRequestBuilder) {
    return NewCallsItemMicrosoftGraphKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMute provides operations to call the mute method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphMute()(*CallsItemMicrosoftGraphMuteRequestBuilder) {
    return NewCallsItemMicrosoftGraphMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPlayPrompt provides operations to call the playPrompt method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphPlayPrompt()(*CallsItemMicrosoftGraphPlayPromptRequestBuilder) {
    return NewCallsItemMicrosoftGraphPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecordResponse provides operations to call the recordResponse method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphRecordResponse()(*CallsItemMicrosoftGraphRecordResponseRequestBuilder) {
    return NewCallsItemMicrosoftGraphRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRedirect provides operations to call the redirect method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphRedirect()(*CallsItemMicrosoftGraphRedirectRequestBuilder) {
    return NewCallsItemMicrosoftGraphRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReject provides operations to call the reject method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphReject()(*CallsItemMicrosoftGraphRejectRequestBuilder) {
    return NewCallsItemMicrosoftGraphRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSubscribeToTone provides operations to call the subscribeToTone method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphSubscribeToTone()(*CallsItemMicrosoftGraphSubscribeToToneRequestBuilder) {
    return NewCallsItemMicrosoftGraphSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTransfer provides operations to call the transfer method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphTransfer()(*CallsItemMicrosoftGraphTransferRequestBuilder) {
    return NewCallsItemMicrosoftGraphTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnmute provides operations to call the unmute method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphUnmute()(*CallsItemMicrosoftGraphUnmuteRequestBuilder) {
    return NewCallsItemMicrosoftGraphUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateRecordingStatus provides operations to call the updateRecordingStatus method.
func (m *CallsCallItemRequestBuilder) MicrosoftGraphUpdateRecordingStatus()(*CallsItemMicrosoftGraphUpdateRecordingStatusRequestBuilder) {
    return NewCallsItemMicrosoftGraphUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) Operations()(*CallsItemOperationsRequestBuilder) {
    return NewCallsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) OperationsById(id string)(*CallsItemOperationsCommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return NewCallsItemOperationsCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Participants provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) Participants()(*CallsItemParticipantsRequestBuilder) {
    return NewCallsItemParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParticipantsById provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ParticipantsById(id string)(*CallsItemParticipantsParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return NewCallsItemParticipantsParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property calls in communications
func (m *CallsCallItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable), nil
}
// ToDeleteRequestInformation delete navigation property calls for communications
func (m *CallsCallItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get calls from communications
func (m *CallsCallItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property calls in communications
func (m *CallsCallItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
