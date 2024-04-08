package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.privacy entity.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters get subjectRightsRequests from privacy
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    m := &SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilder instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get subjectRightsRequests from privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// MicrosoftGraphGetFinalAttachment provides operations to call the getFinalAttachment method.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) MicrosoftGraphGetFinalAttachment()(*SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder) {
    return NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFinalReport provides operations to call the getFinalReport method.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) MicrosoftGraphGetFinalReport()(*SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder) {
    return NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Notes provides operations to manage the notes property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Notes()(*SubjectRightsRequestsItemNotesRequestBuilder) {
    return NewSubjectRightsRequestsItemNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NotesById provides operations to manage the notes property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) NotesById(id string)(*SubjectRightsRequestsItemNotesAuthoredNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote%2Did"] = id
    }
    return NewSubjectRightsRequestsItemNotesAuthoredNoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// Team provides operations to manage the team property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Team()(*SubjectRightsRequestsItemTeamRequestBuilder) {
    return NewSubjectRightsRequestsItemTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get subjectRightsRequests from privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
