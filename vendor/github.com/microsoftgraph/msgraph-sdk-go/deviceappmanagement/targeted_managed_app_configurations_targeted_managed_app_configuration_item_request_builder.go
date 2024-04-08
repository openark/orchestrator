package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters targeted managed app configurations.
type TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters
}
// TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Apps()(*TargetedManagedAppConfigurationsItemAppsRequestBuilder) {
    return NewTargetedManagedAppConfigurationsItemAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppsById provides operations to manage the apps property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) AppsById(id string)(*TargetedManagedAppConfigurationsItemAppsManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return NewTargetedManagedAppConfigurationsItemAppsManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*TargetedManagedAppConfigurationsItemAssignmentsRequestBuilder) {
    return NewTargetedManagedAppConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*TargetedManagedAppConfigurationsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return NewTargetedManagedAppConfigurationsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    m := &TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*TargetedManagedAppConfigurationsItemDeploymentSummaryRequestBuilder) {
    return NewTargetedManagedAppConfigurationsItemDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get targeted managed app configurations.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) MicrosoftGraphAssign()(*TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) {
    return NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTargetApps provides operations to call the targetApps method.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) MicrosoftGraphTargetApps()(*TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsRequestBuilder) {
    return NewTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, requestConfiguration *TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
