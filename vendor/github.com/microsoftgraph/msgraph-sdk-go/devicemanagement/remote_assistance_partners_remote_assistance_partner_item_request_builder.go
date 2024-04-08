package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
type RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters the remote assist partners.
type RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters
}
// RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderInternal instantiates a new RemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) {
    m := &RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder instantiates a new RemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property remoteAssistancePartners for deviceManagement
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the remote assist partners.
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRemoteAssistancePartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable), nil
}
// MicrosoftGraphBeginOnboarding provides operations to call the beginOnboarding method.
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) MicrosoftGraphBeginOnboarding()(*RemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingRequestBuilder) {
    return NewRemoteAssistancePartnersItemMicrosoftGraphBeginOnboardingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisconnect provides operations to call the disconnect method.
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) MicrosoftGraphDisconnect()(*RemoteAssistancePartnersItemMicrosoftGraphDisconnectRequestBuilder) {
    return NewRemoteAssistancePartnersItemMicrosoftGraphDisconnectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property remoteAssistancePartners in deviceManagement
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRemoteAssistancePartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable), nil
}
// ToDeleteRequestInformation delete navigation property remoteAssistancePartners for deviceManagement
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the remote assist partners.
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property remoteAssistancePartners in deviceManagement
func (m *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, requestConfiguration *RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
