package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleGroups\{group-id}
type EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) Ref()(*EntitlementManagementAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
