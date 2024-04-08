package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
type ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters the collection of single-value extended properties defined for the post. Read-only. Nullable.
type ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters
}
// ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal instantiates a new SingleValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    m := &ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder instantiates a new SingleValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property singleValueExtendedProperties for groups
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property singleValueExtendedProperties in groups
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property singleValueExtendedProperties for groups
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property singleValueExtendedProperties in groups
func (m *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SingleValueLegacyExtendedPropertyable, requestConfiguration *ItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
