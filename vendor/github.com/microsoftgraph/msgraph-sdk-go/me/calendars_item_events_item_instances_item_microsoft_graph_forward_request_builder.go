package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder provides operations to call the forward method.
type CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal instantiates a new MicrosoftGraphForwardRequestBuilder and sets the default values.
func NewCalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder) {
    m := &CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/microsoft.graph.forward";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder instantiates a new MicrosoftGraphForwardRequestBuilder and sets the default values.
func NewCalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this action allows the organizer or attendee of a meeting event to forward the meeting request to a new recipient.  If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this action also sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer's copy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-forward?view=graph-rest-1.0
func (m *CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder) Post(ctx context.Context, body CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardForwardPostRequestBodyable, requestConfiguration *CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation this action allows the organizer or attendee of a meeting event to forward the meeting request to a new recipient.  If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this action also sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer's copy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
func (m *CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardForwardPostRequestBodyable, requestConfiguration *CalendarsItemEventsItemInstancesItemMicrosoftGraphForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
