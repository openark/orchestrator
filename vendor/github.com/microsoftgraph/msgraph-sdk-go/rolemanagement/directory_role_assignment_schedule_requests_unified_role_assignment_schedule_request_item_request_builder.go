package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters requests for active role assignments to principals through PIM.
type DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*DirectoryRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*DirectoryRoleAssignmentScheduleRequestsItemAppScopeRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*DirectoryRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get requests for active role assignments to principals through PIM.
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) MicrosoftGraphCancel()(*DirectoryRoleAssignmentScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*DirectoryRoleAssignmentScheduleRequestsItemPrincipalRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*DirectoryRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*DirectoryRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
