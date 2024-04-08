package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Attachments()(*ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) AttachmentsById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Extensions()(*ItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) ExtensionsById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get read-only. Supports $expand.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
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
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MicrosoftGraphForward()(*ItemConversationsItemThreadsItemPostsItemInReplyToMicrosoftGraphForwardRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReply provides operations to call the reply method.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MicrosoftGraphReply()(*ItemConversationsItemThreadsItemPostsItemInReplyToMicrosoftGraphReplyRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMicrosoftGraphReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation read-only. Supports $expand.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
