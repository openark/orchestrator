package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Children()(*ItemTermStoreGroupsItemSetsItemChildrenItemChildrenRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ChildrenById(id string)(*ItemTermStoreGroupsItemSetsItemChildrenItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property children for sites
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get children terms of set in term [store].
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Patch update the navigation property children in sites
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Relations()(*ItemTermStoreGroupsItemSetsItemChildrenItemRelationsRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) RelationsById(id string)(*ItemTermStoreGroupsItemSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Set()(*ItemTermStoreGroupsItemSetsItemChildrenItemSetRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property children for sites
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation children terms of set in term [store].
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property children in sites
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
