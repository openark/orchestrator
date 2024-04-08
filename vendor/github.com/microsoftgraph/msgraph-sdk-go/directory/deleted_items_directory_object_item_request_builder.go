package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedItemsDirectoryObjectItemRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
type DeletedItemsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedItemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters recently deleted items. Read-only. Nullable.
type DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedItemsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// DeletedItemsDirectoryObjectItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsDirectoryObjectItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedItemsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDeletedItemsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsDirectoryObjectItemRequestBuilder) {
    m := &DeletedItemsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/deletedItems/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeletedItemsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDeletedItemsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedItemsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deletedItems for directory
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get recently deleted items. Read-only. Nullable.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphApplication()(*DeletedItemsItemMicrosoftGraphApplicationRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*DeletedItemsItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*DeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*DeletedItemsItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*DeletedItemsItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphGroup()(*DeletedItemsItemMicrosoftGraphGroupRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphRestore()(*DeletedItemsItemMicrosoftGraphRestoreRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*DeletedItemsItemMicrosoftGraphUserRequestBuilder) {
    return NewDeletedItemsItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property deletedItems in directory
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// ToDeleteRequestInformation delete navigation property deletedItems for directory
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation recently deleted items. Read-only. Nullable.
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deletedItems in directory
func (m *DeletedItemsDirectoryObjectItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *DeletedItemsDirectoryObjectItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
