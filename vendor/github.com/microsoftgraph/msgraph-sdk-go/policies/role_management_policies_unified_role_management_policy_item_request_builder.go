package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
type RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters specifies the various policies associated with scopes and roles.
type RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters
}
// RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    m := &RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleManagementPolicies for policies
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EffectiveRules provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) EffectiveRules()(*RoleManagementPoliciesItemEffectiveRulesRequestBuilder) {
    return NewRoleManagementPoliciesItemEffectiveRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EffectiveRulesById provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) EffectiveRulesById(id string)(*RoleManagementPoliciesItemEffectiveRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = id
    }
    return NewRoleManagementPoliciesItemEffectiveRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get specifies the various policies associated with scopes and roles.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable), nil
}
// Patch update the navigation property roleManagementPolicies in policies
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable), nil
}
// Rules provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Rules()(*RoleManagementPoliciesItemRulesRequestBuilder) {
    return NewRoleManagementPoliciesItemRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RulesById provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) RulesById(id string)(*RoleManagementPoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = id
    }
    return NewRoleManagementPoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleManagementPolicies for policies
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation specifies the various policies associated with scopes and roles.
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleManagementPolicies in policies
func (m *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, requestConfiguration *RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
