package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
type EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
type EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
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
// EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Count()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesCountRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accesspackage-list-incompatibleaccesspackages?view=graph-rest-1.0
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Ref()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRefRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
