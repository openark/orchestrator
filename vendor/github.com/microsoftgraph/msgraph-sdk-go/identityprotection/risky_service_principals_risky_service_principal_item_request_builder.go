package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters azure AD service principals that are at risk.
type RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters
}
// RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderInternal instantiates a new RiskyServicePrincipalItemRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) {
    m := &RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals/{riskyServicePrincipal%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder instantiates a new RiskyServicePrincipalItemRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable), nil
}
// History provides operations to manage the history property of the microsoft.graph.riskyServicePrincipal entity.
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) History()(*RiskyServicePrincipalsItemHistoryRequestBuilder) {
    return NewRiskyServicePrincipalsItemHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// HistoryById provides operations to manage the history property of the microsoft.graph.riskyServicePrincipal entity.
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) HistoryById(id string)(*RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyServicePrincipalHistoryItem%2Did"] = id
    }
    return NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property riskyServicePrincipals in identityProtection
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable), nil
}
// ToDeleteRequestInformation delete navigation property riskyServicePrincipals for identityProtection
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property riskyServicePrincipals in identityProtection
func (m *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
