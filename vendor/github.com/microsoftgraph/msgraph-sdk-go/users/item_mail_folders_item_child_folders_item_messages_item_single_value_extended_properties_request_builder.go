package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
type ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters the collection of single-value extended properties defined for the message. Nullable.
type ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters struct {
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
// ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetQueryParameters
}
// ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/singleValueExtendedProperties{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) Count()(*ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesCountRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the collection of single-value extended properties defined for the message. Nullable.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSingleValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyCollectionResponseable), nil
}
// Post create new navigation property to singleValueExtendedProperties for users
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable), nil
}
// ToGetRequestInformation the collection of single-value extended properties defined for the message. Nullable.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to singleValueExtendedProperties for users
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
