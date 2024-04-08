package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\connectedOrganizations\{connectedOrganization-id}\externalSponsors\{directoryObject-id}
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder) Ref()(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsItemRefRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
