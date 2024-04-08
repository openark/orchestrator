package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRequestBuilder provides operations to manage the directory property of the microsoft.graph.roleManagement entity.
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRequestBuilderGetQueryParameters get directory from roleManagement
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters
}
// DirectoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get directory from roleManagement
func (m *DirectoryRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable), nil
}
// Patch update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable), nil
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignments()(*DirectoryRoleAssignmentsRequestBuilder) {
    return NewDirectoryRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentsById(id string)(*DirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return NewDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleAssignmentScheduleInstances provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleInstances()(*DirectoryRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentScheduleInstancesById provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*DirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return NewDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleRequests()(*DirectoryRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewDirectoryRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentScheduleRequestsById provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*DirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return NewDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleAssignmentSchedules provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentSchedules()(*DirectoryRoleAssignmentSchedulesRequestBuilder) {
    return NewDirectoryRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentSchedulesById provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleAssignmentSchedulesById(id string)(*DirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return NewDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleDefinitions()(*DirectoryRoleDefinitionsRequestBuilder) {
    return NewDirectoryRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleDefinitionsById(id string)(*DirectoryRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return NewDirectoryRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleEligibilityScheduleInstances provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleInstances()(*DirectoryRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleEligibilityScheduleInstancesById provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleEligibilityScheduleRequests provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleRequests()(*DirectoryRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleEligibilityScheduleRequestsById provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*DirectoryRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return NewDirectoryRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleEligibilitySchedules provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilitySchedules()(*DirectoryRoleEligibilitySchedulesRequestBuilder) {
    return NewDirectoryRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleEligibilitySchedulesById provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *DirectoryRequestBuilder) RoleEligibilitySchedulesById(id string)(*DirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return NewDirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get directory from roleManagement
func (m *DirectoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
