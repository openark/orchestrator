package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters content in a message hosted by Microsoft Teams - for example, images or code snippets.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetQueryParameters
}
// DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal instantiates a new ChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/hostedContents/{chatMessageHostedContent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder instantiates a new ChatMessageHostedContentItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property hostedContents for teamwork
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
// Patch update the navigation property hostedContents in teamwork
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatMessageHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable), nil
}
// ToDeleteRequestInformation delete navigation property hostedContents for teamwork
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hostedContents in teamwork
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatMessageHostedContentable, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
