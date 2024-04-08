package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters requests for active role assignments to principals through PIM.
type EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*EntitlementManagementRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*EntitlementManagementRoleAssignmentScheduleRequestsItemAppScopeRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*EntitlementManagementRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get requests for active role assignments to principals through PIM.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable), nil
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) MicrosoftGraphCancel()(*EntitlementManagementRoleAssignmentScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*EntitlementManagementRoleAssignmentScheduleRequestsItemPrincipalRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*EntitlementManagementRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*EntitlementManagementRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation requests for active role assignments to principals through PIM.
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *EntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
