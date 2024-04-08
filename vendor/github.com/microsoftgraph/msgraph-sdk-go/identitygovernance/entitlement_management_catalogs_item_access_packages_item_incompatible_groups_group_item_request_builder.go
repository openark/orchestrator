package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\catalogs\{accessPackageCatalog-id}\accessPackages\{accessPackage-id}\incompatibleGroups\{group-id}
type EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) Ref()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
