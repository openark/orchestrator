package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsOfUseAgreementsItemFileRequestBuilder provides operations to manage the file property of the microsoft.graph.agreement entity.
type TermsOfUseAgreementsItemFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TermsOfUseAgreementsItemFileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsOfUseAgreementsItemFileRequestBuilderGetQueryParameters retrieve the details of the default file for an agreement, including the language and version information. The file information is specified through the agreementFile object.
type TermsOfUseAgreementsItemFileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsOfUseAgreementsItemFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsOfUseAgreementsItemFileRequestBuilderGetQueryParameters
}
// TermsOfUseAgreementsItemFileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsOfUseAgreementsItemFileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsOfUseAgreementsItemFileRequestBuilderInternal instantiates a new FileRequestBuilder and sets the default values.
func NewTermsOfUseAgreementsItemFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsOfUseAgreementsItemFileRequestBuilder) {
    m := &TermsOfUseAgreementsItemFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTermsOfUseAgreementsItemFileRequestBuilder instantiates a new FileRequestBuilder and sets the default values.
func NewTermsOfUseAgreementsItemFileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsOfUseAgreementsItemFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsOfUseAgreementsItemFileRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property file for identityGovernance
func (m *TermsOfUseAgreementsItemFileRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the details of the default file for an agreement, including the language and version information. The file information is specified through the agreementFile object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/agreementfile-get?view=graph-rest-1.0
func (m *TermsOfUseAgreementsItemFileRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable), nil
}
// Localizations provides operations to manage the localizations property of the microsoft.graph.agreementFile entity.
func (m *TermsOfUseAgreementsItemFileRequestBuilder) Localizations()(*TermsOfUseAgreementsItemFileLocalizationsRequestBuilder) {
    return NewTermsOfUseAgreementsItemFileLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LocalizationsById provides operations to manage the localizations property of the microsoft.graph.agreementFile entity.
func (m *TermsOfUseAgreementsItemFileRequestBuilder) LocalizationsById(id string)(*TermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization%2Did"] = id
    }
    return NewTermsOfUseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property file in identityGovernance
func (m *TermsOfUseAgreementsItemFileRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable), nil
}
// ToDeleteRequestInformation delete navigation property file for identityGovernance
func (m *TermsOfUseAgreementsItemFileRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the details of the default file for an agreement, including the language and version information. The file information is specified through the agreementFile object.
func (m *TermsOfUseAgreementsItemFileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property file in identityGovernance
func (m *TermsOfUseAgreementsItemFileRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileable, requestConfiguration *TermsOfUseAgreementsItemFileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
