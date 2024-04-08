package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetQueryParameters references to a directory or domain of another organization whose users can request access.
type EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetQueryParameters
}
// EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderInternal instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property connectedOrganizations for identityGovernance
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExternalSponsors provides operations to manage the externalSponsors property of the microsoft.graph.connectedOrganization entity.
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) ExternalSponsors()(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.externalSponsors.item collection
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) ExternalSponsorsById(id string)(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConnectedOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable), nil
}
// InternalSponsors provides operations to manage the internalSponsors property of the microsoft.graph.connectedOrganization entity.
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) InternalSponsors()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.internalSponsors.item collection
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) InternalSponsorsById(id string)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property connectedOrganizations in identityGovernance
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConnectedOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable), nil
}
// ToDeleteRequestInformation delete navigation property connectedOrganizations for identityGovernance
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property connectedOrganizations in identityGovernance
func (m *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, requestConfiguration *EntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
