package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.user entity.
type ItemMessagesMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMessagesMessageItemRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type ItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) Attachments()(*ItemMessagesItemAttachmentsRequestBuilder) {
    return NewItemMessagesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) AttachmentsById(id string)(*ItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesMessageItemRequestBuilder) {
    m := &ItemMessagesMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *ItemMessagesMessageItemRequestBuilder) Content()(*ItemMessagesItemValueContentRequestBuilder) {
    return NewItemMessagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property messages for users
func (m *ItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemMessagesMessageItemRequestBuilder) Extensions()(*ItemMessagesItemExtensionsRequestBuilder) {
    return NewItemMessagesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) ExtensionsById(id string)(*ItemMessagesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemMessagesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *ItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
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
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphCopy()(*ItemMessagesItemMicrosoftGraphCopyRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateForward provides operations to call the createForward method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateForward()(*ItemMessagesItemMicrosoftGraphCreateForwardRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReply provides operations to call the createReply method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReply()(*ItemMessagesItemMicrosoftGraphCreateReplyRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateReplyAll provides operations to call the createReplyAll method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphCreateReplyAll()(*ItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphForward()(*ItemMessagesItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMove provides operations to call the move method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphMove()(*ItemMessagesItemMicrosoftGraphMoveRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReply provides operations to call the reply method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphReply()(*ItemMessagesItemMicrosoftGraphReplyRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReplyAll provides operations to call the replyAll method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphReplyAll()(*ItemMessagesItemMicrosoftGraphReplyAllRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSend provides operations to call the send method.
func (m *ItemMessagesMessageItemRequestBuilder) MicrosoftGraphSend()(*ItemMessagesItemMicrosoftGraphSendRequestBuilder) {
    return NewItemMessagesItemMicrosoftGraphSendRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) MultiValueExtendedProperties()(*ItemMessagesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemMessagesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property messages in users
func (m *ItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
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
func (m *ItemMessagesMessageItemRequestBuilder) SingleValueExtendedProperties()(*ItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *ItemMessagesMessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property messages for users
func (m *ItemMessagesMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *ItemMessagesMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMessagesMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *ItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
