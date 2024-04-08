package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \directoryRoles\{directoryRole-id}\members\{directoryObject-id}
type ItemMembersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole%2Did}/members/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphApplication()(*ItemMembersItemMicrosoftGraphApplicationRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphDevice()(*ItemMembersItemMicrosoftGraphDeviceRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphGroup()(*ItemMembersItemMicrosoftGraphGroupRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOrgContact casts the previous resource to orgContact.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphOrgContact()(*ItemMembersItemMicrosoftGraphOrgContactRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemMembersItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*ItemMembersItemMicrosoftGraphUserRequestBuilder) {
    return NewItemMembersItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of directoryRole entities.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Ref()(*ItemMembersItemRefRequestBuilder) {
    return NewItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
