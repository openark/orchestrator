package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\connectedOrganizations\{connectedOrganization-id}\internalSponsors\{directoryObject-id}
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder) Ref()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsItemRefRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
