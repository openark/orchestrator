package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int32
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItemFilterMicrosoftGraphApplyTopItemsFilterApplyTopItemsFilterPostRequestBody) SetCount(value *int32)() {
    m.count = value
}
