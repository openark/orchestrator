package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type ItemItemsDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type ItemItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsDriveItemItemRequestBuilderGetQueryParameters
}
// ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Analytics()(*ItemItemsItemAnalyticsRequestBuilder) {
    return NewItemItemsItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Children()(*ItemItemsItemChildrenRequestBuilder) {
    return NewItemItemsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ChildrenById(id string)(*ItemItemsItemChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return NewItemItemsItemChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemsDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    m := &ItemItemsDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Content()(*ItemItemsItemContentRequestBuilder) {
    return NewItemItemsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property items for drives
func (m *ItemItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get all items contained in the drive. Read-only. Nullable.
func (m *ItemItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ListItem()(*ItemItemsItemListItemRequestBuilder) {
    return NewItemItemsItemListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckin provides operations to call the checkin method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphCheckin()(*ItemItemsItemMicrosoftGraphCheckinRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckout provides operations to call the checkout method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphCheckout()(*ItemItemsItemMicrosoftGraphCheckoutRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCopy provides operations to call the copy method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphCopy()(*ItemItemsItemMicrosoftGraphCopyRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateLink provides operations to call the createLink method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphCreateLink()(*ItemItemsItemMicrosoftGraphCreateLinkRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphCreateUploadSession()(*ItemItemsItemMicrosoftGraphCreateUploadSessionRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDelta provides operations to call the delta method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphDelta()(*ItemItemsItemMicrosoftGraphDeltaRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeltaWithToken provides operations to call the delta method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphDeltaWithToken(token *string)(*ItemItemsItemMicrosoftGraphDeltaWithTokenRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token)
}
// MicrosoftGraphFollow provides operations to call the follow method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphFollow()(*ItemItemsItemMicrosoftGraphFollowRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphGetActivitiesByInterval()(*ItemItemsItemMicrosoftGraphGetActivitiesByIntervalRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemItemsItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime)
}
// MicrosoftGraphInvite provides operations to call the invite method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphInvite()(*ItemItemsItemMicrosoftGraphInviteRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPreview provides operations to call the preview method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphPreview()(*ItemItemsItemMicrosoftGraphPreviewRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphRestore()(*ItemItemsItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSearchWithQ provides operations to call the search method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphSearchWithQ(q *string)(*ItemItemsItemMicrosoftGraphSearchWithQRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q)
}
// MicrosoftGraphUnfollow provides operations to call the unfollow method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphUnfollow()(*ItemItemsItemMicrosoftGraphUnfollowRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphValidatePermission provides operations to call the validatePermission method.
func (m *ItemItemsDriveItemItemRequestBuilder) MicrosoftGraphValidatePermission()(*ItemItemsItemMicrosoftGraphValidatePermissionRequestBuilder) {
    return NewItemItemsItemMicrosoftGraphValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property items in drives
func (m *ItemItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Permissions()(*ItemItemsItemPermissionsRequestBuilder) {
    return NewItemItemsItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) PermissionsById(id string)(*ItemItemsItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemItemsItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Subscriptions()(*ItemItemsItemSubscriptionsRequestBuilder) {
    return NewItemItemsItemSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) SubscriptionsById(id string)(*ItemItemsItemSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewItemItemsItemSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Thumbnails()(*ItemItemsItemThumbnailsRequestBuilder) {
    return NewItemItemsItemThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ThumbnailsById(id string)(*ItemItemsItemThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewItemItemsItemThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for drives
func (m *ItemItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *ItemItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in drives
func (m *ItemItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Versions()(*ItemItemsItemVersionsRequestBuilder) {
    return NewItemItemsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) VersionsById(id string)(*ItemItemsItemVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewItemItemsItemVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Workbook()(*ItemItemsItemWorkbookRequestBuilder) {
    return NewItemItemsItemWorkbookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
