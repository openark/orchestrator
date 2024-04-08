package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\catalogs\{accessPackageCatalog-id}\accessPackages\{accessPackage-id}\incompatibleAccessPackages\{accessPackage-id1}
type EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages/{accessPackage%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) Ref()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
