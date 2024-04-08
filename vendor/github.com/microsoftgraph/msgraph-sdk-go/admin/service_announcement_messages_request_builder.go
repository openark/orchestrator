package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesRequestBuilder provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
type ServiceAnnouncementMessagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementMessagesRequestBuilderGetQueryParameters retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
type ServiceAnnouncementMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ServiceAnnouncementMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceAnnouncementMessagesRequestBuilderGetQueryParameters
}
// ServiceAnnouncementMessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesRequestBuilderInternal instantiates a new MessagesRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesRequestBuilder) {
    m := &ServiceAnnouncementMessagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementMessagesRequestBuilder instantiates a new MessagesRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ServiceAnnouncementMessagesRequestBuilder) Count()(*ServiceAnnouncementMessagesCountRequestBuilder) {
    return NewServiceAnnouncementMessagesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceannouncement-list-messages?view=graph-rest-1.0
func (m *ServiceAnnouncementMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceUpdateMessageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageCollectionResponseable), nil
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphArchive()(*ServiceAnnouncementMessagesMicrosoftGraphArchiveRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFavorite provides operations to call the favorite method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphFavorite()(*ServiceAnnouncementMessagesMicrosoftGraphFavoriteRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkRead provides operations to call the markRead method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphMarkRead()(*ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkUnread provides operations to call the markUnread method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphMarkUnread()(*ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphUnarchive()(*ServiceAnnouncementMessagesMicrosoftGraphUnarchiveRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnfavorite provides operations to call the unfavorite method.
func (m *ServiceAnnouncementMessagesRequestBuilder) MicrosoftGraphUnfavorite()(*ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteRequestBuilder) {
    return NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Post create new navigation property to messages for admin
func (m *ServiceAnnouncementMessagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceAnnouncementMessagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceUpdateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable), nil
}
// ToGetRequestInformation retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
func (m *ServiceAnnouncementMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to messages for admin
func (m *ServiceAnnouncementMessagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceAnnouncementMessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
