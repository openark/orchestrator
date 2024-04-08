package print

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SharesItemAllowedGroupsGroupItemRequestBuilder builds and executes requests for operations under \print\shares\{printerShare-id}\allowedGroups\{group-id}
type SharesItemAllowedGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewSharesItemAllowedGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewSharesItemAllowedGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedGroupsGroupItemRequestBuilder) {
    m := &SharesItemAllowedGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/shares/{printerShare%2Did}/allowedGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSharesItemAllowedGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewSharesItemAllowedGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharesItemAllowedGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharesItemAllowedGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of print entities.
func (m *SharesItemAllowedGroupsGroupItemRequestBuilder) Ref()(*SharesItemAllowedGroupsItemRefRequestBuilder) {
    return NewSharesItemAllowedGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
