package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody 
type ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The percent property
    percent *int32
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["percent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercent(val)
        }
        return nil
    }
    return res
}
// GetPercent gets the percent property value. The percent property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) GetPercent()(*int32) {
    return m.percent
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("percent", m.GetPercent())
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPercent sets the percent property value. The percent property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyBottomPercentFilterApplyBottomPercentFilterPostRequestBody) SetPercent(value *int32)() {
    m.percent = value
}
