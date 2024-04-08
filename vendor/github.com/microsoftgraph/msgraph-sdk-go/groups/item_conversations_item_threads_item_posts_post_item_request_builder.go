package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsPostItemRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type ItemConversationsItemThreadsItemPostsPostItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters get posts from groups
type ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) Attachments()(*ItemConversationsItemThreadsItemPostsItemAttachmentsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) AttachmentsById(id string)(*ItemConversationsItemThreadsItemPostsItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemConversationsItemThreadsItemPostsPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsPostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemConversationsItemThreadsItemPostsPostItemRequestBuilder instantiates a new PostItemRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsPostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsPostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) Extensions()(*ItemConversationsItemThreadsItemPostsItemExtensionsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) ExtensionsById(id string)(*ItemConversationsItemThreadsItemPostsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get posts from groups
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable), nil
}
// InReplyTo provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) InReplyTo()(*ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) MicrosoftGraphForward()(*ItemConversationsItemThreadsItemPostsItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReply provides operations to call the reply method.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) MicrosoftGraphReply()(*ItemConversationsItemThreadsItemPostsItemMicrosoftGraphReplyRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemMicrosoftGraphReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) MultiValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) SingleValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get posts from groups
func (m *ItemConversationsItemThreadsItemPostsPostItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
