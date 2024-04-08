package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
type ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters the to [term] of the relation. The term to which the relationship is defined.
type ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) {
    m := &ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did1}/sets/{set%2Did}/children/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}/toTerm{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
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
// ToGetRequestInformation the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
