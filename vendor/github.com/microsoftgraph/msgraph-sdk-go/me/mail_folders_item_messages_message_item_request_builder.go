package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MailFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MailFoldersItemMessagesMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MailFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// MailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*MailFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewMailFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) AttachmentsById(id string)(*MailFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMailFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMailFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMailFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersItemMessagesMessageItemRequestBuilder) {
    m := &MailFoldersItemMessagesMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMailFoldersItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMailFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Content()(*MailFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewMailFoldersItemMessagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property messages for me
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*MailFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewMailFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) ExtensionsById(id string)(*MailFoldersItemMessagesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMailFoldersItemMessagesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the collection of messages in the mailFolder.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
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
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCopy()(*MailFoldersItemMessagesItemMicrosoftGraphCopyRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateForward provides operations to call the createForward method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateForward()(*MailFoldersItemMessagesItemMicrosoftGraphCreateForwardRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReply provides operations to call the createReply method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReply()(*MailFoldersItemMessagesItemMicrosoftGraphCreateReplyRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReplyAll provides operations to call the createReplyAll method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReplyAll()(*MailFoldersItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphForward()(*MailFoldersItemMessagesItemMicrosoftGraphForwardRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMove provides operations to call the move method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphMove()(*MailFoldersItemMessagesItemMicrosoftGraphMoveRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReply provides operations to call the reply method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphReply()(*MailFoldersItemMessagesItemMicrosoftGraphReplyRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReplyAll provides operations to call the replyAll method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphReplyAll()(*MailFoldersItemMessagesItemMicrosoftGraphReplyAllRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSend provides operations to call the send method.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MicrosoftGraphSend()(*MailFoldersItemMessagesItemMicrosoftGraphSendRequestBuilder) {
    return NewMailFoldersItemMessagesItemMicrosoftGraphSendRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedProperties()(*MailFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMailFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MailFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMailFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property messages in me
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
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
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedProperties()(*MailFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMailFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MailFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMailFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for me
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in me
func (m *MailFoldersItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *MailFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
