package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.serviceUpdateMessage entity.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters a collection of serviceAnnouncementAttachments.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal instantiates a new ServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    m := &ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}/attachments/{serviceAnnouncementAttachment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder instantiates a new ServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the admin entity.
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Content()(*ServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilder) {
    return NewServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property attachments for admin
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of serviceAnnouncementAttachments.
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable), nil
}
// Patch update the navigation property attachments in admin
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable), nil
}
// ToDeleteRequestInformation delete navigation property attachments for admin
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of serviceAnnouncementAttachments.
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attachments in admin
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementAttachmentable, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
