package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TodoListsItemTasksTodoTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.todoTaskList entity.
type TodoListsItemTasksTodoTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type TodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TodoListsItemTasksTodoTaskItemRequestBuilderGetQueryParameters
}
// TodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) Attachments()(*TodoListsItemTasksItemAttachmentsRequestBuilder) {
    return NewTodoListsItemTasksItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentsById(id string)(*TodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentBase%2Did"] = id
    }
    return NewTodoListsItemTasksItemAttachmentsAttachmentBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AttachmentSessions provides operations to manage the attachmentSessions property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentSessions()(*TodoListsItemTasksItemAttachmentSessionsRequestBuilder) {
    return NewTodoListsItemTasksItemAttachmentSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentSessionsById provides operations to manage the attachmentSessions property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) AttachmentSessionsById(id string)(*TodoListsItemTasksItemAttachmentSessionsAttachmentSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentSession%2Did"] = id
    }
    return NewTodoListsItemTasksItemAttachmentSessionsAttachmentSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ChecklistItems provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ChecklistItems()(*TodoListsItemTasksItemChecklistItemsRequestBuilder) {
    return NewTodoListsItemTasksItemChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChecklistItemsById provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ChecklistItemsById(id string)(*TodoListsItemTasksItemChecklistItemsChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return NewTodoListsItemTasksItemChecklistItemsChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewTodoListsItemTasksTodoTaskItemRequestBuilderInternal instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewTodoListsItemTasksTodoTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksTodoTaskItemRequestBuilder) {
    m := &TodoListsItemTasksTodoTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTodoListsItemTasksTodoTaskItemRequestBuilder instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewTodoListsItemTasksTodoTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoListsItemTasksTodoTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoListsItemTasksTodoTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for me
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) Extensions()(*TodoListsItemTasksItemExtensionsRequestBuilder) {
    return NewTodoListsItemTasksItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ExtensionsById(id string)(*TodoListsItemTasksItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewTodoListsItemTasksItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable), nil
}
// LinkedResources provides operations to manage the linkedResources property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) LinkedResources()(*TodoListsItemTasksItemLinkedResourcesRequestBuilder) {
    return NewTodoListsItemTasksItemLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LinkedResourcesById provides operations to manage the linkedResources property of the microsoft.graph.todoTask entity.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) LinkedResourcesById(id string)(*TodoListsItemTasksItemLinkedResourcesLinkedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource%2Did"] = id
    }
    return NewTodoListsItemTasksItemLinkedResourcesLinkedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property tasks in me
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable), nil
}
// ToDeleteRequestInformation delete navigation property tasks for me
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in me
func (m *TodoListsItemTasksTodoTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TodoTaskable, requestConfiguration *TodoListsItemTasksTodoTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
