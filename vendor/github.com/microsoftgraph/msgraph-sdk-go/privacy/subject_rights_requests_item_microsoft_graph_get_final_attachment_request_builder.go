package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder provides operations to call the getFinalAttachment method.
type SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderInternal instantiates a new MicrosoftGraphGetFinalAttachmentRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder) {
    m := &SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}/microsoft.graph.getFinalAttachment()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder instantiates a new MicrosoftGraphGetFinalAttachmentRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFinalAttachment
func (m *SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getFinalAttachment
func (m *SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemMicrosoftGraphGetFinalAttachmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
