package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a servicePrincipal object.
type ServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalItemRequestBuilderGetQueryParameters
}
// ServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppRoleAssignedTo provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*ItemAppRoleAssignedToRequestBuilder) {
    return NewItemAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignedToById provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedToById(id string)(*ItemAppRoleAssignedToAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewItemAppRoleAssignedToAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignmentsById(id string)(*ItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*ItemClaimsMappingPoliciesRequestBuilder) {
    return NewItemClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.claimsMappingPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPoliciesById(id string)(*ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) CreatedObjects()(*ItemCreatedObjectsRequestBuilder) {
    return NewItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) CreatedObjectsById(id string)(*ItemCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DelegatedPermissionClassifications provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*ItemDelegatedPermissionClassificationsRequestBuilder) {
    return NewItemDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DelegatedPermissionClassificationsById provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassificationsById(id string)(*ItemDelegatedPermissionClassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification%2Did"] = id
    }
    return NewItemDelegatedPermissionClassificationsDelegatedPermissionClassificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete a servicePrincipal object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceprincipal-delete?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Endpoints()(*ItemEndpointsRequestBuilder) {
    return NewItemEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EndpointsById provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) EndpointsById(id string)(*ItemEndpointsEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return NewItemEndpointsEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentials()(*ItemFederatedIdentityCredentialsRequestBuilder) {
    return NewItemFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FederatedIdentityCredentialsById provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*ItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return NewItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get retrieve the properties and relationships of a servicePrincipal object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceprincipal-get?view=graph-rest-1.0
func (m *ServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.homeRealmDiscoveryPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) MemberOfById(id string)(*ItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAddKey provides operations to call the addKey method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphAddKey()(*ItemMicrosoftGraphAddKeyRequestBuilder) {
    return NewItemMicrosoftGraphAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAddPassword provides operations to call the addPassword method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphAddPassword()(*ItemMicrosoftGraphAddPasswordRequestBuilder) {
    return NewItemMicrosoftGraphAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAddTokenSigningCertificate provides operations to call the addTokenSigningCertificate method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphAddTokenSigningCertificate()(*ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder) {
    return NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*ItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*ItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveKey provides operations to call the removeKey method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphRemoveKey()(*ItemMicrosoftGraphRemoveKeyRequestBuilder) {
    return NewItemMicrosoftGraphRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemovePassword provides operations to call the removePassword method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphRemovePassword()(*ItemMicrosoftGraphRemovePasswordRequestBuilder) {
    return NewItemMicrosoftGraphRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *ServicePrincipalItemRequestBuilder) MicrosoftGraphRestore()(*ItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*ItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) OwnedObjects()(*ItemOwnedObjectsRequestBuilder) {
    return NewItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) OwnedObjectsById(id string)(*ItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.servicePrincipals.item.owners.item collection
func (m *ServicePrincipalItemRequestBuilder) OwnersById(id string)(*ItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update entity in servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// ToDeleteRequestInformation delete a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*ItemTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePoliciesById(id string)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*ItemTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePoliciesById(id string)(*ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToPatchRequestInformation update entity in servicePrincipals
func (m *ServicePrincipalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOfById(id string)(*ItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
