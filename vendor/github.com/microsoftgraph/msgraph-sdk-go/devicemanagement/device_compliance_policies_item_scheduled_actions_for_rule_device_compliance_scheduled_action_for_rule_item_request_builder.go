package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
type DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
type DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal instantiates a new DeviceComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    m := &DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder instantiates a new DeviceComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable), nil
}
// Patch update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable), nil
}
// ScheduledActionConfigurations provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurations()(*DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScheduledActionConfigurationsById provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurationsById(id string)(*DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceActionItem%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
