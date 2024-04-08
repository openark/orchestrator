package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse 
type ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *int32
}
// NewItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse instantiates a new ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse and sets the default values.
func NewItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse()(*ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) {
    m := &ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ItemItemsItemWorkbookTablesItemRowsMicrosoftGraphCountCountResponse) SetValue(value *int32)() {
    m.value = value
}
