package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    m := &ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/relations/{relation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property relations for groups
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// FromTerm provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm()(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemFromTermRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemFromTermRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable), nil
}
// Patch update the navigation property relations in groups
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable), nil
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Set()(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property relations for groups
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property relations in groups
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToTerm provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm()(*ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemToTermRequestBuilder) {
    return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenItemRelationsItemToTermRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
