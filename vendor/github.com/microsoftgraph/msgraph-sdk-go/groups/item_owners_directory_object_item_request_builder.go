package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\owners\{directoryObject-id}
type ItemOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphApplication()(*ItemOwnersItemMicrosoftGraphApplicationRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphDevice()(*ItemOwnersItemMicrosoftGraphDeviceRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphGroup()(*ItemOwnersItemMicrosoftGraphGroupRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOrgContact casts the previous resource to orgContact.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphOrgContact()(*ItemOwnersItemMicrosoftGraphOrgContactRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemOwnersItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*ItemOwnersItemMicrosoftGraphUserRequestBuilder) {
    return NewItemOwnersItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
func (m *ItemOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemOwnersItemRefRequestBuilder) {
    return NewItemOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
