package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
type TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
type TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters
}
// TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal instantiates a new AgreementFileVersionItemRequestBuilder and sets the default values.
func NewTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    m := &TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file/localizations/{agreementFileLocalization%2Did}/versions/{agreementFileVersion%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder instantiates a new AgreementFileVersionItemRequestBuilder and sets the default values.
func NewTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property versions for identityGovernance
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable), nil
}
// Patch update the navigation property versions in identityGovernance
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable), nil
}
// ToDeleteRequestInformation delete navigation property versions for identityGovernance
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property versions in identityGovernance
func (m *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, requestConfiguration *TermsOfUseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
