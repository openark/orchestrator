package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) AttachmentsById(id string)(*ItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Content()(*ItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property messages for users
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ExtensionsById(id string)(*ItemMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the collection of messages in the mailFolder.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// MicrosoftGraphCopy provides operations to call the copy method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCopy()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCopyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateForward provides operations to call the createForward method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateForward()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateForwardRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReply provides operations to call the createReply method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReply()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateReplyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReplyAll provides operations to call the createReplyAll method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReplyAll()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphForward()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMove provides operations to call the move method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphMove()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphMoveRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReply provides operations to call the reply method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphReply()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphReplyRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReplyAll provides operations to call the replyAll method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphReplyAll()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphReplyAllRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSend provides operations to call the send method.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphSend()(*ItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphSendRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMicrosoftGraphSendRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedProperties()(*ItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property messages in users
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedProperties()(*ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for users
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of messages in the mailFolder.
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in users
func (m *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
