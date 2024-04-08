package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody 
type ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fields property
    fields []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSortFieldable
    // The matchCase property
    matchCase *bool
    // The method property
    method *string
}
// NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody instantiates a new ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody()(*ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) {
    m := &ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookSortFieldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSortFieldable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSortFieldable)
            }
            m.SetFields(res)
        }
        return nil
    }
    res["matchCase"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchCase(val)
        }
        return nil
    }
    res["method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethod(val)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The fields property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) GetFields()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSortFieldable) {
    return m.fields
}
// GetMatchCase gets the matchCase property value. The matchCase property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) GetMatchCase()(*bool) {
    return m.matchCase
}
// GetMethod gets the method property value. The method property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) GetMethod()(*string) {
    return m.method
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("matchCase", m.GetMatchCase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("method", m.GetMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFields sets the fields property value. The fields property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) SetFields(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSortFieldable)() {
    m.fields = value
}
// SetMatchCase sets the matchCase property value. The matchCase property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) SetMatchCase(value *bool)() {
    m.matchCase = value
}
// SetMethod sets the method property value. The method property
func (m *ItemItemsItemWorkbookTablesItemSortMicrosoftGraphApplyApplyPostRequestBody) SetMethod(value *string)() {
    m.method = value
}
