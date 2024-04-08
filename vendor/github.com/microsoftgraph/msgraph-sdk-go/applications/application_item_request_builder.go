package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ApplicationItemRequestBuilder provides operations to manage the collection of application entities.
type ApplicationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplicationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationItemRequestBuilderGetQueryParameters get the properties and relationships of an application object.
type ApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationItemRequestBuilderGetQueryParameters
}
// ApplicationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    m := &ApplicationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) CreatedOnBehalfOf()(*ItemCreatedOnBehalfOfRequestBuilder) {
    return NewItemCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/application-delete?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExtensionProperties provides operations to manage the extensionProperties property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) ExtensionProperties()(*ItemExtensionPropertiesRequestBuilder) {
    return NewItemExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionPropertiesById provides operations to manage the extensionProperties property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) ExtensionPropertiesById(id string)(*ItemExtensionPropertiesExtensionPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extensionProperty%2Did"] = id
    }
    return NewItemExtensionPropertiesExtensionPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentials()(*ItemFederatedIdentityCredentialsRequestBuilder) {
    return NewItemFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FederatedIdentityCredentialsById provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*ItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return NewItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get the properties and relationships of an application object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/application-get?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// HomeRealmDiscoveryPoliciesById provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Logo provides operations to manage the media for the application entity.
func (m *ApplicationItemRequestBuilder) Logo()(*ItemLogoRequestBuilder) {
    return NewItemLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAddKey provides operations to call the addKey method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphAddKey()(*ItemMicrosoftGraphAddKeyRequestBuilder) {
    return NewItemMicrosoftGraphAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAddPassword provides operations to call the addPassword method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphAddPassword()(*ItemMicrosoftGraphAddPasswordRequestBuilder) {
    return NewItemMicrosoftGraphAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*ItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*ItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveKey provides operations to call the removeKey method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphRemoveKey()(*ItemMicrosoftGraphRemoveKeyRequestBuilder) {
    return NewItemMicrosoftGraphRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemovePassword provides operations to call the removePassword method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphRemovePassword()(*ItemMicrosoftGraphRemovePasswordRequestBuilder) {
    return NewItemMicrosoftGraphRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphRestore()(*ItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetVerifiedPublisher provides operations to call the setVerifiedPublisher method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphSetVerifiedPublisher()(*ItemMicrosoftGraphSetVerifiedPublisherRequestBuilder) {
    return NewItemMicrosoftGraphSetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnsetVerifiedPublisher provides operations to call the unsetVerifiedPublisher method.
func (m *ApplicationItemRequestBuilder) MicrosoftGraphUnsetVerifiedPublisher()(*ItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilder) {
    return NewItemMicrosoftGraphUnsetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.owners.item collection
func (m *ApplicationItemRequestBuilder) OwnersById(id string)(*ItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of an application object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/application-update?view=graph-rest-1.0
func (m *ApplicationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// ToDeleteRequestInformation delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of an application object.
func (m *ApplicationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) TokenIssuancePolicies()(*ItemTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.tokenIssuancePolicies.item collection
func (m *ApplicationItemRequestBuilder) TokenIssuancePoliciesById(id string)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.application entity.
func (m *ApplicationItemRequestBuilder) TokenLifetimePolicies()(*ItemTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.tokenLifetimePolicies.item collection
func (m *ApplicationItemRequestBuilder) TokenLifetimePoliciesById(id string)(*ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToPatchRequestInformation update the properties of an application object.
func (m *ApplicationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
