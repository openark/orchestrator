package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleAccessPackages\{accessPackage-id1}
type EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages/{accessPackage%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) Ref()(*EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
