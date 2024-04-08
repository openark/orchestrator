package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRejectedSendersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\rejectedSenders\{directoryObject-id}
type ItemRejectedSendersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemRejectedSendersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRejectedSendersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedSendersDirectoryObjectItemRequestBuilder) {
    m := &ItemRejectedSendersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/rejectedSenders/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemRejectedSendersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRejectedSendersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRejectedSendersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRejectedSendersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
func (m *ItemRejectedSendersDirectoryObjectItemRequestBuilder) Ref()(*ItemRejectedSendersItemRefRequestBuilder) {
    return NewItemRejectedSendersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
