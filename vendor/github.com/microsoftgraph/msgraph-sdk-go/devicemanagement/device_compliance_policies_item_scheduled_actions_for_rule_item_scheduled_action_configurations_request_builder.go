package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
type DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
type DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetQueryParameters
}
// DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal instantiates a new ScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    m := &DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder instantiates a new ScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Count()(*DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsCountRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceActionItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemCollectionResponseable), nil
}
// Post create new navigation property to scheduledActionConfigurations for deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable), nil
}
// ToGetRequestInformation the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scheduledActionConfigurations for deviceManagement
func (m *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, requestConfiguration *DeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
