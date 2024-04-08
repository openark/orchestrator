package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder provides operations to call the cancel method.
type BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) {
    m := &BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/calendarView/{bookingAppointment%2Did}/microsoft.graph.cancel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel the specified bookingAppointment in the specified bookingBusiness and send a message to the involved customer and staff members.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingappointment-cancel?view=graph-rest-1.0
func (m *BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) Post(ctx context.Context, body BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation cancel the specified bookingAppointment in the specified bookingBusiness and send a message to the involved customer and staff members.
func (m *BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemCalendarViewItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
