package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.todoTask entity.
type TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetQueryParameters a collection of file attachments for the task.
type TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetQueryParameters
}
// NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderInternal instantiates a new AttachmentBaseItemRequestBuilder and sets the default values.
func NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) {
    m := &TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/attachments/{attachmentBase%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder instantiates a new AttachmentBaseItemRequestBuilder and sets the default values.
func NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) Content()(*TodoListsItemTasksItemAttachmentsItemValueContentRequestBuilder) {
    return NewTodoListsItemTasksItemAttachmentsItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property attachments for me
func (m *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of file attachments for the task.
func (m *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentBaseable), nil
}
// ToDeleteRequestInformation delete navigation property attachments for me
func (m *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of file attachments for the task.
func (m *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
