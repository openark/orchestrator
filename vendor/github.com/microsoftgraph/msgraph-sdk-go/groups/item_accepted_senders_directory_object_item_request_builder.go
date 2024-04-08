package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAcceptedSendersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\acceptedSenders\{directoryObject-id}
type ItemAcceptedSendersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemAcceptedSendersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAcceptedSendersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAcceptedSendersDirectoryObjectItemRequestBuilder) {
    m := &ItemAcceptedSendersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/acceptedSenders/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAcceptedSendersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAcceptedSendersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAcceptedSendersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAcceptedSendersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of group entities.
func (m *ItemAcceptedSendersDirectoryObjectItemRequestBuilder) Ref()(*ItemAcceptedSendersItemRefRequestBuilder) {
    return NewItemAcceptedSendersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
