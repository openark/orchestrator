package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder provides operations to call the unpublish method.
type BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderInternal instantiates a new MicrosoftGraphUnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder) {
    m := &BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/microsoft.graph.unpublish";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder instantiates a new MicrosoftGraphUnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post make the scheduling page of this business not available to external customers. Set the **isPublished** property to `false`, and the **publicUrl** property to `null`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingbusiness-unpublish?view=graph-rest-1.0
func (m *BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation make the scheduling page of this business not available to external customers. Set the **isPublished** property to `false`, and the **publicUrl** property to `null`.
func (m *BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemMicrosoftGraphUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
