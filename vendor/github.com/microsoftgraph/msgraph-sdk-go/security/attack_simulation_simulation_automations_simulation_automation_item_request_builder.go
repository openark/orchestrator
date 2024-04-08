package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder provides operations to manage the simulationAutomations property of the microsoft.graph.attackSimulationRoot entity.
type AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetQueryParameters represents simulation automation created to run on a tenant.
type AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetQueryParameters
}
// AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderInternal instantiates a new SimulationAutomationItemRequestBuilder and sets the default values.
func NewAttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) {
    m := &AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/attackSimulation/simulationAutomations/{simulationAutomation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder instantiates a new SimulationAutomationItemRequestBuilder and sets the default values.
func NewAttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property simulationAutomations for security
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents simulation automation created to run on a tenant.
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable), nil
}
// Patch update the navigation property simulationAutomations in security
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSimulationAutomationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable), nil
}
// Runs provides operations to manage the runs property of the microsoft.graph.simulationAutomation entity.
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) Runs()(*AttackSimulationSimulationAutomationsItemRunsRequestBuilder) {
    return NewAttackSimulationSimulationAutomationsItemRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RunsById provides operations to manage the runs property of the microsoft.graph.simulationAutomation entity.
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) RunsById(id string)(*AttackSimulationSimulationAutomationsItemRunsSimulationAutomationRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["simulationAutomationRun%2Did"] = id
    }
    return NewAttackSimulationSimulationAutomationsItemRunsSimulationAutomationRunItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property simulationAutomations for security
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation represents simulation automation created to run on a tenant.
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property simulationAutomations in security
func (m *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SimulationAutomationable, requestConfiguration *AttackSimulationSimulationAutomationsSimulationAutomationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
