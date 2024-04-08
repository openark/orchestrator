package devices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredOwners\{directoryObject-id}
type ItemRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MicrosoftGraphAppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphAppRoleAssignment()(*ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndpoint casts the previous resource to endpoint.
func (m *ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphEndpoint()(*ItemRegisteredOwnersItemMicrosoftGraphEndpointRequestBuilder) {
    return NewItemRegisteredOwnersItemMicrosoftGraphEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemRegisteredOwnersItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredOwnersItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*ItemRegisteredOwnersItemMicrosoftGraphUserRequestBuilder) {
    return NewItemRegisteredOwnersItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of device entities.
func (m *ItemRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemRegisteredOwnersItemRefRequestBuilder) {
    return NewItemRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
